package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/pkg/player"
	"pathe.co/zinx/utils"
	"pathe.co/zinx/ziface"
)

type Connection struct {
	TcpServer    ziface.IServer
	Conn         *net.TCPConn           // Current connection's socket TCP socket
	ConnID       uint32                 // Current connection's ID, also known as SessionID, ID is globally unique
	isClosed     bool                   // Current connection's close status
	MsgHandler   ziface.IMsgHandle      // Message handler, which manages MsgIDs and corresponding handling methods
	ExitBuffChan chan bool              // Channel to inform that the connection has exited/stopped
	msgChan      chan []byte            // Unbuffered channel for message communication between the read and write Goroutines
	msgBuffChan  chan []byte            // The buffered channel used for message communication between the reading and writing goroutines
	property     map[string]interface{} // Connection properties
	propertyLock sync.RWMutex           // Lock for protecting concurrent property modifications

	Player player.Player
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:    server,
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		MsgHandler:   msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		property:     make(map[string]interface{}),
	}

	c.TcpServer.GetConnMgr().Add(c)
	p := *player.NewPlayer()
	c.TcpServer.GetGame().AddPlayer(&p)

	b, _ := p.SetForTransport()
	c.SendBuffMsg(0, []byte(b))
	c.SendBuffMsg(1, []byte(c.Player.ShowTakes()))

	return c
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			fmt.Println("Data to be sent to client", data)
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:", err, "Conn Writer exit")
				return
			}

		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buffered Data error:", err, "Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}

		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Start(g *game.Game) {
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.ConnID)
	// If the current connection is already closed
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// ==================
	// If the user registered a callback function for this connection's closure, it should be called explicitly at this moment
	c.TcpServer.CallOnConnStop(c)
	// Close the socket connection
	c.Conn.Close()
	// Close the writer
	c.ExitBuffChan <- true
	// Remove the connection from the connection manager
	c.TcpServer.GetConnMgr().Remove(c)
	// Close all channels of this connection
	close(c.ExitBuffChan)
	close(c.msgBuffChan)
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// Create a data packing/unpacking object
		dp := NewDataPack()

		// Read the client's message header
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			c.ExitBuffChan <- true
			continue
		}

		// Unpack the message, obtain msgid and datalen, and store them in msg
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			c.ExitBuffChan <- true
			continue
		}

		// Read the data based on dataLen and store it in msg.Data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// Get the current connection ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// Get the remote client address information
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send the Message data to the remote TCP client
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	fmt.Println(data)

	// Package the data and send it
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id =", msgId)
		return errors.New("Pack error msg")
	}

	// Write back to the client
	// Change the previous direct write using conn.Write to sending the message to the Channel for the Writer to read
	c.msgChan <- msg

	return nil
}

func (c *Connection) SendBuffMsg(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when sending buffered message")
	}
	// Pack the data and send it
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Println("Pack error msg ID =", msgID)
		return errors.New("Pack error message")
	}

	// Write to the client
	c.msgBuffChan <- msg

	return nil
}

// Set connection property
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

// Get connection property
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

// Remove connection property
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

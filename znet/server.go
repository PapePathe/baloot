package znet

import (
	"errors"
	"fmt"
	"net"

	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/utils"
	"pathe.co/zinx/ziface"
)

// Implementation of the iServer interface, defining a Server service class.
type Server struct {
	ConnMgr     ziface.IConnManager
	IPVersion   string
	IP          string
	game        game.Game
	msgHandler  ziface.IMsgHandle
	Name        string
	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
	Port        int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// echo business
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)
	// Start a goroutine to do the server Linster business.
	go func() {
		// 0 Start the Worker pool mechanism
		s.msgHandler.StartWorkerPool()

		// 1. Get a TCP address.
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2. Listen to the server address.
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		// The server has successfully started listening.
		fmt.Println("start Zinx server  ", s.Name, " succ, now listening...")

		// TODO server.go should have a method to generate ID automatically
		var cid uint32 = 0
		// 3. Start the server network connection business.
		for {

			// 3.1. block and wait for the client to establish a connection request
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// 2. Set the maximum connection limit for the server
			//    If the limit is exceeded, close the new connection
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			// 3.2. TODO Server.Start() set the maximum connection control of the server. If the maximum connection is exceeded, then close this new connection

			// 3.3. handle the new connection request business, where handler and conn are bound
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++

			// 3.4. start the handling business of the current connection
			go dealConn.Start(&s.game)
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	// Server.Stop() needs to stop or clean up other connection information or other information that needs to be cleared.
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	select {}
}

// Router function: Register a router business method for the current server to handle client connections
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)

	fmt.Println("Add Router success!")
}

// GetConnMgr returns the connection manager
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

/*
Create a server instance
*/
func NewServer() ziface.IServer {
	// Initialize the global configuration file first
	utils.GlobalObject.Reload()

	s := &Server{
		Name:       utils.GlobalObject.Name, // Get from global parameters
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,    // Get from global parameters
		Port:       utils.GlobalObject.TcpPort, // Get from global parameters
		msgHandler: NewMsgHandle(),             // Initialize msgHandler
		ConnMgr:    NewConnManager(),           // Create a ConnManager
		game:       *game.NewGame(),
	}
	return s
}

func (s *Server) GetGame() *game.Game {
	return &s.game
}

// Set the hook function to be called when a connection is created for the server
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// Set the hook function to be called when a connection is about to be disconnected for the server
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// Invoke the OnConnStart hook function for the connection
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// Invoke the OnConnStop hook function for the connection
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

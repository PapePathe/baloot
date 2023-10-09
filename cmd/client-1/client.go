package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"pathe.co/zinx/pkg/player"
	"pathe.co/zinx/znet"
)

/*
Simulate client
*/
func main() {

	fmt.Println("Client Test ... start")
	// Wait for 3 seconds before sending the test request to give the server a chance to start
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		msg, err := readMessage(conn)
		if err != nil {
			fmt.Println(err)
		}
		switch msg.Id {
		case 0:
			player := player.NewPlayer()
			fmt.Println("Player Hand Sync")
			if err := player.GetFromTransport(msg.Data); err != nil {
				fmt.Println(err)
			}
			player.PreetyHand()
			fmt.Println("What is your take? \nTout\n Cent\n Pique\n Coeur\n Carreau\n Trefle")
			var take int
			fmt.Scanln(&take)
			if take > 0 && take < 7 {
				dp := znet.NewDataPack()
				msg, err := dp.Pack(znet.NewMsgPackage(2, []byte{byte(take)}))
				if err != nil {
					fmt.Println("Pack error msg ID =", msg)
				}
				conn.Write(msg)
			}
		case 1:
			fmt.Println("It's your turn to set your take, previous takes: ...")
		case 2:
			fmt.Println("It's your turn to play")
		case 3:
			fmt.Println("Another player added card on deck")
		case 4:
			fmt.Println("It's your turn to play")
		}

		time.Sleep(1 * time.Second)
	}
}

func readMessage(conn net.Conn) (*znet.Message, error) {
	// Pack the message
	dp := znet.NewDataPack()
	// Read the head part from the stream
	headData := make([]byte, dp.GetHeadLen())
	_, err := io.ReadFull(conn, headData) // ReadFull fills the buffer until it's full
	if err != nil {
		fmt.Println("read head error")
		return nil, err
	}

	// Unpack the headData into a message
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return nil, err
	}

	if msgHead.GetDataLen() > 0 {
		// The message has data, so we need to read the data part
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		// Read the data bytes from the stream based on the dataLen
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return nil, err
		}
		return msg, nil

	}
	return nil, errors.New("something went wrong reading message from connection")
}

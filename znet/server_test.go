package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

/*
Simulate the client
*/
func ClientTest() {

	fmt.Println("Client Test ... start")

	// Wait for 3 seconds to give the server a chance to start its service
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

// Test function for the Server module
func TestServer(t *testing.T) {

	/*
	   Server test
	*/
	// 1. Create a server handle s
	s := NewServer()

	/*
	   Client test
	*/
	go ClientTest()

	// 2. Start the service
	s.Serve()
}

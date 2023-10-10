package ziface

import (
	"net"

	"pathe.co/zinx/pkg/game"
)

// Define the connection interface
type IConnection interface {
	// Start the connection, making the current connection work
	Start(*game.Game)
	// Stop the connection, ending the current connection state
	Stop()
	// Get the raw socket TCPConn from the current connection
	GetTCPConnection() *net.TCPConn
	// Get the current connection ID
	GetConnID() uint32
	// Get the remote client's address information
	RemoteAddr() net.Addr
	// Directly send the Message data to the remote TCP client
	SendMsg(msgId uint32, data []byte) error
	// Send Message data directly to the remote TCP client (buffered)
	SendBuffMsg(msgID uint32, data []byte) error // Add buffered message sending interface
	// Set connection attributes
	SetProperty(key string, value interface{})
	// Get connection attributes
	GetProperty(key string) (interface{}, error)
	// Remove connection attributes
	RemoveProperty(key string)
}

// Define an interface for handling connection business uniformly
type HandFunc func(*net.TCPConn, []byte, int) error

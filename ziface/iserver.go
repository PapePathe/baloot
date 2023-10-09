package ziface

import "pathe.co/zinx/pkg/game"

// Define the server interface
type IServer interface {
	// Start the server method
	Start()
	// Stop the server method
	Stop()
	// Start the business service method
	Serve()
	// Router function: Register a router business method for the current server to handle client connections
	AddRouter(msgId uint32, router IRouter)
	// Get the connection manager
	GetConnMgr() IConnManager
	// Set the hook function to be called when a connection is created for this server
	SetOnConnStart(func(IConnection))
	// Set the hook function to be called when a connection is about to be disconnected for this server
	SetOnConnStop(func(IConnection))
	// Invoke the OnConnStart hook function for the connection
	CallOnConnStart(conn IConnection)
	// Invoke the OnConnStop hook function for the connection
	CallOnConnStop(conn IConnection)

	GetGame() *game.Game
}

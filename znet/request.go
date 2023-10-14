package znet

import (
	"pathe.co/zinx/pkg/game"
	"pathe.co/zinx/ziface"
)

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
	game *game.Game
	pid  int
}

// GetConnection retrieves the connection information of the request
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData retrieves the data of the request message
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// Get the ID of the request message
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}

func (r *Request) GetGame() *game.Game {
	return r.game
}

func (r *Request) GetPlayerID() int {
	return r.pid
}

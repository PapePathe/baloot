package router

import (
	"fmt"

	"pathe.co/zinx/ziface"
	"pathe.co/zinx/znet"
)

type PlayerTakeRouter struct {
	znet.BaseRouter
}

// PlayerTake Handle
func (this *PlayerTakeRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PlayerTakeRouter Handle")
	fmt.Println("recv from client: msgId=", request.GetMsgID(), ", data=")

	//takeID := int(request.GetData()[0])

	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

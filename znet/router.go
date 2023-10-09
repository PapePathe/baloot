package znet

import "pathe.co/zinx/ziface"

// When implementing a router, embed this base class and override its methods as needed
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}

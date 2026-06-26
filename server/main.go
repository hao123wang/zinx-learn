package main

import (
	"fmt"
	"zinx-learn/ziface"
	"zinx-learn/znet"
)

type EchoRouter struct {
	znet.BaseRouter
}

func (e *EchoRouter) PreHandle(req ziface.IRequest) {
	fmt.Printf("连接 %d prehandle\n", req.GetConnection().GetConnID())
}

func (e *EchoRouter) Handle(req ziface.IRequest) {
	fmt.Printf("连接 %d handle\n", req.GetConnection().GetConnID())
}

func (e *EchoRouter) PostHandle(req ziface.IRequest) {
	fmt.Printf("连接 %d posthandle\n", req.GetConnection().GetConnID())
}

func main() {
	s := znet.NewServer("zinx-learn", "0.0.0.0", 8080)
	s.AddRouter(&EchoRouter{})
	s.Start()
}

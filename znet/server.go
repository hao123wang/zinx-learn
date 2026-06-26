package znet

import (
	"fmt"
	"net"
	"zinx-learn/ziface"
)

type Server struct {
	Name      string // 服务器名称
	IPVersion string // ip版本，如tcp4
	IP        string // 服务器ip
	Port      int    // 服务器端口

	Router ziface.IRouter // 消息路由，负责处理接收到的request
}

func NewServer(name, ip string, port int) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp",
		IP:        ip,
		Port:      port,
	}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	fmt.Printf("[START] Server listener at addr: %s\n", addr)
	defer s.Stop()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	var connID uint32
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept conn err: %v\n", err)
		}
		// 为每个连接创建一个连接管理
		connID++
		connection := NewConnection(conn, connID, s.Router)
		go connection.Start()
	}
}

func handle(req ziface.IRequest) error {
	fmt.Printf("recv from conn: %d, req: %+v\n", req.GetConnection().GetConnID(), req.GetData())
	// 服务器打印信息
	fmt.Print(string(req.GetData()))
	// 回显给客户端
	req.GetConnection().GetConn().Write(req.GetData())

	return nil
}

func (s *Server) Stop() {
	fmt.Println("[END] server stop")
}

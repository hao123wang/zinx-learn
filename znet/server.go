package znet

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	Name      string // 服务器名称
	IPVersion string // ip版本，如tcp4
	IP        string // 服务器ip
	Port      int    // 服务器端口
}

func NewServer(name, ip string, port int) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp",
		IP:        ip,
		Port:      port,
	}
}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	fmt.Printf("[START] Server listener at addr: %s\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept conn err: %v\n", err)
		}
		// 为每个连接开启协程处理
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// 接收来自该连接的消息
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client conn exit\n")
				return
			}
			fmt.Printf("recv buf err: %v\n", err)
			continue
		}
		// 服务器打印信息
		fmt.Print(string(buf[:n]))
		// 回显给客户端
		conn.Write(buf[:n])
	}
}

func (s *Server) Stop() {

}

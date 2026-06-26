package znet

import (
	"fmt"
	"net"
	"zinx-learn/ziface"
)

type Connection struct {
	Conn     net.Conn // 当前连接
	ConnID   uint32   // 当前连接id
	IsClosed bool     // 当前连接是否已关闭

	HandlerFunc ziface.HandlerFunc // 连接的处理函数

	ExitBuffChan chan struct{}
}

func NewConnection(conn net.Conn, connID uint32, handlerFunc ziface.HandlerFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		IsClosed:     false,
		HandlerFunc:  handlerFunc,
		ExitBuffChan: make(chan struct{}, 1),
	}
	return c
}

func (c *Connection) read() {
	defer func() {
		fmt.Println("conn reader exit!")
		c.Stop()
	}()
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- struct{}{}
			return
		}

		// 将消息组装为request
		req := NewRequest(c, buf[:n])

		// 处理请求
		if err := c.HandlerFunc(req); err != nil {
			fmt.Println("connID ", c.ConnID, " handle is error")
			c.ExitBuffChan <- struct{}{}
			return
		}
	}
}

func (c *Connection) Start() {
	// 启动当前连接的读协程
	go c.read()

	// 监听退出通道
	<-c.ExitBuffChan
}

func (c *Connection) Stop() {
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true

	c.Conn.Close()
	c.ExitBuffChan <- struct{}{}
	close(c.ExitBuffChan)
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetConn() net.Conn {
	return c.Conn
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

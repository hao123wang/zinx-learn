package ziface

import "net"

type IConnection interface {
	Start()                  // 启动连接，让当前连接开始工作
	Stop()                   // 停止连接，结束当前连接状态
	GetConnID() uint32       // 获取客户端地址信息
	GetConn() net.Conn       // 获取底层连接
	GetRemoteAddr() net.Addr // 获取远程客户端的地址
}

type HandlerFunc func(request IRequest) error

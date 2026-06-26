package znet

import "zinx-learn/ziface"

type Request struct {
	connection ziface.IConnection
	data       []byte
}

func NewRequest(connection ziface.IConnection, data []byte) *Request {
	return &Request{
		connection: connection,
		data:       data,
	}
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.connection
}

func (r *Request) GetData() []byte {
	return r.data
}

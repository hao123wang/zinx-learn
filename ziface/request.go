package ziface

type IRequest interface {
	GetConnection() IConnection // 获取请求相关的连接信息
	GetData() []byte            // 获取请求数据
}

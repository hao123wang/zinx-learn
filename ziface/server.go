package ziface

type IServer interface {
	Start()
	Stop()
	AddRouter(router IRouter)
}

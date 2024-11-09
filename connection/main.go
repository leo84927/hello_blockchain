package connection

import "context"

var (
	Ctx, CtxClose = context.WithCancel(context.TODO())
)

type Conn interface {
	// Init 连线初始化
	Init()
	// Close 关闭连线
	Close()
}

func Init(option []Conn) {
	for _, instance := range option {
		instance.Init()
	}
}

func Close(option []Conn) {
	for _, instance := range option {
		instance.Close()
	}
}

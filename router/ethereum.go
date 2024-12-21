package router

import (
	"hello_blockchain/controller"

	"github.com/fasthttp/router"
)

func EthereumRouter() *router.Router {

	handler := newRoute()

	ethereumGp := handler.Group("/ethereum")
	{
		ethereumCtl := new(controller.EthereumController)
		{
			// 生成地址
			ethereumGp.GET("/create/address", ethereumCtl.CreateAddress)
			// 取得余额
			ethereumGp.GET("/get/balance", ethereumCtl.GetBalance)
		}
	}

	return handler
}

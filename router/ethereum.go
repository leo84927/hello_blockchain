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
		ethereumGp.GET("/get/balance", ethereumCtl.GetBalance)
	}

	return handler
}

package router

import (
	"fmt"
	"hello_blockchain/lib/log"
	"os"
	"runtime/debug"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func newRoute() *router.Router {
	// 不可直接返回 &router.Router{...}, router.New() 有做其他欄位的初始化
	handler := router.New()
	handler.PanicHandler = apiServerPanic
	handler.NotFound = notFound
	handler.RedirectFixedPath = false
	handler.RedirectTrailingSlash = false
	handler.HandleMethodNotAllowed = false
	return handler
}

func notFound(fctx *fasthttp.RequestCtx) {
	fctx.SetStatusCode(404)
}

func methodNotAllowed(fctx *fasthttp.RequestCtx) {
	fctx.SetStatusCode(405)
}

func apiServerPanic(fctx *fasthttp.RequestCtx, rcv interface{}) {

	msg := fmt.Errorf("%+v", rcv)

	fmt.Println(msg)
	errStack := debug.Stack()
	_, _ = os.Stderr.Write(errStack)
	log.LogError(log.FilePanic, msg, fmt.Sprintf("error stack:\n%s", errStack))

	if r := recover(); r != nil {
		fmt.Println("recovered failed", r)
	}

	fctx.SetStatusCode(500)
}

package middleware

import (
	_const "hello_blockchain/const"
	"hello_blockchain/helper"

	"github.com/valyala/fasthttp"
)

type CheckPoint func(fctx *fasthttp.RequestCtx) (_const.ErrorCode, error)

func Use(next fasthttp.RequestHandler, middlewareList []CheckPoint) fasthttp.RequestHandler {
	return func(fctx *fasthttp.RequestCtx) {
		for _, cb := range middlewareList {
			if code, err := cb(fctx); code != _const.Success || err != nil {
				helper.Fail(fctx, code, err.Error())
				return
			}
		}

		next(fctx)
	}
}

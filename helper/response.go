package helper

import (
	_const "hello_blockchain/const"
	"time"

	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

// successResponse 成功返回
type successResponse struct {
	Sign   int64            `json:"sign"`
	Code   _const.ErrorCode `json:"code"`
	Result interface{}      `json:"data"`
}

// failResponse 失敗返回
type failResponse struct {
	Sign    int64            `json:"sign"`
	Code    _const.ErrorCode `json:"code"`
	Message string           `json:"msg"`
	Result  interface{}      `json:"data"`
}

func Success(fctx *fasthttp.RequestCtx, result ...interface{}) {

	res := successResponse{
		Sign:   time.Now().UnixMilli(),
		Code:   _const.ErrorCode(200),
		Result: result,
	}

	// TODO 待驗證
	//// case 0: res = { "Sign": 1234567890, "Code": 200, "Result": {} }，no switch: res = { "Sign": 1234567890, "Code": 200, "Result": [] }
	//// case 1: res = { "Sign": 1234567890, "Code": 200, "Result": "data" }，no switch: res = { "Sign": 1234567890, "Code": 200, "Result": [data] }
	//switch len(result) {
	//case 0:
	//	res.Result = struct{}{}
	//case 1:
	//	res.Result = result[0]
	//default:
	//	res.Result = result
	//}

	b, err := sonic.Marshal(res)
	if err != nil {
		fctx.SetBody([]byte(err.Error()))
		return
	}

	fctx.SetBody(b)
}

func Fail(fctx *fasthttp.RequestCtx, errorCode _const.ErrorCode, message string, result ...interface{}) {

	res := failResponse{
		Sign:    time.Now().UnixMilli(),
		Code:    errorCode,
		Message: message,
		Result:  result,
	}
	//if result == nil {
	//	res.Result = struct{}{}
	//}

	b, err := sonic.Marshal(res)
	if err != nil {
		fctx.SetBody([]byte(err.Error()))
		return
	}

	fctx.SetBody(b)
}

package controller

import (
	_const "hello_blockchain/const"
	"hello_blockchain/helper"
	"hello_blockchain/lib/elk"
	"hello_blockchain/service"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type EthereumController struct{}

func (EthereumController) CreateAddress(fctx *fasthttp.RequestCtx) {

	address, err := service.EthereumService.CreateAddress(fctx)
	if err != nil {
		elk.LogError(logrus.ErrorLevel, err, nil)
		helper.Fail(fctx, _const.ErrorUnknown, err.Error())
		return
	}

	helper.Success(fctx, address)
}

func (EthereumController) GetBalance(fctx *fasthttp.RequestCtx) {

}

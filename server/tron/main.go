package main

import (
	"fmt"
	"hello_blockchain/connection"
	"hello_blockchain/connection/cache"
	"hello_blockchain/helper"
	"hello_blockchain/lib/log"

	"github.com/rotisserie/eris"
)

const (
	_serverName = "tron"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			panicErr := eris.New(fmt.Sprintf("%+v", r))
			log.LogPanic(panicErr)
			panic(r)
		}
	}()

	helper.SetServerName(_serverName)

	for _, instance := range []connection.Conn{
		cache.NewRedis(),
	} {
		instance.Init()
	}
}

package main

import (
	"fmt"
	"hello_blockchain/connection"
	"hello_blockchain/connection/cache"
	"hello_blockchain/connection/db"
	"hello_blockchain/helper"
	"hello_blockchain/lib/log"

	"github.com/rotisserie/eris"
)

const (
	_serverName = "ethereum"
)

func main() {

	// 根據server所需，指定要初始化哪些連線
	needInit := []connection.Conn{
		cache.NewRedis(),
		db.NewMysql(),
	}

	defer func() {
		connection.Close(needInit)
		if r := recover(); r != nil {
			panicErr := eris.New(fmt.Sprintf("%+v", r))
			log.LogPanic(panicErr)
			panic(r)
		}
	}()

	// 設定server名稱
	helper.SetServerName(_serverName)

	// 初始化
	connection.Init(needInit)
}

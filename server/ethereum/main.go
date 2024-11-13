package main

import (
	"fmt"
	"hello_blockchain/connection"
	"hello_blockchain/helper"
	"hello_blockchain/lib/log"

	"github.com/rotisserie/eris"
)

const (
	_serverName = "ethereum"
)

func main() {

	defer func() {
		// 關閉連線
		connection.Close()

		if r := recover(); r != nil {
			panicErr := eris.New(fmt.Sprintf("%+v", r))
			log.LogPanic(panicErr)
			panic(r)
		}
	}()

	// 設定server名稱
	helper.SetServerName(_serverName)

	// 根據server所需，指定要初始化哪些連線
	connection.InitClient(connection.ClientOptions{
		ServiceName: _serverName,
		NeedHttp:    true,
		NeedRedis:   true,
		NeedGorm:    true,
	})
}

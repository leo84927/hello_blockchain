package main

import (
	"fmt"
	"hello_blockchain/config"
	"hello_blockchain/connection"
	"hello_blockchain/helper"
	"hello_blockchain/lib/log"
	"hello_blockchain/middleware"
	"hello_blockchain/router"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/rotisserie/eris"
)

func main() {

	defer func() {
		// 關閉連線
		connection.Close()

		if r := recover(); r != nil {
			panicErr := eris.New(fmt.Sprintf("%+v", r))
			log.LogError(log.FilePanic, panicErr)
			panic(r)
		}
	}()

	// 設定server名稱
	serverName := config.EthereumServerName
	port := config.EthereumServerPort
	helper.SetServerName(serverName)
	fmt.Println("StartTime:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("ServerName:", helper.GetServerName())
	fmt.Println("ServerPort:", port)

	// 根據server所需，指定要初始化哪些連線
	connection.InitClient(connection.ClientOptions{
		ServiceName: serverName,
		NeedHttp:    true,
		NeedRedis:   true,
		NeedGorm:    true,
	})

	srv := &fasthttp.Server{
		Handler: fasthttp.TimeoutHandler(
			middleware.Use(router.EthereumRouter().Handler, nil),
			config.EthereumServerTimeout,
			`{"status": "false","data":"Server response timed out, please try again later."}`,
		),
		ReadTimeout:        config.EthereumServerTimeout,
		WriteTimeout:       config.EthereumServerTimeout,
		Name:               helper.GetServerName(),
		MaxRequestBodySize: 4 * 1024 * 1024,
	}
	fmt.Println("server running......")

	if err := srv.ListenAndServe(":" + port); err != nil {
		log.LogError(log.FilePanic, err, "Error in ListenAndServe")
		return
	}
}

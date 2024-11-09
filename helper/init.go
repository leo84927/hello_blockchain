package helper

import (
	"context"
)

var (
	Ctx, ctxClose = context.WithCancel(context.TODO())
	_serverName   string
)

func SetServerName(name string) {
	_serverName = name
}

func GetServerName() string {
	return _serverName
}

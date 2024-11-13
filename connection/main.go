package connection

import (
	"context"

	"github.com/valyala/fasthttp"
)

var (
	Ctx        = context.Background()
	ServerName string

	_fc             *fasthttp.Client
	_gormInterface  GormInterface
	_redisInterface RedisInterface
)

type ClientOptions struct {
	ServiceName string
	NeedHttp    bool
	NeedRedis   bool
	NeedGorm    bool
}

func InitClient(opt ClientOptions) {
	ServerName = opt.ServiceName

	if opt.NeedHttp {
		_fc = NewFasthttpClient()
	}
	if opt.NeedRedis {
		_redisInterface = NewRedisConn()
	}
	if opt.NeedGorm {
		_gormInterface = NewGormConn()
	}
}

func Close() {
	if _redisInterface != nil {
		_redisInterface.Close()
	}
	if _gormInterface != nil {
		_gormInterface.Close()
	}
}

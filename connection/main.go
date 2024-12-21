package connection

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/valyala/fasthttp"
)

var (
	Ctx        = context.Background()
	ServerName string

	_fc             *fasthttp.Client
	_gormInterface  GormInterface
	_redisInterface RedisInterface
	_ethClient      *ethclient.Client
)

type ClientOptions struct {
	ServiceName string
	NeedHttp    bool
	NeedRedis   bool
	NeedGorm    bool
	NeedEth     bool
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
	if opt.NeedEth {
		_ethClient = NewEthClient()
	}
}

func Close() {
	if _redisInterface != nil {
		_redisInterface.Close()
	}
	if _gormInterface != nil {
		_gormInterface.Close()
	}
	if _ethClient != nil {
		_ethClient.Close()
	}
}

package cache

import (
	"context"
	"hello_blockchain/config"

	"github.com/redis/go-redis/v9"
	"github.com/rotisserie/eris"
)

type _redis struct {
	client *redis.Client // redis连线
}

func NewRedis() *_redis {
	return &_redis{}
}

func (c *_redis) Init() {
	client := redis.NewClient(&redis.Options{
		Addr:         config.RedisConn.Addr,
		Password:     config.RedisConn.Password,
		DB:           config.RedisConn.DB,
		DialTimeout:  config.RedisConn.DialTimeout,
		ReadTimeout:  config.RedisConn.ReadTimeout,
		WriteTimeout: config.RedisConn.WriteTimeout,
		PoolSize:     config.RedisConn.PoolSize,
		PoolTimeout:  config.RedisConn.PoolTimeout,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(eris.Wrap(err, "redis init error"))
	} else {
		c.client = client
	}
}

func (c *_redis) Close() {
	if c.client != nil {
		err := c.client.Close()
		if err != nil {
			panic(eris.Wrap(err, "redis close error"))
		}
	}
}

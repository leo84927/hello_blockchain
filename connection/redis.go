package connection

import "github.com/redis/go-redis/v9"

type RedisInterface interface {
	Client() *redis.Client
	Close()
}

func GetRedisClient() *redis.Client {
	return _redisInterface.Client()
}

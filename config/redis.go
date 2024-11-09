package config

import (
	"hello_blockchain/lib/env"
	"time"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

var (
	RedisConn = RedisConfig{
		Addr:     env.GetEnvString("REDIS_HOST", "127.0.0.1"),
		Password: env.GetEnvString("REDIS_PASSWORD", "Redis@123456"),
		DB:       env.GetEnvInt("REDIS_DB", 0),

		DialTimeout:  env.GetEnvDuration("REDIS_DIAL_TIMEOUT", 5) * time.Second,
		ReadTimeout:  env.GetEnvDuration("REDIS_READ_TIMEOUT", 5) * time.Second,
		WriteTimeout: env.GetEnvDuration("REDIS_WRITE_TIMEOUT", 5) * time.Second,
		PoolSize:     env.GetEnvInt("REDIS_POOL_SIZE", 10),
		PoolTimeout:  env.GetEnvDuration("REDIS_POOL_TIMEOUT", 20) * time.Second,
	}
)

package config

import (
	"hello_blockchain/lib/env"
	"time"
)

var (
	HttpProxy   = env.GetEnvString("HTTP_PROXY", "0.0.0.0")
	HttpTimeout = time.Duration(env.GetEnvInt("HTTP_TIMEOUT", 20)) * time.Second
)

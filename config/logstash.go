package config

import "hello_blockchain/lib/env"

var (
	LogstashHost = env.GetEnvString("LOGSTASH_HOST", "127.0.0.1")
)

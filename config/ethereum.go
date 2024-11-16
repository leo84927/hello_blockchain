package config

import "hello_blockchain/lib/env"

var (
	EthereumServerName    = env.GetEnvString("ETHEREUM_SERVER_NAME", "ethereum")
	EthereumServerPort    = env.GetEnvString("ETHEREUM_SERVER_PORT", "10101")
	EthereumServerTimeout = env.GetEnvDuration("ETHEREUM_SERVER_TIMEOUT", 10)
)

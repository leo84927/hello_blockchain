package config

import "hello_blockchain/lib/env"

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
	LogSql   bool
}

var (
	MysqlConn = MysqlConfig{
		Username: env.GetEnvString("DB_USERNAME", "root"),
		Password: env.GetEnvString("DB_PASSWORD", "root"),
		Host:     env.GetEnvString("DB_HOST", "127.0.0.1"),
		Port:     env.GetEnvString("DB_PORT", "3306"),
		Database: env.GetEnvString("DB_DATABASE", "dev"),
		Charset:  "utf8mb4",
		LogSql:   env.GetEnvBool("DB_LOG_SQL", false),
	}
)

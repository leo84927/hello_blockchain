package env

import (
	"time"

	"github.com/rotisserie/eris"
	"github.com/spf13/viper"
)

var (
	env = newEnvInstance()
)

func newEnvInstance() *viper.Viper {
	envInstance := viper.New()

	// 设定档的档名、格式、路径
	envInstance.SetConfigName(".env")
	envInstance.SetConfigType("dotenv")
	envInstance.AddConfigPath(".")

	// 执行读取设定
	err := envInstance.ReadInConfig()
	if err != nil {
		panic(eris.Wrap(err, "load env failed"))
	}

	return envInstance
}

func GetEnvString(key string, defaultValue string) string {
	if env.IsSet(key) {
		return env.GetString(key)
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if env.IsSet(key) {
		return env.GetInt(key)
	}
	return defaultValue
}

func GetEnvInt64(key string, defaultValue int64) int64 {
	if env.IsSet(key) {
		return env.GetInt64(key)
	}
	return defaultValue
}

func GetEnvBool(key string, defaultValue bool) bool {
	if env.IsSet(key) {
		return env.GetBool(key)
	}
	return defaultValue
}

func GetEnvFloat64(key string, defaultValue float64) float64 {
	if env.IsSet(key) {
		return env.GetFloat64(key)
	}
	return defaultValue
}

// GetEnvDuration 以秒为单位(需特别注意单位，返回的值要再乘以秒)
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if env.IsSet(key) {
		return env.GetDuration(key) * time.Second
	}
	return defaultValue * time.Second
}

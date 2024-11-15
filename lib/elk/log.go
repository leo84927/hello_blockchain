package elk

import (
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

func LogError(level logrus.Level, err error, data map[string]interface{}, msgSlice ...string) {
	if err == nil {
		return
	}

	var errorMsg string
	switch len(msgSlice) {
	case 0:
		errorMsg = err.Error()
	case 1:
		errorMsg = msgSlice[0]
	default:
		errorMsg, _ = sonic.MarshalString(msgSlice)
	}

	LogToLogstash(level, err, errorMsg, data)
}

package log

import (
	"fmt"
	"hello_blockchain/helper"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	FilePanic = "panic_"
	FileError = "error_"
)

var (
	loggerMap sync.Map            // 双层的 sync.Map，第一层的key为日期，第二层的key为文件名，第二层的value为 *zerolog.Logger
	logFile   = "./log/%s-%s.log" // 一定要有两个%s 第一个会是文件名，第二个会是日期
)

// withLevel 根据层级记录日志 (目前只记录panic)
func withLevel(level zerolog.Level, skip int, fileName, msg string, data ...interface{}) {
	switch fileName {
	case FilePanic:
		fileName += helper.GetServerName()
	default:
		return
	}

	if len(data) > 1 {
		if jsonString, err := sonic.MarshalString(data); err != nil {
			msg = fmt.Sprintf("logger MarshalString failed: %s, data: %+v", err.Error(), data)
		} else {
			msg += " " + jsonString
		}
	} else if stringData, ok := data[0].(string); len(data) == 1 && ok {
		msg += " " + stringData
	} else {
		// TODO
		msg += " " + fmt.Sprintf("%v", data)
	}

	msg = getMsgPrefix(skip) + msg
	getLogger(fileName).WithLevel(level).Msg(msg)
}

// getMsgPrefix 返回呼叫點的文件字串，ex：tron/main.go:102
func getMsgPrefix(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "runtime.Caller error"
	}

	// 只取呼叫点的文件名和上一层的资料夹
	pathSlice := strings.Split(file, "/")
	l := len(pathSlice)
	if l > 2 {
		file = pathSlice[l-2] + "/" + pathSlice[l-1]
	}

	return file + ":" + strconv.Itoa(line) + " "
}

// getLogger 取得日志实例
func getLogger(fileName string) *zerolog.Logger {

	todayDate := time.Now().Format("2006-01-02")
	yesterdayDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// 删除昨天以前的日志
	loggerMap.Range(func(key, value interface{}) bool {
		// 若key不是今天也不是昨天，从map中删除
		if dateKey, ok := key.(string); ok && dateKey != todayDate && dateKey != yesterdayDate {
			loggerMap.Delete(key)
		}

		return true
	})

	// 判断今天map是否已建立过，若未建立则自动新增
	fileMap, _ := loggerMap.LoadOrStore(todayDate, &sync.Map{})
	logger, ok := fileMap.(*sync.Map).Load(fileName)
	if !ok {
		fileWriter := &lumberjack.Logger{
			Filename:   fmt.Sprintf(logFile, fileName, todayDate),
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		}
		output := zerolog.ConsoleWriter{
			Out:        fileWriter,
			TimeFormat: "[15:04:05.000]",
			NoColor:    true,
		}

		output.FormatMessage = func(i interface{}) string {
			return i.(string)
		}
		output.FormatFieldName = func(i interface{}) string {
			return i.(string) + "="
		}
		output.FormatFieldValue = func(i interface{}) string {
			return i.(string)
		}

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		newLogger := zerolog.New(output).With().Timestamp().Logger()
		logger = &newLogger

		// 存入 fileMap
		fileMap.(*sync.Map).Store(fileName, logger)
	}

	return logger.(*zerolog.Logger)
}

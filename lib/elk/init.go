package elk

import (
	"fmt"
	"hello_blockchain/config"
	"hello_blockchain/connection"
	"hello_blockchain/lib/log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

var (
	_logrus *logrus.Logger
)

type HTTPWriter struct {
	URL    string
	Method string
}

func (w *HTTPWriter) Write(p []byte) (n int, err error) {

	body, code, err := connection.HttpRequest(w.URL, w.Method, p, nil)
	if err != nil {
		log.LogError(log.FileError, err, "Send HttpRequest to logstash failed")
		return 0, err
	}
	if code != fasthttp.StatusOK {
		errMsg := eris.New(string(body))
		log.LogError(log.FileError, errMsg, "Send HttpRequest to logstash status code not OK")
		return 0, errMsg
	}

	return len(p), nil
}

func getLogger() *logrus.Logger {
	return &logrus.Logger{
		Out: &HTTPWriter{
			URL:    config.LogstashHost,
			Method: "PUT",
		},
		Formatter: &logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:03:04",
			DisableHTMLEscape: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				pc, file, line, ok := runtime.Caller(7)
				if !ok {
					filename := filepath.Base(f.File)
					return f.Function, fmt.Sprintf("%s:%d", filename, f.Line)
				}

				// 只取呼叫点的文件名和上一层的资料夹
				pathSlice := strings.Split(file, "/")
				l := len(pathSlice)
				if l > 2 {
					file = pathSlice[l-2] + "/" + pathSlice[l-1]
				}

				return runtime.FuncForPC(pc).Name(), file + ":" + strconv.Itoa(line) + " "
			},
		},
		ReportCaller: true,
		Level:        logrus.WarnLevel,
	}
}

func LogToLogstash(level logrus.Level, err error, msg string, data map[string]interface{}) {
	if _logrus == nil {
		_logrus = getLogger()
	}

	_logrus.WithFields(data).WithError(err).Log(level, msg)
}

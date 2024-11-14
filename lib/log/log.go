package log

import (
	"github.com/bytedance/sonic"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
)

func LogError(fileName string, err error, msgSlice ...string) {
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
	if fileName == FilePanic {
		withLevel(zerolog.PanicLevel, 3, FilePanic, errorMsg, formatErrorToJson(err))
		return
	}

	withLevel(zerolog.ErrorLevel, 3, fileName, errorMsg, formatErrorToJson(err))
}

func formatErrorToJson(err error) string {
	str, marshalErr := sonic.MarshalString(
		eris.ToCustomJSON(
			err,
			eris.NewDefaultJSONFormat(eris.FormatOptions{
				InvertOutput: true, // Flag that inverts the error output (wrap errors shown first).
				WithTrace:    true, // Flag that enables stack trace output.
				InvertTrace:  true, // Flag that inverts the stack trace output (top of call stack shown first).
				WithExternal: true, // Flag that enables external error output.
			}),
		),
	)

	if marshalErr != nil {
		return eris.ToString(err, true)
	}
	return str
}

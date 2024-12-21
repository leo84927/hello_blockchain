package _const

type ErrorCode uint16

const (
	Success = ErrorCode(200)

	ErrorUnknown = ErrorCode(500)
)

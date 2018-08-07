package httpservice

import "fmt"

var (
	// ErrorCtx context key for error
	ErrorCtx = "ERR_CTX"

	// ErrorCode0 正常值
	ErrorCode0 = "0"
	// ErrorCode1 错误值
	ErrorCode1 = "1"

	// ErrorCode9999 系统异常
	ErrorCode9999 = "9999"

	// BizError1001 调用参数异常
	BizError1001 = NewBizError("1001")
	// BizError1002 http请求失败
	BizError1002 = NewBizError("1002")
	// BizError9000 请求鉴权超时
	BizError9000 = NewBizError("9000")
	// BizError9001 请求鉴权非法
	BizError9001 = NewBizError("9001")
)

// StatusError reports an unsuccessful exit by a command.
type StatusError struct {
	Status     string
	StatusCode int
}

func (e StatusError) Error() string {
	return fmt.Sprintf("Status: %s, Code: %d", e.Status, e.StatusCode)
}

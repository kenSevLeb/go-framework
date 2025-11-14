package errors

import (
	"fmt"
	"net/http"
)

// Error
type Error struct {
	// error code
	Code int `json:"code"`
	// error message
	Message string `json:"message"`
	// 隐藏内容
	HideMessage string
}

// Error strings
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ErrorWithHide() string {
	return fmt.Sprintf("%s:%s", e.Message, e.HideMessage)
}

// New
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// 格式化
func Sprintf(code int, format string, v ...interface{}) *Error {
	return New(code, fmt.Sprintf(format, v...))
}

// 附加前置信息
func (e *Error) With(msg string) *Error {
	e.Message = fmt.Sprintf("%s: %s", msg, e.Message)
	return e
}

// 附加后置信息
func (e *Error) Append(msg string) *Error {
	e.Message = fmt.Sprintf("%s: %s", e.Message, msg)
	return e
}

// 隐藏信息
func (e *Error) Hide(v interface{}) *Error {
	e.HideMessage = fmt.Sprintf("%v", v)
	return e
}

// generates a 401 error.
func Unauthorized(format string, v ...interface{}) *Error {
	return New(http.StatusUnauthorized, fmt.Sprintf(format, v...))
}

// generates a 403 error.
func Forbidden(format string, v ...interface{}) *Error {
	return New(http.StatusForbidden, fmt.Sprintf(format, v...))
}

// generates a 404 error.
func NotFound(format string, v ...interface{}) *Error {
	return New(http.StatusNotFound, fmt.Sprintf(format, v...))
}

// generates a 408 error.
func Timeout(format string, v ...interface{}) *Error {
	return New(http.StatusRequestTimeout, fmt.Sprintf(format, v...))
}

// InternalServerError generates a 500 error.
func InternalServer(format string, v ...interface{}) *Error {
	return New(http.StatusInternalServerError, fmt.Sprintf(format, v...))
}

// InvalidParam 参数错误
func InvalidParam(format string, v ...interface{}) *Error {
	return New(1001, fmt.Sprintf(format, v...))
}

// QueryFailed 查询失败
func QueryFailed(format string, v ...interface{}) *Error {
	return New(1002, fmt.Sprintf(format, v...))
}

// SaveFailed 保存失败
func SaveFailed(format string, v ...interface{}) *Error {
	return New(1003, fmt.Sprintf(format, v...))
}

// DeleteFailed 删除失败
func DeleteFailed(format string, v ...interface{}) *Error {
	return New(1004, fmt.Sprintf(format, v...))
}

// Logic 逻辑错误
func Logic(format string, v ...interface{}) *Error {
	return New(1005, fmt.Sprintf(format, v...))
}

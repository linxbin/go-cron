package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = message
	return &Error{Code: code, Message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code, e.Message)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Message, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	newError.Details = append(newError.Details, details...)
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case NotFound.Code:
		return http.StatusNotFound
	case InvalidParams.Code:
		fallthrough
	case LoginError.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}

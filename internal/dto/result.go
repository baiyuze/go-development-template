package dto

import "net/http"

type Result[T any] struct {
	Data  T
	Error error
}

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Err     error  `json:"err,omitempty"`
}

const (
	CodeOK        = 200
	CodeBadReq    = 400
	CodeServerErr = 500
)

func ServiceFail[T any](err error) Result[T] {
	var zero T
	return Result[T]{
		Data:  zero,
		Error: err,
	}
}
func ServiceSuccess[T any](data T) Result[T] {
	return Result[T]{
		Data:  data,
		Error: nil,
	}
}

// Ok 快捷构造函数
func Ok[T any](data T) Response[T] {
	return Response[T]{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func Fail(code int, msg string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: msg,
	}
}

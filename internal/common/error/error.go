package errs

// func
import (
	"errors"
)

type PanicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"err"` // 不暴露内部 error，可选
}

func (e *PanicError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewPanic(code int, msg string, err error) *PanicError {
	return &PanicError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func New(msg string) error {
	return errors.New(msg)
}

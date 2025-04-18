package errs

// func
import "errors"

type PanicError struct {
	Code    int
	Message string
	Err     error
}

func (e *PanicError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
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

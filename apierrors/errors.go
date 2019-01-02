package apierrors

import (
	"errors"
)

// New returns an error that formats as the given text.
func New(err error, status int, values map[string]string) error {
	return &ErrorObject{
		Code:    status,
		Keys:    values,
		Message: err.Error(),
	}
}

// ErrorObject is a trivial implementation of error.
type ErrorObject struct {
	Code    int
	Keys    map[string]string
	Message string
}

func (e *ErrorObject) Error() string {
	return e.Message
}

// Status represents the status code to return from error
func (e *ErrorObject) Status() int {
	return e.Code
}

// Values represents a list of key value pairs to return from error
func (e *ErrorObject) Values() map[string]string {
	return e.Keys
}

// A list of error messages for Dataset API
var (
	ErrCourseNotFound = errors.New("course not found")
	ErrInternalServer = errors.New("internal server error")

	NotFoundMap = map[error]bool{
		ErrCourseNotFound: true,
	}
)

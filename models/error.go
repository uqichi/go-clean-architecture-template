package models

import (
	"fmt"
)

const (
	ErrNotFound  = 1000
	ErrInvalid   = 1001
	ErrDuplicate = 1002
)

var errorText = map[int]string{
	ErrNotFound:  "Record Not Found",
	ErrInvalid:   "Invalid parameter",
	ErrDuplicate: "Duplicate record",
}

// ErrorText returns a text for the application error code. It returns the empty
// string if the code is unknown.
func ErrorText(code int) string {
	return errorText[code]
}

// Error represents an application error.
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// NewError creates a new Error instance.
func NewError(code int, err ...error) *Error {
	he := &Error{Code: code, Message: ErrorText(code)}
	if len(err) > 0 {
		he.Message = err[0].Error()
	}
	return he
}

// Error makes it compatible with `error` interface.
func (e *Error) Error() string {
	return fmt.Sprintf("code=%d, message=%v", e.Code, e.Message)
}

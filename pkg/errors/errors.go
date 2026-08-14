// Package errors defines the fixture error type with a machine-readable code.
package errors

import "fmt"

type Code int

const (
	CodeUnknown Code = iota
	CodeInvalidInput
	CodeNotFound
	CodeInternal
)

type Error struct {
	Code    Code
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("code=%d %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("code=%d %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error { return e.Cause }

func New(code Code, msg string) *Error       { return &Error{Code: code, Message: msg} }
func Wrap(code Code, msg string, cause error) *Error {
	return &Error{Code: code, Message: msg, Cause: cause}
}

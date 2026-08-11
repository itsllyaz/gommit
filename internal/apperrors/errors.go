// Package errors provides typed errors for gommit commands.
package errors

import (
	"errors"
	"fmt"
)

// Code identifies error categories.
type Code string

const (
	CodeNotRepo    Code = "NOT_REPO"
	CodeGitFailure Code = "GIT_FAILURE"
	CodeAIFailure  Code = "AI_FAILURE"
	CodeConfig     Code = "CONFIG"
	CodePlugin     Code = "PLUGIN"
	CodeUsage      Code = "USAGE"
)

// Error is a gommit domain error with code and optional cause.
type Error struct {
	Code    Code
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error { return e.Cause }

// New creates a domain error.
func New(code Code, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Wrap wraps an existing error with code and message.
func Wrap(code Code, message string, cause error) *Error {
	return &Error{Code: code, Message: message, Cause: cause}
}

// IsCode reports whether err matches a gommit error code.
func IsCode(err error, code Code) bool {
	var ge *Error
	if errors.As(err, &ge) {
		return ge.Code == code
	}
	return false
}

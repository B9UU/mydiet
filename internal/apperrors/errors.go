package apperrors

import (
	"fmt"
)

// AppError represents a custom application error
type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

// ErrorCode represents different types of application errors
type ErrorCode int

const (
	// Generic error codes
	ErrInternal ErrorCode = iota
	ErrValidation
	ErrNotFound
	ErrDatabase
	ErrAuthentication
	ErrAuthorization
)

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap allows error unwrapping
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code ErrorCode, message string, err ...error) *AppError {
	appErr := &AppError{
		Code:    code,
		Message: message,
	}
	if len(err) > 0 {
		appErr.Err = err[0]
	}
	return appErr
}

// Wrap creates a new error with additional context
func Wrap(err error, message string) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return &AppError{
			Code:    appErr.Code,
			Message: message,
			Err:     appErr,
		}
	}
	return &AppError{
		Code:    ErrInternal,
		Message: message,
		Err:     err,
	}
}

// Is checks if the error matches a specific error code
func Is(err error, code ErrorCode) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == code
}
package httperrors

import (
	"fmt"
	"net/http"
)

// HTTPError represents a generic HTTP error
type HTTPError interface {
	Error() string
	StatusCode() int
}

type genericError struct {
	code    int
	message string
}

func (e *genericError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}

func (e *genericError) StatusCode() int {
	return e.code
}

// BadRequest creates a new 400 Bad Request error
func BadRequest(message string) HTTPError {
	return &genericError{
		code:    http.StatusBadRequest,
		message: message,
	}
}

// Unauthorized creates a new 401 Unauthorized error
func Unauthorized(message string) HTTPError {
	return &genericError{
		code:    http.StatusUnauthorized,
		message: message,
	}
}

// Forbidden creates a new 403 Forbidden error
func Forbidden(message string) HTTPError {
	return &genericError{
		code:    http.StatusForbidden,
		message: message,
	}
}

// TooManyRequests creates a new 429 Too Many Requests error
func TooManyRequests(message string) HTTPError {
	return &genericError{
		code:    http.StatusTooManyRequests,
		message: message,
	}
}

// InternalServerError creates a new 500 Internal Server Error
func InternalServerError(message string) HTTPError {
	return &genericError{
		code:    http.StatusInternalServerError,
		message: message,
	}
}

// ErrorHandler is a helper function to write HTTPError to http.ResponseWriter
func ErrorHandler(w http.ResponseWriter, err HTTPError) {
	w.WriteHeader(err.StatusCode())
	w.Write([]byte(err.Error()))
}

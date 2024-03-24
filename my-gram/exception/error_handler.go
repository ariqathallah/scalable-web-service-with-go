package exception

import "net/http"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrBadRequest(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ErrInternalServer(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func ErrNotFound(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func ErrUnauthorized(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func ErrForbidden(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

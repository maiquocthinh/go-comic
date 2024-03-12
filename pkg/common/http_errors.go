package common

import (
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"-"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
}

func (e *ApiError) RootError() error {
	if err, ok := e.RootErr.(*ApiError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *ApiError) Error() string {
	return e.RootErr.Error()
}

// return type ApiError with err: origin error, message: message of error, statusCode: http status code
func NewFullApiError(err error, statusCode int, message string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Internal Server Error")
func NewInternalApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Internal Server Error"
	}

	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Bad Request Error")
func NewBadRequestApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Bad Request Error"
	}

	return &ApiError{
		StatusCode: http.StatusBadRequest,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Conflict Error")
func NewConflictApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Conflict Error"
	}

	return &ApiError{
		StatusCode: http.StatusConflict,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "NotFound Resource Error")
func NewNotFoundApiError(err error, message string) *ApiError {
	if message == "" {
		message = "NotFound Resource Error"
	}

	return &ApiError{
		StatusCode: http.StatusNotFound,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Forbidden Error")
func NewForbiddenApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Forbidden Error"
	}

	return &ApiError{
		StatusCode: http.StatusNotFound,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Unauthorized Error")
func NewUnauthorizedApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Unauthorized Error"
	}

	return &ApiError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    err,
		Message:    message,
	}
}

// return type ApiError with err: origin error, message: message of error (default message: "Server Timeout Error")
func NewTimeoutApiError(err error, message string) *ApiError {
	if message == "" {
		message = "Server Timeout Error"
	}

	return &ApiError{
		StatusCode: http.StatusRequestTimeout,
		RootErr:    err,
		Message:    message,
	}
}

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

func NewFullApiError(err error, statusCode int, message string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		RootErr:    err,
		Message:    message,
	}
}

func NewInternalApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		RootErr:    err,
		Message:    message,
	}
}

func NewBadRequestApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		RootErr:    err,
		Message:    message,
	}
}

func NewConflictApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusConflict,
		RootErr:    err,
		Message:    message,
	}
}

func NewNotFoundApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		RootErr:    err,
		Message:    message,
	}
}

func NewForbiddenApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		RootErr:    err,
		Message:    message,
	}
}

func NewUnauthorizedApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    err,
		Message:    message,
	}
}

func NewTimeoutApiError(err error, message string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusRequestTimeout,
		RootErr:    err,
		Message:    message,
	}
}

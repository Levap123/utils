package apperror

import "net/http"

type AppError struct {
	ErrorMessage  string `json:"error_mesage,omitempty"`
	ClientMessage string `json:"client_message,omitempty"`
	Status        int    `json:"status,omitempty"`
}

func NewError(err error, clientMsg string, status int) *AppError {
	return &AppError{
		ErrorMessage:  err.Error(),
		ClientMessage: clientMsg,
		Status:        status,
	}
}

func (e *AppError) Error() string {
	return e.ErrorMessage
}

func MakeBadRequestErr(err error, clientMsg string) *AppError {
	return NewError(err, clientMsg, http.StatusBadRequest)
}

func MakeNotFoundErr(err error, clientMsg string) *AppError {
	return NewError(err, clientMsg, http.StatusNotFound)
}

func MakeUnoauthorizedErr(err error) *AppError {
	return NewError(err, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

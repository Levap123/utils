package apperror

import "net/http"

type AppError struct {
	Err              error  `json:"err,omitempty"`
	ClientMessage    string `json:"client_message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Status           int    `json:"status,omitempty"`
}

func NewError(err error, clientMsg string, status int) *AppError {
	return &AppError{
		Err:              err,
		ClientMessage:    clientMsg,
		DeveloperMessage: err.Error(),
		Status:           status,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
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

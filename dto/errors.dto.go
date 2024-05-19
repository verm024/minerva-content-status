package dto

import "net/http"

type CustomErrorInterface interface {
	Error() string
}

type CustomError struct {
	Message string
	Code    uint
}

func (err *CustomError) Error() string {
	return err.Message
}

func NewCustomError(message string, code uint) *CustomError {
	var computedCode uint = code
	if code == 0 {
		computedCode = http.StatusInternalServerError
	}

	return &CustomError{message, computedCode}
}

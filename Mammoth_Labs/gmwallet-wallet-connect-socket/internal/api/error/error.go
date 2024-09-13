package error

import "gmwallet-connect-go/internal/api/response"

type AppError struct {
	ErrorCode    int
	ErrorMessage response.Message
	ErrorDetails string
}

func (e *AppError) Code() int {
	return e.ErrorCode
}
func (e *AppError) Message() response.Message {
	return e.ErrorMessage
}
func (e *AppError) Error() string {
	return e.ErrorDetails
}

func NewAppError(code int, message response.Message, details string) *AppError {
	return &AppError{
		ErrorCode:    code,
		ErrorMessage: message,
		ErrorDetails: details,
	}
}

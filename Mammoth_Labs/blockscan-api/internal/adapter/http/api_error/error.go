package api_error

import "blockscan-go/internal/core/common/utils"

type ApiError interface {
	Error() string
	Stack() string
	Code() int
	AddStack(stackTrace string)
	SetCode(code int)
}

type ApiCommonError struct {
	StatusCode int             `json:"status_code"`
	Message    utils.ErrorType `json:"message"`
	StackTrace string          `json:"stack_trace"`
}

func NewApiError(code int, message utils.ErrorType, stackTrace string) *ApiCommonError {
	return &ApiCommonError{
		StatusCode: code,
		Message:    message,
		StackTrace: stackTrace,
	}
}

func (e *ApiCommonError) Error() string {
	return string(e.Message)
}

func (e *ApiCommonError) Stack() string {
	return e.StackTrace
}

func (e *ApiCommonError) Code() int {
	return e.StatusCode
}

func (e *ApiCommonError) SetCode(code int) {
	e.StatusCode = code
}
func (e *ApiCommonError) AddStack(stackTrace string) {
	e.StackTrace += "\n" + stackTrace
}

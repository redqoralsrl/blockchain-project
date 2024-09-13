package response

import "blockscan-go/internal/core/common/utils"

type ApiResponse[T any] struct {
	Code    int             `json:"code"`
	Message utils.ErrorType `json:"message"`
	Error   *string         `json:"error"`
	Data    T               `json:"data"`
}

type ApiResponseSwagger[T any] struct {
	Code    int             `json:"code"`
	Message utils.ErrorType `json:"message"`
	Error   *string         `json:"error"`
	Data    T               `json:"data"`
}

func NewApiResponse[T any](code int, message utils.ErrorType, err *string, data T) *ApiResponse[T] {
	return &ApiResponse[T]{
		Code:    code,
		Message: message,
		Error:   err,
		Data:    data,
	}
}

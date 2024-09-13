package response

type Message string

var (
	Success       Message = "success"
	Fail          Message = "fail"
	NotFound      Message = "not found"
	InvalidInput  Message = "invalid input"
	DatabaseError Message = "Database Error"
	InternalError Message = "Internal Server Error"
)

type ApiResponse[T any] struct {
	Code    int     `json:"code"`
	Message Message `json:"message"`
	Error   *string `json:"error"`
	Data    T       `json:"data"`
}

type Response[T any] struct {
	Data *T `json:"data"`
}

func NewApiResponse[T any](code int, message Message, err *string, data T) *ApiResponse[T] {
	return &ApiResponse[T]{
		Code:    code,
		Message: message,
		Error:   err,
		Data:    data,
	}
}

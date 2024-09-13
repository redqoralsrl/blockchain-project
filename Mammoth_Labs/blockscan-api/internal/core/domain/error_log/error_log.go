package error_log

import "time"

type ErrorLog struct {
	ID           int
	Timestamp    time.Time
	IPAddress    string
	UserAgent    string
	Path         string
	HttpMethod   string
	RequestUrl   string
	ErrorCode    int
	ErrorMessage string
	StackTrace   string
}

type Reader interface {
}

type Writer interface {
	Create(e *ErrorLog) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface{}

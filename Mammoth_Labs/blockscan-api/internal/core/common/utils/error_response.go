package utils

import "fmt"

type ErrorType string

const (
	NoDataFound              ErrorType = "No data found"
	SyntaxError              ErrorType = "Syntax"
	DuplicateError           ErrorType = "Duplicate"
	InvalidSQLStatementError ErrorType = "Invalid sql statement"
	BadConnectionError       ErrorType = "Bad connection"
	JsonParseError           ErrorType = "Invalid character"
	ForeignKeyError          ErrorType = "Foreign key"
	CheckViolationError      ErrorType = "Check violation"
	Success                  ErrorType = "Success"
	Fail                     ErrorType = "Fail"
	NotFound                 ErrorType = "Not found"
	InvalidInput             ErrorType = "Invalid input"
	DatabaseError            ErrorType = "Database"
	InternalError            ErrorType = "Internal"
)

type ServiceError struct {
	StackTrace string
	StatusCode int
	Message    string
	ErrorType  ErrorType
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s error %s", e.ErrorType, e.Message)
}

func (e *ServiceError) Code() int {
	return e.StatusCode
}

func (e *ServiceError) Stack() string {
	return e.StackTrace
}

func (e *ServiceError) AddStack(stack string) string {
	e.StackTrace = fmt.Sprintf("%s\n%s", e.StackTrace, stack)
	return e.StackTrace
}

func (e *ServiceError) SetCode(code int) {
	e.StatusCode = code
}

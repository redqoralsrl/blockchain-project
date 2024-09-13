package middleware

import (
	"blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/error_log"
	"blockscan-go/internal/core/domain/error_log/service"
	"errors"
	"github.com/labstack/echo/v4"
	"go/types"
	"net/http"
	"time"
)

func ErrorHandler(errorLogService *service.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {

				var e *echo.HTTPError
				var apiErr *api_error.ApiCommonError
				var serviceErr *utils.ServiceError

				errString := err.Error()
				code := http.StatusInternalServerError
				stack := ""

				switch {
				case errors.As(err, &serviceErr):
					code = serviceErr.Code()
					stack = serviceErr.Stack()
					errString = serviceErr.Error()
				case errors.As(err, &apiErr):
					code = apiErr.Code()
					stack = apiErr.Stack()
					errString = apiErr.Error()
				case errors.As(err, &e):
					code = e.Code
				}

				errorLog := &error_log.ErrorLog{
					Timestamp:    time.Now(),
					IPAddress:    c.RealIP(),
					UserAgent:    c.Request().UserAgent(),
					Path:         c.Path(),
					HttpMethod:   c.Request().Method,
					RequestUrl:   c.Request().URL.String(),
					ErrorCode:    code,
					ErrorMessage: errString,
					StackTrace:   stack,
				}

				if _, ok := errorLogService.Create(errorLog); ok != nil {
					c.Logger().Error(ok)
				}

				return c.JSON(code, response.ApiResponse[*types.Nil]{
					Code:    code,
					Message: utils.Fail,
					Error:   &errString,
					Data:    nil,
				})
			}

			return nil
		}
	}
}

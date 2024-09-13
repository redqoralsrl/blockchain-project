package middleware

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/error_log/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func SetupMiddleware(e *echo.Echo, newLogger *zap.Logger, errorLogService *service.Service) {
	// 각 요청에 고유한 요청 ID를 추가합니다
	e.Use(middleware.RequestID())
	// 제공된 Zap 로거를 사용하여 요청을 기록하는 커스텀 미들웨어
	e.Use(utils.WrapRequestLogger(newLogger))
	// 체인의 어느 곳에서든 발생하는 패닉을 복구하고 중앙 HTTPErrorHandler로 제어를 넘깁니다
	e.Use(middleware.Recover())
	// 기본 설정으로 Cross-Origin Resource Sharing (CORS)을 활성화합니다
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"https://mysite.com", "https://mysite.net"},
	//	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	//}))
	// 각 요청에 대한 최대 지속 시간을 설정합니다. 이 시간이 초과되면 요청이 종료됩니다
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))
	//e.Use(middleware.CSRF())
	// 응답에 보안 관련 헤더를 추가합니다
	e.Use(middleware.Secure())
	// 남용을 방지하기 위해 속도 제한을 추가합니다
	e.Use(utils.NewRateLimiterMiddleware())
	// 요청 본문의 크기를 5MB로 제한합니다
	e.Use(middleware.BodyLimit("5M"))
	// 제공된 오류 로그 서비스를 사용하여 오류를 기록하는 커스텀 오류 처리기
	e.Use(ErrorHandler(errorLogService))
	// 요청 페이로드에 대한 커스텀 검증기를 설정합니다
	e.Validator = utils.NewCustomValidator()
}

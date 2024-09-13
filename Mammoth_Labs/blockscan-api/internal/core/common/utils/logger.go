package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
)

// NewLogger creates and configures a new zap logger instance.
func config(stage string) (*zap.Logger, error) {
	var config zap.Config

	if stage == "dev" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 컬러로 레벨 표시
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 시간 포맷 ISO8601
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder      // 파일명만 표시
		config.DisableStacktrace = true                                     // 스택트레이스 비활성화
		config.InitialFields = map[string]interface{}{
			"appVersion":  "hex-echo-go.v0.0.1",
			"environment": stage,
		} // 모든 로그 메시지에 기본으로 포함될 필드
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // 로그 레벨 설정
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 레벨 표시
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 시간 포맷 ISO8601
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // 파일명만 표시
		config.DisableStacktrace = true                                // 스택트레이스 비활성화
		config.InitialFields = map[string]interface{}{
			"appVersion":  "hex-echo-go.v0.0.1",
			"environment": stage,
		} // 모든 로그 메시지에 기본으로 포함될 필드
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel) // 로그 레벨 설정
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func WrapRequestLogger(logger *zap.Logger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:          true,
		LogStatus:       true,
		LogMethod:       true,
		LogHost:         true,
		LogLatency:      true,
		LogRequestID:    true,
		LogResponseSize: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("request_id", c.Response().Header().Get(echo.HeaderXRequestID)),
				zap.String("method", v.Method),
				zap.String("host", v.Host),
				zap.String("real_ip", c.RealIP()),
				zap.String("response_size", strconv.FormatInt(v.ResponseSize, 10)),
				zap.String("latency", v.Latency.String()),
				zap.Any("header", v.Headers),
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		}})

}

func NewLogger(stage string) (*zap.Logger, error) {
	logger, err := config(stage)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

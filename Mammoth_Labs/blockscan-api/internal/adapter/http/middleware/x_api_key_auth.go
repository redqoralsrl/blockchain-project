package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"blockscan-go/internal/config"
)

func XApiKeyAuth(conf *config.Config) echo.MiddlewareFunc {
	xApiKey := conf.XApiKey

	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-Api-Key",
		Validator: func(auth string, c echo.Context) (bool, error) {
			if len(auth) == 0 {
				return false, errors.New("X-Api-Key is nil")
			} else if auth != xApiKey {
				return false, errors.New("X-Api-Key not matched")
			}

			return auth == xApiKey, nil
		},
	})
}

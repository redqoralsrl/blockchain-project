package crypto_currency

import (
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/crypto_currency"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase crypto_currency.UseCase
}

func CryptoCurrencyHandler(s crypto_currency.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	cryptoCurrencyGroup := e.Group("/api/v1/coin")

	cryptoCurrencyGroup.GET("/price", h.GetPrice, apiAuth)
}

// GetPrice
// @Summary Get coins price
// @ID GetPrice
// @Tags Crypto Currency
// @Description Get / 코인 가격 받아오는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Success 200 {object} response.ApiResponseSwagger[crypto_currency.GetCryptoCurrencyData]
// @Router /coin/price [get]
func (h *handler) GetPrice(c echo.Context) error {
	data, err := h.useCase.Get()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*crypto_currency.GetCryptoCurrencyData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

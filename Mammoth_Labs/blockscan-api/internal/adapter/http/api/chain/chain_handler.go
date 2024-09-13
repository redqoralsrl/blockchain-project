package chain

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/chain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type handler struct {
	useCase chain.UseCase
}

func ChainHandler(s chain.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	chainGroup := e.Group("/api/v1/chain")

	chainGroup.GET("/:chain_id", h.GetChain, apiAuth)
	chainGroup.GET("/token/:chain_id/:contract_address", h.GetToken, apiAuth)
	chainGroup.GET("/tokens/:chain_id", h.GetTokens, apiAuth)
}

// GetChain
// @Summary Get Chain / *공용
// @ID GetChain
// @Tags Chain
// @Description GetChain / chain 관련 정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[chain.Chain]
// @Router /chain/{chain_id} [get]
func (h *handler) GetChain(c echo.Context) error {
	chainStr := c.Param("chain_id")

	chainNum, convertErr := strconv.Atoi(chainStr)
	if convertErr != nil {
		err := apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, convertErr.Error())
		return err
	}

	data, err := h.useCase.Get(chainNum)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*chain.Chain](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

// GetToken
// @Summary Get Token / *공용
// @ID GetToken
// @Tags Chain
// @Description GetToken / token 관련 정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Param contract_address path string true "컨트랙트 주소"
// @Success 200 {object} response.ApiResponseSwagger[chain.Chain]
// @Router /chain/token/{chain_id}/{contract_address} [get]
func (h *handler) GetToken(c echo.Context) error {
	var input chain.GetTokenChainInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if input.ContractAddress == "0x" || input.ContractAddress == "" {
		err := apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, "contract_address invalid")
		return err
	}

	data, err := h.useCase.GetToken(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*chain.Chain](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

// GetTokens
// @Summary Get Tokens / *공용
// @ID GetTokens
// @Tags Chain
// @Description GetTokens / token 관련 정보들을 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[[]chain.Chain]
// @Router /chain/tokens/{chain_id} [get]
func (h *handler) GetTokens(c echo.Context) error {
	var input chain.GetTokensChainInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	list, err := h.useCase.GetTokens(&input)
	if err != nil {
		return err
	}

	if list == nil {
		list = []*chain.Chain{}
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*chain.Chain](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

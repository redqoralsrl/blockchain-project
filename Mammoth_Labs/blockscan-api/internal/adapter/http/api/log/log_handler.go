package log

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase log.UseCase
}

func LogHandler(s log.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	logGroup := e.Group("/api/v1/log")

	logGroup.GET("/nft/activities/:wallet_address", h.LogGetNFTsByWallet, apiAuth)
	logGroup.POST("/nft/activities/:wallet_address", h.LogGetSearchNFTsByWallet, apiAuth)
}

// LogGetNFTsByWallet
// @Summary Get NFTs by wallet / *공용
// @ID LogGetNFTsByWallet
// @Tags Log
// @Description GetByWallet / 해당 지갑에서 발생한 NFT 로그 기록 가져오는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param wallet_address path string true "월렛주소"
// @Param chain_id query int true "블록체인 체인넘버"
// @Param take query int true "가져올 아이템 수 최대 100"
// @Param skip query int true "건너뛸 아이템 수 (옵션) 기본 0"
// @Success 200 {object} response.ApiResponseSwagger[[]log.GetLogNFTsByWalletData]
// @Router /log/nft/activities/{wallet_address} [get]
func (h *handler) LogGetNFTsByWallet(c echo.Context) error {
	var input log.GetLogNFTsByWalletInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if input.Take == 0 {
		input.Take = 10
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := utils.ValidateWalletAddress(input.WalletAddress); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	data, err := h.useCase.GetNFTsByWallet(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*log.GetLogNFTsByWalletData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

// LogGetSearchNFTsByWallet
// @Summary Get NFT logs by wallet / *공용
// @ID LogGetSearchNFTsByWallet
// @Tags Log
// @Description LogGetSearchNFTsByWallet / 해당 지갑에서 발생한 NFT 로그 날짜, 종류에 따라 가져오는 API / *공용(gmwallet, edem)
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param wallet_address path string true "월렛주소"
// @Param body body log.GetSearchByWalletInput true "Request body"
// @Success 200 {object} response.ApiResponseSwagger[[]log.GetLogNFTsByWalletData]
// @Router /log/nft/activities/{wallet_address} [post]
func (h *handler) LogGetSearchNFTsByWallet(c echo.Context) error {
	var input log.GetSearchNFTsByWalletInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if input.Take == 0 {
		input.Take = 10
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := utils.ValidateWalletAddress(input.WalletAddress); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	data, err := h.useCase.GetSearchNFTsByWallet(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*log.GetLogNFTsByWalletData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

package giant_mammoth

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/blockchain/gmmt/staking"
	"blockscan-go/internal/core/blockchain/gmmt/swap"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/giant_mammoth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type handler struct {
	useCase giant_mammoth.UseCase
}

func GiantMammothHandler(s giant_mammoth.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	giantMammothGroup := e.Group("/api/v1/giant_mammoth")

	giantMammothGroup.GET("/staking/list/:chain_id", h.GetStakingList, apiAuth)
	giantMammothGroup.GET("/staking/info/:chain_id/:wallet_address", h.GetStakingByAccount, apiAuth)
	giantMammothGroup.GET("/swap/list/:chain_id", h.GetSwapPairList, apiAuth)
}

// GetStakingList
// @Summary Get Staking list / *공용
// @ID GetStakingList
// @Tags GiantMammoth
// @Description Get Staking list / Gmmt staking 정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[staking.Staking]
// @Router /giant_mammoth/staking/list/{chain_id} [get]
func (h *handler) GetStakingList(c echo.Context) error {
	chainStr := c.Param("chain_id")

	chainNum, convertErr := strconv.Atoi(chainStr)
	if convertErr != nil {
		err := apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, convertErr.Error())
		return err
	}

	data, err := h.useCase.GetStakingList(chainNum)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*staking.Staking](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

// GetStakingByAccount
// @Summary Get Staking By Account / *공용
// @ID GetStakingByAccount
// @Tags GiantMammoth
// @Description Get Staking By Account / 해당 지갑에 관련된 staking 정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Param wallet_address path string true "지갑주소"
// @Success 200 {object} response.ApiResponseSwagger[[]staking.ValidatorByAccountData]
// @Router /giant_mammoth/staking/info/{chain_id}/{wallet_address} [get]
func (h *handler) GetStakingByAccount(c echo.Context) error {
	var input giant_mammoth.GetStakingByAccountInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := utils.ValidateWalletAddress(input.WalletAddress); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	var list []*staking.ValidatorByAccountData
	list, err := h.useCase.GetStakingByAccount(input.WalletAddress, input.ChainId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*staking.ValidatorByAccountData](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

// GetSwapPairList
// @Summary Get Swap Pair List / *공용
// @ID GetSwapPairList
// @Tags GiantMammoth
// @Description Get Swap Pair List / 아이보리 스왑 pair list를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param chain_id path int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[[]swap.Swap]
// @Router /giant_mammoth/swap/list/{chain_id} [get]
func (h *handler) GetSwapPairList(c echo.Context) error {
	chainStr := c.Param("chain_id")

	chainNum, convertErr := strconv.Atoi(chainStr)
	if convertErr != nil {
		err := apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, convertErr.Error())
		return err
	}

	data, err := h.useCase.GetSwapPairList(chainNum)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*swap.Swap](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

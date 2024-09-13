package nft

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/nft"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase nft.UseCase
}

func NftHandler(s nft.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	nftGroup := e.Group("/api/v1/nft")

	nftGroup.GET("/:wallet_address", h.NftGet, apiAuth)
	nftGroup.GET("/detail/:hash/:token_id", h.NftGetDetail, apiAuth)
	nftGroup.GET("/detail/:hash/:token_id/:wallet_address", h.NftGetDetailOfWallet, apiAuth)
}

// NftGet
// @Summary  Get Nfts by wallet / *공용
// @ID NftGet
// @Tags Nft
// @Description NftGet / 월렛이 보유한 NFT들을 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param wallet_address path string true "지갑주소"
// @Param chain_id query int true "블록체인 체인넘버"
// @Param take query int true "가져올 아이템 수 최대 100"
// @Param skip query int true "건너뛸 아이템 수 (옵션) 기본 0"
// @Success 200 {object} response.ApiResponseSwagger[[]nft.GetNftData]
// @Router /nft/{wallet_address} [get]
func (h *handler) NftGet(c echo.Context) error {
	var input nft.GetNftInput

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

	var list []*nft.GetNftData
	list, err := h.useCase.Get(&input)
	if err != nil {
		return err
	}

	if list == nil {
		list = []*nft.GetNftData{}
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*nft.GetNftData](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

// NftGetDetail
// @Summary  Get Nft Detail / *공용
// @ID NftGetDetail
// @Tags Nft
// @Description NftGetDetail / NFT 상세정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param hash path string true "컨트랙트 주소"
// @Param token_id path string true "토큰 아이디"
// @Param chain_id query int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[nft.GetNFtDetailData]
// @Router /nft/detail/{hash}/{token_id} [get]
func (h *handler) NftGetDetail(c echo.Context) error {
	var input nft.GetNftDetailInput

	if err := c.Bind(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		err = apiError.NewApiError(http.StatusBadRequest, utils.InvalidInput, err.Error())
		return err
	}

	data, err := h.useCase.GetDetail(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*nft.GetNFtDetailData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

// NftGetDetailOfWallet
// @Summary  Get Nft Detail of Wallet / *공용
// @ID NftGetDetailOfWallet
// @Tags Nft
// @Description NftGetDetailOfWallet / 지갑이 보유한 NFT 상세정보를 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param hash path string true "컨트랙트 주소"
// @Param token_id path string true "토큰 아이디"
// @Param wallet_address path string true "지갑주소"
// @Param chain_id query int true "블록체인 체인넘버"
// @Success 200 {object} response.ApiResponseSwagger[nft.GetNFtDetailData]
// @Router /nft/detail/{hash}/{token_id}/{wallet_address} [get]
func (h *handler) NftGetDetailOfWallet(c echo.Context) error {
	var input nft.GetNftDetailOfWalletInput

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

	data, err := h.useCase.GetDetailOfWallet(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*nft.GetNFtDetailData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

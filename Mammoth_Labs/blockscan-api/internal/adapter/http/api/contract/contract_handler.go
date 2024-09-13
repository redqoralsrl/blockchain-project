package contract

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/contract"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase contract.UseCase
}

func ContractHandler(s contract.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	contractGroup := e.Group("/api/v1/contract")

	contractGroup.GET("/wallet/nft/:wallet_address", h.ContractGetWalletNFTsByCollection, apiAuth)
	contractGroup.GET("/:hash/wallet/nft/:wallet_address", h.ContractGetCollectionNFTsForWallet, apiAuth)
}

// ContractGetWalletNFTsByCollection
// @Summary  Get wallet nfts by collection / *공용
// @ID ContractGetWalletNFTsByCollection
// @Tags Contract
// @Description ContractGetWalletNFTsByCollection / 월렛이 보유한 NFT들을 컬렉션 별로 분류해서 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param wallet_address path string true "지갑주소"
// @Param chain_id query int true "블록체인 체인넘버"
// @Param take query int true "가져올 아이템 수 최대 100"
// @Param skip query int true "건너뛸 아이템 수 (옵션) 기본 0"
// @Success 200 {object} response.ApiResponseSwagger[[]contract.GetContractWalletNFTsByCollectionData]
// @Router /contract/wallet/nft/{wallet_address} [get]
func (h *handler) ContractGetWalletNFTsByCollection(c echo.Context) error {
	var input contract.GetContractWalletNFTsByCollectionInput

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

	var list []*contract.GetContractWalletNFTsByCollectionData
	list, err := h.useCase.GetWalletNFTsByCollection(&input)
	if err != nil {
		return err
	}

	if list == nil {
		list = []*contract.GetContractWalletNFTsByCollectionData{}
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*contract.GetContractWalletNFTsByCollectionData](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

// ContractGetCollectionNFTsForWallet
// @Summary Get collection nfts for wallet / *공용
// @ID ContractGetCollectionNFTsForWallet
// @Tags Contract
// @Description ContractGetCollectionNFTsForWallet / 해당 컨트랙트에 월렛이 보유한 NFT들을 반환하는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param hash path string true "컨트랙트 주소"
// @Param wallet_address path string true "지갑주소"
// @Param chain_id query int true "블록체인 체인넘버"
// @Param take query int true "가져올 아이템 수 최대 100"
// @Param skip query int true "건너뛸 아이템 수 (옵션) 기본 0"
// @Success 200 {object} response.ApiResponseSwagger[contract.GetContractCollectionNFTsForWalletData]
// @Router /contract/{hash}/wallet/nft/{wallet_address} [get]
func (h *handler) ContractGetCollectionNFTsForWallet(c echo.Context) error {
	var input contract.GetContractCollectionNFTsForWalletInput

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

	var data *contract.GetContractCollectionNFTsForWalletData
	data, err := h.useCase.GetCollectionNFTsForWallet(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[*contract.GetContractCollectionNFTsForWalletData](
		http.StatusOK,
		utils.Success,
		nil,
		data,
	))
}

package transaction

import (
	apiError "blockscan-go/internal/adapter/http/api_error"
	"blockscan-go/internal/adapter/http/response"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/transaction"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase transaction.UseCase
}

func TransactionHandler(s transaction.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	h := handler{s}

	transactionGroup := e.Group("/api/v1/transaction")

	transactionGroup.POST("", h.GetTransaction, apiAuth)
	transactionGroup.GET("/recent/:wallet_address", h.GetAllTransaction, apiAuth)
}

// GetTransaction
// @Summary get coin, token transactions / *공용
// @ID GetTransaction
// @Tags Transaction
// @Description GetTransaction / 해당 지갑에서 코인/토큰 전송내역 가져오는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param body body transaction.GetTransactionInput true "Request body"
// @Success 200 {object} response.ApiResponseSwagger[[]transaction.GetTransactionData]
// @Router /transaction [post]
func (h *handler) GetTransaction(c echo.Context) error {
	var input transaction.GetTransactionInput

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

	var list []*transaction.GetTransactionData
	list, err := h.useCase.Get(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*transaction.GetTransactionData](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

// GetAllTransaction
// @Summary get coin,token recent transactions / *공용
// @ID GetAllTransaction
// @Tags Transaction
// @Description GetAllTransaction / 해당 지갑에서 코인/토큰 최근 전송내역 가져오는 API / *공용
// @Produce json
// @Param X-Api-Key header string true "x api key"
// @Param wallet_address path string true "지갑주소"
// @Param chain_id query int true "블록체인 넘버"
// @Param take query int true "가져올 아이템 수 최대 100"
// @Param skip query int true "건너뛸 아이템 수 (옵션) 기본 0"
// @Success 200 {object} response.ApiResponseSwagger[[]transaction.GetTransactionData]
// @Router /transaction/recent/{wallet_address} [get]
func (h *handler) GetAllTransaction(c echo.Context) error {
	var input transaction.GetAllTransactionInput

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

	var list []*transaction.GetTransactionData
	list, err := h.useCase.GetAll(&input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewApiResponse[[]*transaction.GetTransactionData](
		http.StatusOK,
		utils.Success,
		nil,
		list,
	))
}

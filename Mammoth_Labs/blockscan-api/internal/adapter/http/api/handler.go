package api

import (
	chainHandler "blockscan-go/internal/adapter/http/api/chain"
	contractHandler "blockscan-go/internal/adapter/http/api/contract"
	cryptoCurrencyHandler "blockscan-go/internal/adapter/http/api/crypto_currency"
	erc721Handler "blockscan-go/internal/adapter/http/api/erc721"
	giantMammothHandler "blockscan-go/internal/adapter/http/api/giant_mammoth"
	logHandler "blockscan-go/internal/adapter/http/api/log"
	nftHandler "blockscan-go/internal/adapter/http/api/nft"
	transactionHandler "blockscan-go/internal/adapter/http/api/transaction"
	chainUseCase "blockscan-go/internal/core/domain/chain"
	contractUseCase "blockscan-go/internal/core/domain/contract"
	cryptoCurrencyUseCase "blockscan-go/internal/core/domain/crypto_currency"
	erc721UseCase "blockscan-go/internal/core/domain/erc721"
	giantMammothUseCase "blockscan-go/internal/core/domain/giant_mammoth"
	logUseCase "blockscan-go/internal/core/domain/log"
	nftUseCase "blockscan-go/internal/core/domain/nft"
	transactionUseCase "blockscan-go/internal/core/domain/transaction"
	"github.com/labstack/echo/v4"
)

type Services struct {
	ContractService       contractUseCase.UseCase
	Erc721Service         erc721UseCase.UseCase
	LogService            logUseCase.UseCase
	CryptoCurrencyService cryptoCurrencyUseCase.UseCase
	ChainService          chainUseCase.UseCase
	TransactionService    transactionUseCase.UseCase
	GiantMammothService   giantMammothUseCase.UseCase
	NftService            nftUseCase.UseCase
}

func Handlers(services *Services, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	contractHandler.ContractHandler(services.ContractService, e, apiAuth)
	erc721Handler.Erc721Handler(services.Erc721Service, e, apiAuth)
	logHandler.LogHandler(services.LogService, e, apiAuth)
	cryptoCurrencyHandler.CryptoCurrencyHandler(services.CryptoCurrencyService, e, apiAuth)
	chainHandler.ChainHandler(services.ChainService, e, apiAuth)
	transactionHandler.TransactionHandler(services.TransactionService, e, apiAuth)
	giantMammothHandler.GiantMammothHandler(services.GiantMammothService, e, apiAuth)
	nftHandler.NftHandler(services.NftService, e, apiAuth)
}

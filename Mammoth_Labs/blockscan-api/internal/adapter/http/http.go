package http

import (
	"blockscan-go/internal/adapter/http/api"
	_ "blockscan-go/internal/adapter/http/docs"
	"blockscan-go/internal/adapter/http/middleware"
	"blockscan-go/internal/config"
	stakingCaller "blockscan-go/internal/core/blockchain/gmmt/staking"
	swapCaller "blockscan-go/internal/core/blockchain/gmmt/swap"
	ivoryrouter "blockscan-go/internal/core/common/abi/Router"
	"blockscan-go/internal/core/common/abi/factory"
	"blockscan-go/internal/core/common/abi/pair"
	"blockscan-go/internal/core/common/abi/staking"
	chainPostgresql "blockscan-go/internal/core/domain/chain/postgresql"
	chainService "blockscan-go/internal/core/domain/chain/service"
	contractPostgresql "blockscan-go/internal/core/domain/contract/postgresql"
	contractService "blockscan-go/internal/core/domain/contract/service"
	cryptoCurrencyPostgresql "blockscan-go/internal/core/domain/crypto_currency/postgresql"
	cryptoCurrencyService "blockscan-go/internal/core/domain/crypto_currency/service"
	erc721Postgresql "blockscan-go/internal/core/domain/erc721/postgresql"
	erc721Service "blockscan-go/internal/core/domain/erc721/service"
	errorLogPostgresql "blockscan-go/internal/core/domain/error_log/postgresql"
	errorLogService "blockscan-go/internal/core/domain/error_log/service"
	giantMammothService "blockscan-go/internal/core/domain/giant_mammoth/service"
	logPostgresql "blockscan-go/internal/core/domain/log/postgresql"
	logService "blockscan-go/internal/core/domain/log/service"
	nftPostgresql "blockscan-go/internal/core/domain/nft/postgresql"
	nftService "blockscan-go/internal/core/domain/nft/service"
	transactionPostgresql "blockscan-go/internal/core/domain/transaction/postgresql"
	transactionService "blockscan-go/internal/core/domain/transaction/service"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
)

// @title blockscan-go API
// @version 1.0
// @description API Backend Template

// @license.name Go Echo
// @host tracking.gmmtchain.io
// @BasePath /api/v1
func Run(db *sql.DB, logger *zap.Logger, transactionManager *postgresql.Manager, conf *config.Config) {
	e := echo.New()

	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// setup storage
	newErrorLogStorage := errorLogPostgresql.NewErrorLogStorage(db)
	newContractStorage := contractPostgresql.NewContractStorage(db)
	newErc721Storage := erc721Postgresql.NewErc721Storage(db)
	newLogStorage := logPostgresql.NewLogStorage(db)
	newCryptoCurrency := cryptoCurrencyPostgresql.NewCryptoCurrencyStorage(db)
	newChain := chainPostgresql.NewChainStorage(db)
	newTransaction := transactionPostgresql.NewTransactionStorage(db)
	newNft := nftPostgresql.NewNftStorage(db)

	// errorLog service
	newErrorLogService := errorLogService.NewService(newErrorLogStorage, logger)

	// setup middleware
	middleware.SetupMiddleware(e, logger, newErrorLogService)

	// Api Key
	XApiKeyAuth := middleware.XApiKeyAuth(conf)

	// ip country code adapter
	//ipCountryCodeAdapter := port.NewIpCountryCodeAdapter()

	// blockchain setting
	// staking
	stakingAbi, _ := abi.JSON(strings.NewReader(string(staking.StakingMetaData.ABI)))
	// factory
	factoryAddress := &swapCaller.NetworkOfCA{
		TestNetworkAddress: "0x0e9e740319A4A2f4ea89Af6fd3A8B8016D6fe7e9",
		MainNetworkAddress: "0x6aD630595ADC6717119aB5c8192e1CEd94E0C587",
	}
	// router
	routerAddress := &swapCaller.NetworkOfCA{
		TestNetworkAddress: "0xd78eC64aFB1273d80ae947fF0797830828975c9D",
		MainNetworkAddress: "0x92d8bF464931aab6323dab18d56bBb37e119DE53",
	}
	factoryAbi, _ := abi.JSON(strings.NewReader(string(factory.FactoryMetaData.ABI)))
	routerAbi, _ := abi.JSON(strings.NewReader(string(ivoryrouter.IvoryrouterMetaData.ABI)))
	pairAbi, _ := abi.JSON(strings.NewReader(string(pair.PairMetaData.ABI)))

	newStaking := stakingCaller.NewCallers(conf, &stakingAbi, logger)
	newFactory := swapCaller.NewCallers(conf, &factoryAbi, &routerAbi, &pairAbi, factoryAddress, routerAddress, logger)

	// Jwt

	// service
	services := api.Services{
		ContractService:       contractService.NewService(newContractStorage, transactionManager, logger),
		Erc721Service:         erc721Service.NewService(newErc721Storage, transactionManager, logger),
		LogService:            logService.NewService(newLogStorage, transactionManager, logger),
		CryptoCurrencyService: cryptoCurrencyService.NewService(newCryptoCurrency, transactionManager, logger),
		ChainService:          chainService.NewService(newChain, transactionManager, logger),
		TransactionService:    transactionService.NewService(newTransaction, transactionManager, logger),
		GiantMammothService:   giantMammothService.NewService(newStaking, newFactory, logger),
		NftService:            nftService.NewService(newNft, transactionManager, logger),
	}

	api.Handlers(&services, e, XApiKeyAuth)

	// api
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hex-echo-go backend API")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "server is running")
	})

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler, echoMiddleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "sid" && password == "3725" {
			return true, nil
		}
		return false, nil
	}))

	if err := e.Start(":" + conf.ApiPort); err != nil {
		log.Fatalf("Error running API: %v", err)
	}
}

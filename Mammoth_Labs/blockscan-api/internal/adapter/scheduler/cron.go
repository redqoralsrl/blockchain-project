package scheduler

import (
	"blockscan-go/internal/adapter/scheduler/coin_price"
	"blockscan-go/internal/adapter/scheduler/nft"
	"blockscan-go/internal/config"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"

	attributesPostgresql "blockscan-go/internal/core/domain/attributes/postgresql"
	attributesService "blockscan-go/internal/core/domain/attributes/service"
	cryptoCurrencyPostgresql "blockscan-go/internal/core/domain/crypto_currency/postgresql"
	cryptoCurrencyService "blockscan-go/internal/core/domain/crypto_currency/service"
	erc1155Postgresql "blockscan-go/internal/core/domain/erc1155/postgresql"
	erc1155Service "blockscan-go/internal/core/domain/erc1155/service"
	erc721Postgresql "blockscan-go/internal/core/domain/erc721/postgresql"
	erc721Service "blockscan-go/internal/core/domain/erc721/service"
)

type CronService struct {
	cron *cron.Cron
	once sync.Once
}

func NewCronService() *CronService {
	return &CronService{}
}

func (cs *CronService) GetCronInstance() *cron.Cron {
	cs.once.Do(func() {
		cs.cron = cron.New()
		cs.cron.Start()
	})
	return cs.cron
}

func (cs *CronService) Stop() {
	if cs.cron != nil {
		cs.cron.Stop()
	}
}

func Run(db *sql.DB, tx *postgresql.Manager, logger *zap.Logger, conf *config.Config) (*cron.Cron, *nft.NftImageParser, *coin_price.CoinPrice) {
	// Cron Setting
	cronService := NewCronService()
	cronInstance := cronService.GetCronInstance()
	cronInstance.Start()

	// DB
	newErc721Service := erc721Service.NewService(erc721Postgresql.NewErc721Storage(db), tx, logger)
	newAttributesService := attributesService.NewService(attributesPostgresql.NewAttributesStorage(db), tx, logger)
	newCryptoCurrencyService := cryptoCurrencyService.NewService(cryptoCurrencyPostgresql.NewCryptoCurrencyStorage(db), tx, logger)
	newErc1155Service := erc1155Service.NewService(erc1155Postgresql.NewErc1155Storage(db), tx, logger)

	nftImageParser := nft.NewNftImageParser(newErc721Service, newErc1155Service, newAttributesService, tx, cronInstance, logger)
	coinPrice := coin_price.NewCoinPrice(newCryptoCurrencyService, tx, cronInstance, logger, conf)

	return cronInstance, nftImageParser, coinPrice
}

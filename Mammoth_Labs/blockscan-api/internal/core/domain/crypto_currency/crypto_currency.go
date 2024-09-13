package crypto_currency

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
)

type CryptoCurrency struct {
	Id           int     `json:"id"`
	Timestamp    int     `json:"timestamp"`
	MorningDate  int     `json:"morning_date"`
	MidnightDate int     `json:"midnight_date"`
	LastUpdated  int     `json:"last_updated"`
	EthPrice     float64 `json:"eth_price"`
	MmtPrice     float64 `json:"mmt_price"`
	GmmtPrice    float64 `json:"gmmt_price"`
	MaticPrice   float64 `json:"matic_price"`
	BnbPrice     float64 `json:"bnb_price"`
}

type Reader interface {
	Get() (*GetCryptoCurrencyData, error)
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateCryptoCurrencyInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get() (*GetCryptoCurrencyData, *utils.ServiceError)
}

type CronUseCase interface {
	Create(queryRower postgresql.Query, input *CreateCryptoCurrencyInput) *utils.ServiceError
}

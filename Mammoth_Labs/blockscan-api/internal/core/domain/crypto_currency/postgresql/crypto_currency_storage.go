package postgresql

import (
	"blockscan-go/internal/core/domain/crypto_currency"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
)

type CryptoCurrencyStorage struct {
	db *sql.DB
}

func NewCryptoCurrencyStorage(db *sql.DB) *CryptoCurrencyStorage {
	return &CryptoCurrencyStorage{db}
}

func (s *CryptoCurrencyStorage) Create(queryRower postgresql.Query, input *crypto_currency.CreateCryptoCurrencyInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into crypto_currency (timestamp, morning_date, midnight_date, last_updated, eth_price, mmt_price, gmmt_price, matic_price, bnb_price)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		returning id;
	`

	ethPriceStr := fmt.Sprintf("%f", input.EthPrice)
	mmtPriceStr := fmt.Sprintf("%f", input.MmtPrice)
	gmmtPriceStr := fmt.Sprintf("%f", input.GmmtPrice)
	maticPriceStr := fmt.Sprintf("%f", input.MaticPrice)
	bnbPriceStr := fmt.Sprintf("%f", input.BnbPrice)

	_, err := queryRower.Exec(
		query,
		input.Timestamp,
		input.MorningDate,
		input.MidnightDate,
		input.LastUpdated,
		ethPriceStr,
		mmtPriceStr,
		gmmtPriceStr,
		maticPriceStr,
		bnbPriceStr,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CryptoCurrencyStorage) Get() (*crypto_currency.GetCryptoCurrencyData, error) {
	query := `
		select timestamp, eth_price, mmt_price, gmmt_price, matic_price, bnb_price
		from crypto_currency
		order by timestamp desc
		limit 1;
	`

	var cryptoData = &crypto_currency.GetCryptoCurrencyData{}

	err := s.db.QueryRow(
		query,
	).Scan(
		&cryptoData.Timestamp,
		&cryptoData.EthPrice,
		&cryptoData.MmtPrice,
		&cryptoData.GmmtPrice,
		&cryptoData.MaticPrice,
		&cryptoData.BnbPrice,
	)

	if err != nil {
		return nil, err
	}

	return cryptoData, nil
}

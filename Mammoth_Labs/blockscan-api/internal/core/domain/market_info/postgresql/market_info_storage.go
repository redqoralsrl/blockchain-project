package postgresql

import (
	"blockscan-go/internal/core/domain/market_info"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
	"math/big"
)

type MarketInfoStorage struct {
	db *sql.DB
}

func NewMarketInfoStorage(db *sql.DB) *MarketInfoStorage {
	return &MarketInfoStorage{db}
}

func (s *MarketInfoStorage) Create(queryRower postgresql.Query, input *market_info.CreateMarketInfoInput) (*market_info.MarketInfo, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into market_info (chain_id, log_id, transaction_hash, collection, seller, buyer, volume, volume_symbol, volume_contract_address)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		returning id, chain_id, log_id, transaction_hash, collection, seller, buyer, volume, volume_symbol, volume_contract_address;
	`

	volumeStr := input.Volume.String()
	var getVolume string

	marketInfoData := &market_info.MarketInfo{}

	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.LogId,
		input.TransactionHash,
		input.Collection,
		input.Seller,
		input.Buyer,
		volumeStr,
		input.VolumeSymbol,
		input.VolumeContractAddress,
	).Scan(
		&marketInfoData.ID,
		&marketInfoData.ChainId,
		&marketInfoData.LogId,
		&marketInfoData.TransactionHash,
		&marketInfoData.Collection,
		&marketInfoData.Seller,
		&marketInfoData.Buyer,
		&getVolume,
		&marketInfoData.VolumeSymbol,
		&marketInfoData.VolumeContractAddress,
	)

	if err != nil {
		return nil, err
	}

	if getVolume != "" {
		volumeInt := new(big.Int)
		volumeInt, ok := volumeInt.SetString(getVolume, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert number string to big.Int")
		}

		marketInfoData.Volume = *volumeInt
	}

	return marketInfoData, err
}

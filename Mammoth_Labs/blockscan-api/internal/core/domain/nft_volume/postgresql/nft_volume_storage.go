package postgresql

import (
	"blockscan-go/internal/core/domain/nft_volume"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
	"math/big"
)

type NftVolumeStorage struct {
	db *sql.DB
}

func NewNftVolumeStorage(db *sql.DB) *NftVolumeStorage {
	return &NftVolumeStorage{db}
}

func (s *NftVolumeStorage) Create(queryRower postgresql.Query, input *nft_volume.CreateNftVolumeInput) (*nft_volume.NftVolume, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into nft_volume (chain_id, log_id, from_address, to_address, value, contract, symbol, timestamp, transaction_hash, event)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning id, chain_id, log_id, from_address, to_address, value, contract, symbol, timestamp, transaction_hash, event;
	`

	valueStr := input.Value.String()
	var getValue string

	nftVolumeData := &nft_volume.NftVolume{}
	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.LogId,
		input.FromAddress,
		input.ToAddress,
		valueStr,
		input.Contract,
		input.Symbol,
		input.Timestamp,
		input.TransactionHash,
		input.Event,
	).Scan(
		&nftVolumeData.ID,
		&nftVolumeData.ChainId,
		&nftVolumeData.LogId,
		&nftVolumeData.FromAddress,
		&nftVolumeData.ToAddress,
		&getValue,
		&nftVolumeData.Contract,
		&nftVolumeData.Symbol,
		&nftVolumeData.Timestamp,
		&nftVolumeData.TransactionHash,
		&nftVolumeData.Event,
	)

	if getValue != "" {
		nftVolumeInt := new(big.Int)
		nftVolumeInt, ok := nftVolumeInt.SetString(getValue, 10)
		if !ok {
			return nil, fmt.Errorf("nftVolumeInt failed to convert number string to big.Int")
		}
		nftVolumeData.Value = *nftVolumeInt
	}

	if err != nil {
		return nil, err
	}

	return nftVolumeData, nil
}

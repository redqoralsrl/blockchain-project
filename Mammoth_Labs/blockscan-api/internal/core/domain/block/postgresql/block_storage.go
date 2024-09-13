package postgresql

import (
	"blockscan-go/internal/core/domain/block"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
	"math/big"
)

type BlockStorage struct {
	db *sql.DB
}

func NewBlockStorage(db *sql.DB) *BlockStorage {
	return &BlockStorage{db}
}

func (s *BlockStorage) Create(queryRower postgresql.Query, input *block.CreateBlockInput) (*block.Block, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into block (chain_id, difficulty, extra_data, gas_limit, gas_used, hash, logs_bloom, miner, mix_hash, nonce, number_hex, number, parent_hash, receipts_root, sha3_uncles, size, state_root, timestamp, total_difficulty, transactions_root, uncles)
		values ($1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
		returning id, chain_id, difficulty, extra_data, gas_limit, gas_used, hash, logs_bloom, miner, mix_hash, nonce, number_hex, number, parent_hash, receipts_root, sha3_uncles, size, state_root, timestamp, total_difficulty, transactions_root, uncles;
	`

	numberStr := input.Number.String()
	var getNumberStr string

	blockData := &block.Block{}
	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.Difficulty,
		input.ExtraData,
		input.GasLimit,
		input.GasUsed,
		input.Hash,
		input.LogsBloom,
		input.Miner,
		input.MixHash,
		input.Nonce,
		input.NumberHex,
		numberStr,
		input.ParentHash,
		input.ReceiptsRoot,
		input.Sha3Uncles,
		input.Size,
		input.StateRoot,
		input.Timestamp,
		input.TotalDifficulty,
		input.TransactionsRoot,
		input.Uncles,
	).Scan(
		&blockData.ID,
		&blockData.ChainId,
		&blockData.Difficulty,
		&blockData.ExtraData,
		&blockData.GasLimit,
		&blockData.GasUsed,
		&blockData.Hash,
		&blockData.LogsBloom,
		&blockData.Miner,
		&blockData.MixHash,
		&blockData.Nonce,
		&blockData.NumberHex,
		&getNumberStr,
		&blockData.ParentHash,
		&blockData.ReceiptsRoot,
		&blockData.Sha3Uncles,
		&blockData.Size,
		&blockData.StateRoot,
		&blockData.Timestamp,
		&blockData.TotalDifficulty,
		&blockData.TransactionsRoot,
		&blockData.Uncles,
	)

	if err != nil {
		return nil, err
	}

	numberInt := new(big.Int)
	numberInt, ok := numberInt.SetString(getNumberStr, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert number string to big.Int")
	}

	blockData.Number = *numberInt

	return blockData, err
}

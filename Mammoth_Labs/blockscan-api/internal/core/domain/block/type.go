package block

import "math/big"

type CreateBlockInput struct {
	ChainId          int     `json:"chain_id" validate:"required"`
	Difficulty       string  `json:"difficulty" validate:"required"`
	ExtraData        string  `json:"extra_data" validate:"required"`
	GasLimit         string  `json:"gas_limit" validate:"required"`
	GasUsed          string  `json:"gas_used" validate:"required"`
	Hash             string  `json:"hash" validate:"required"`
	LogsBloom        string  `json:"logs_bloom" validate:"required"`
	Miner            string  `json:"miner" validate:"required"`
	MixHash          string  `json:"mix_hash" validate:"required"`
	Nonce            string  `json:"nonce" validate:"required"`
	NumberHex        string  `json:"number_hex" validate:"required"`
	Number           big.Int `json:"number" validate:"required"`
	ParentHash       string  `json:"parent_hash" validate:"required"`
	ReceiptsRoot     string  `json:"receipts_root" validate:"required"`
	Sha3Uncles       string  `json:"sha3_uncles" validate:"required"`
	Size             string  `json:"size" validate:"required"`
	StateRoot        string  `json:"state_root" validate:"required"`
	Timestamp        int     `json:"timestamp"`
	TotalDifficulty  string  `json:"total_difficulty" validate:"required"`
	TransactionsRoot string  `json:"transactions_root" validate:"required"`
	Uncles           string  `json:"uncles"`
}

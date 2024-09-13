package block

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Block struct {
	ID               int     `json:"id"`
	ChainId          int     `json:"chain_id"`
	Difficulty       string  `json:"difficulty"`
	ExtraData        string  `json:"extra_data"`
	GasLimit         string  `json:"gas_limit"`
	GasUsed          string  `json:"gas_used"`
	Hash             string  `json:"hash"`
	LogsBloom        string  `json:"logs_bloom"`
	Miner            string  `json:"miner"`
	MixHash          string  `json:"mix_hash"`
	Nonce            string  `json:"nonce"`
	NumberHex        string  `json:"number_hex"`
	Number           big.Int `json:"number"`
	ParentHash       string  `json:"parent_hash"`
	ReceiptsRoot     string  `json:"receipts_root"`
	Sha3Uncles       string  `json:"sha3_uncles"`
	Size             string  `json:"size"`
	StateRoot        string  `json:"state_root"`
	Timestamp        int     `json:"timestamp"`
	TotalDifficulty  string  `json:"total_difficulty"`
	TransactionsRoot string  `json:"transactions_root"`
	Uncles           string  `json:"uncles"`
}

type Reader interface {
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateBlockInput) (*Block, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	Create(queryRower postgresql.Query, input *CreateBlockInput) (*Block, *utils.ServiceError)
}

package transaction

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Transaction struct {
	ID                int     `json:"id"`
	ChainId           int     `json:"chain_id"`
	BlockId           int     `json:"block_id"`
	ContractId        int     `json:"contract_id,omitempty"`
	Timestamp         int     `json:"timestamp"`
	BlockHash         string  `json:"block_hash"`
	BlockNumberHex    string  `json:"block_number_hex"`
	BlockNumber       big.Int `json:"block_number"`
	FromAddress       string  `json:"from_address"`
	Gas               string  `json:"gas"`
	GasPrice          string  `json:"gas_price"`
	Hash              string  `json:"hash"`
	Input             string  `json:"input"`
	Nonce             string  `json:"nonce"`
	R                 string  `json:"r"`
	S                 string  `json:"s"`
	ToAddress         string  `json:"to_address,omitempty"`
	TransactionIndex  string  `json:"transaction_index"`
	Type              string  `json:"type"`
	V                 string  `json:"v"`
	Value             big.Int `json:"value"`
	ContractAddress   string  `json:"contract_address,omitempty"`
	CumulativeGasUsed string  `json:"cumulative_gas_used"`
	GasUsed           string  `json:"gas_used"`
	LogsBloom         string  `json:"logs_bloom"`
	Status            string  `json:"status"`
}

type Reader interface {
	GetCoin(input *GetTransactionInput) ([]*GetTransactionData, error)
	GetToken(input *GetTransactionInput) ([]*GetTransactionData, error)
	GetAll(input *GetAllTransactionInput) ([]*GetTransactionData, error)
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateTransactionInput) (*Transaction, error)
	Update(queryRower postgresql.Query, input *UpdateTransactionInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(input *GetTransactionInput) ([]*GetTransactionData, *utils.ServiceError)
	GetAll(input *GetAllTransactionInput) ([]*GetTransactionData, *utils.ServiceError)
}

type RpcUseCase interface {
	Create(queryRower postgresql.Query, input *CreateTransactionInput) (*Transaction, *utils.ServiceError)
	Update(queryRower postgresql.Query, input *UpdateTransactionInput) *utils.ServiceError
}

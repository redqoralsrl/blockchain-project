package transaction

import (
	"math/big"
)

type CreateTransactionInput struct {
	ChainId           int     `json:"chain_id" validate:"required"`
	BlockId           int     `json:"block_id" validate:"required"`
	Timestamp         int     `json:"timestamp" validate:"required"`
	BlockHash         string  `json:"block_hash" validate:"required"`
	BlockNumberHex    string  `json:"block_number_hex" validate:"required"`
	BlockNumber       big.Int `json:"block_number" validate:"required"`
	FromAddress       string  `json:"from_address" validate:"required"`
	Gas               string  `json:"gas" validate:"required"`
	GasPrice          string  `json:"gas_price" validate:"required"`
	Hash              string  `json:"hash" validate:"required"`
	Input             string  `json:"input" validate:"required"`
	Nonce             string  `json:"nonce" validate:"required"`
	R                 string  `json:"r" validate:"required"`
	S                 string  `json:"s" validate:"required"`
	ToAddress         string  `json:"to_address,omitempty"`
	TransactionIndex  string  `json:"transaction_index" validate:"required"`
	Type              string  `json:"type" validate:"required"`
	V                 string  `json:"v" validate:"required"`
	Value             big.Int `json:"value" validate:"required"`
	ContractAddress   string  `json:"contract_address,omitempty"`
	CumulativeGasUsed string  `json:"cumulative_gas_used" validate:"required"`
	GasUsed           string  `json:"gas_used" validate:"required"`
	LogsBloom         string  `json:"logs_bloom" validate:"required"`
	Status            string  `json:"status"`
}

type UpdateTransactionInput struct {
	TransactionId int `json:"transaction_id" validate:"required"`
	ContractId    int `json:"contract_id" validate:"required"`
}

type GetTransactionInput struct {
	WalletAddress   string `json:"wallet_address" validate:"required"`
	ChainId         int    `json:"chain_id" validate:"required"`
	ChainType       string `json:"chain_type" validate:"omitempty,oneof=coin token"`
	ContractAddress string `json:"contract_address" validate:"omitempty"`
	Take            int    `json:"take" validate:"omitempty,min=0"`
	Skip            int    `json:"skip"`
	OrderBy         string `json:"order_by" validate:"omitempty,oneof=asc desc" description:"정렬 방식"`
	Type            string `json:"type" validate:"omitempty,oneof=send received" description:"트랜잭션 타입(send, received)"`
}

type GetTransactionData struct {
	Hash            string `json:"hash"`
	Timestamp       int    `json:"timestamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	GasUsed         string `json:"gas_used"`
	Gas             string `json:"gas"`
	GasPrice        string `json:"gas_price"`
	Info            string `json:"info"`
	Type            string `json:"type"`
	Status          string `json:"status"`
	ChainId         int    `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	TransactionFee  string `json:"transaction_fee"`
}

type GetAllTransactionInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
	Take          int    `query:"take" validate:"required"`
	Skip          int    `query:"skip"`
}

package nft_volume

import "math/big"

type CreateNftVolumeInput struct {
	ChainId         int     `json:"chain_id" validate:"required"`
	LogId           int     `json:"log_id" validate:"required"`
	FromAddress     string  `json:"from_address" validate:"required"`
	ToAddress       string  `json:"to_address" validate:"required"`
	Value           big.Int `json:"value" validate:"required"`
	Contract        string  `json:"contract" validate:"required"`
	Symbol          string  `json:"symbol" validate:"required"`
	Timestamp       int     `json:"timestamp" validate:"required"`
	TransactionHash string  `json:"transaction_hash" validate:"required"`
	Event           string  `json:"event" validate:"required"`
}

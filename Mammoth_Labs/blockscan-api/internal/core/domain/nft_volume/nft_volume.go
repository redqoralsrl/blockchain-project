package nft_volume

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type NftVolume struct {
	ID              int     `json:"id"`
	ChainId         int     `json:"chain_id"`
	LogId           int     `json:"log_id"`
	FromAddress     string  `json:"from_address"`
	ToAddress       string  `json:"to_address"`
	Value           big.Int `json:"value"`
	Contract        string  `json:"contract"`
	Symbol          string  `json:"symbol"`
	Timestamp       int     `json:"timestamp"`
	TransactionHash string  `json:"transaction_hash"`
	Event           string  `json:"event"`
}

type Reader interface {
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateNftVolumeInput) (*NftVolume, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	Create(queryRower postgresql.Query, input *CreateNftVolumeInput) (*NftVolume, *utils.ServiceError)
}

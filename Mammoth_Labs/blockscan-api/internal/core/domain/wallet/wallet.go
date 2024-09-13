package wallet

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Wallet struct {
	ID            int     `json:"id"`
	Address       string  `json:"address"`
	NickName      string  `json:"nick_name,omitempty"`
	Profile       string  `json:"profile,omitempty"`
	Nonce         big.Int `json:"nonce"`
	SignHash      string  `json:"sign_hash,omitempty"`
	SignTimestamp int     `json:"sign_timestamp,omitempty"`
}

type Reader interface {
	Get(queryRower postgresql.Query, address string) (*Wallet, error)
}

type Writer interface {
	Create(queryRower postgresql.Query, address string) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	Get(queryRower postgresql.Query, address string) (*Wallet, *utils.ServiceError)
	Create(queryRower postgresql.Query, address string) (int, *utils.ServiceError)
}

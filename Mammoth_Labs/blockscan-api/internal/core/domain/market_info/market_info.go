package market_info

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type MarketInfo struct {
	ID                    int     `json:"id"`
	ChainId               int     `json:"chain_id"`
	LogId                 int     `json:"log_id"`
	TransactionHash       string  `json:"transaction_hash"`
	Collection            string  `json:"collection"`
	Seller                string  `json:"seller"`
	Buyer                 string  `json:"buyer"`
	Volume                big.Int `json:"volume"`
	VolumeSymbol          string  `json:"volume_symbol"`
	VolumeContractAddress string  `json:"volume_contract_address"`
}

type Reader interface {
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateMarketInfoInput) (*MarketInfo, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	Create(queryRower postgresql.Query, input *CreateMarketInfoInput) (*MarketInfo, *utils.ServiceError)
}

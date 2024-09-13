package tracking_info

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
	"time"
)

type TrackingInfo struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ChainId     int       `json:"chain_id"`
	BlockHeight big.Int   `json:"block_height"`
	IsOperation bool      `json:"is_operation"`
}

type Reader interface {
	Get(chainId int) (*TrackingInfo, error)
}

type Writer interface {
	Increase(queryRower postgresql.Query, chainId int) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	//Get(chainId int) (*TrackingInfo, *tracking_info_errors.TrackingInfoGetError)
	//Increase(chainId int) *tracking_info_errors.TrackingInfoIncreaseError
}

type RpcUseCase interface {
	Get(chainId int) (*TrackingInfo, *utils.ServiceError)
	Increase(chainId int) *utils.ServiceError
}

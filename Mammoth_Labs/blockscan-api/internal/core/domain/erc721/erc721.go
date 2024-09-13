package erc721

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Erc721 struct {
	ID                  int     `json:"id"`
	ChainId             int     `json:"chain_id"`
	ContractId          int     `json:"contract_id"`
	WalletId            int     `json:"wallet_id"`
	TokenId             big.Int `json:"token_id"`
	Amount              big.Int `json:"amount"`
	Url                 string  `json:"url,omitempty"`
	ImageUrl            string  `json:"image_url,omitempty"`
	Name                string  `json:"name,omitempty"`
	Description         string  `json:"description,omitempty"`
	ExternalUrl         string  `json:"external_url,omitempty"`
	IsUndefinedMetadata bool    `json:"is_undefined_metadata"`
	FloorPrice          big.Int `json:"floor_price,omitempty"`
	AwsImageUrl         string  `json:"aws_image_url,omitempty"`
}

type Reader interface {
	GetEmptyUrlErc721List() ([]*GetEmptyUrlErc721, error)
}

type Writer interface {
	MoveErc721(queryRower postgresql.Query, input *MoveErc721Input) error
	UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) error
	Update(queryRower postgresql.Query, input *attributes.CreateErc721AttributesInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	MoveErc721(queryRower postgresql.Query, input *MoveErc721Input) *utils.ServiceError
}

type CronUseCase interface {
	GetEmptyUrlErc721List() []*GetEmptyUrlErc721
	UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) *utils.ServiceError
	Update(queryRower postgresql.Query, input *attributes.CreateErc721AttributesInput) *utils.ServiceError
}

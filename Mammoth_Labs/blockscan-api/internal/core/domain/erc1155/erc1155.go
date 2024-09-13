package erc1155

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Erc1155 struct {
	ID                  int     `json:"id"`
	ChainId             int     `json:"chain_id"`
	ContractId          int     `json:"contract_id"`
	TokenId             big.Int `json:"token_id"`
	Amount              big.Int `json:"amount"`
	Url                 string  `json:"url,omitempty"`
	ImageUrl            string  `json:"image_url,omitempty"`
	Name                string  `json:"name,omitempty"`
	Description         string  `json:"description,omitempty"`
	ExternalUrl         string  `json:"external_url,omitempty"`
	IsUndefinedMetadata bool    `json:"is_undefined_metadata"`
	AwsImageUrl         string  `json:"aws_image_url,omitempty"`
}

type Erc1155Owner struct {
	ID        int     `json:"id"`
	ChainId   int     `json:"chain_id"`
	Erc1155Id int     `json:"erc_1155_id"`
	WalletId  int     `json:"wallet_id"`
	Amount    big.Int `json:"amount"`
}

type Reader interface {
	GetEmptyUrlErc1155List() ([]*GetEmptyUrlErc1155, error)
}

type Writer interface {
	MoveErc1155(queryRower postgresql.Query, input *MoveErc1155Input) error
	UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) error
	Update(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
	MoveErc1155(queryRower postgresql.Query, input *MoveErc1155Input) *utils.ServiceError
}

type CronUseCase interface {
	GetEmptyUrlErc1155List() []*GetEmptyUrlErc1155
	UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) *utils.ServiceError
	Update(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) *utils.ServiceError
}

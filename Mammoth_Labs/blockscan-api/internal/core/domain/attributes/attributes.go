package attributes

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
)

type Attributes struct {
	ID          int    `json:"id"`
	ChainId     int    `json:"chain_id"`
	ContractId  int    `json:"contract_id"`
	Erc721Id    int    `json:"erc721_id,omitempty"`
	Erc1155Id   int    `json:"erc1155_id,omitempty"`
	TraitType   string `json:"trait_type,omitempty"`
	Value       string `json:"value,omitempty"`
	DisplayType string `json:"display_type,omitempty"`
}

type Reader interface {
}

type Writer interface {
	CreateErc721(queryRower postgresql.Query, input *CreateErc721AttributesInput) error
	CreateErc1155(queryRower postgresql.Query, input *CreateErc1155AttributesInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
}

type RpcUseCase interface {
}

type CronUseCase interface {
	CreateErc721(queryRower postgresql.Query, input *CreateErc721AttributesInput) *utils.ServiceError
	CreateErc1155(queryRower postgresql.Query, input *CreateErc1155AttributesInput) *utils.ServiceError
}

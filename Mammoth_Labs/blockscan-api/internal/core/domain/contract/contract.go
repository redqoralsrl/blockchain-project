package contract

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Contract struct {
	ID                  int     `json:"id"`
	ChainId             int     `json:"chain_id"`
	TransactionCreateId int     `json:"transaction_create_id,omitempty"`
	Timestamp           int     `json:"timestamp"`
	Hash                string  `json:"hash"`
	Name                string  `json:"name,omitempty"`
	Symbol              string  `json:"symbol,omitempty"`
	Decimals            int     `json:"decimals"`
	Description         string  `json:"description,omitempty"`
	IsErc20             bool    `json:"is_erc_20"`
	IsErc721            bool    `json:"is_erc_721"`
	IsErc1155           bool    `json:"is_erc_1155"`
	TotalSupply         big.Int `json:"total_supply"`
	Creator             string  `json:"creator"`
	AwsLogoImage        string  `json:"aws_logo_image,omitempty"`
	AwsBannerImage      string  `json:"aws_banner_image,omitempty"`
	Twitter             string  `json:"twitter,omitempty"`
	Instagram           string  `json:"instagram,omitempty"`
	Homepage            string  `json:"homepage,omitempty"`
	DayVolume           big.Int `json:"day_volume"`
	WeekVolume          big.Int `json:"week_volume"`
	MonthVolume         big.Int `json:"month_volume"`
	AllVolume           big.Int `json:"all_volume"`
}

type Reader interface {
	Get(queryRower postgresql.Query, input *GetContractInput) (*Contract, error)
	GetId(queryRower postgresql.Query, input *GetContractIdInput) (int, error)
	GetType(queryRower postgresql.Query, input *GetContractTypeInput) (int, error)
	GetWalletNFTsByCollection(input *GetContractWalletNFTsByCollectionInput) ([]*GetContractWalletNFTsByCollectionData, error)
	GetCollectionNFTsForWallet(input *GetContractCollectionNFTsForWalletInput) (*GetContractCollectionNFTsForWalletData, error)
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateContractInput) (*Contract, error)
	UpdateType(queryRower postgresql.Query, input *UpdateContractTypeInput) error
	UpdateAllVolume(queryRower postgresql.Query, input *UpdateContractAllVolumeInput) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetWalletNFTsByCollection(input *GetContractWalletNFTsByCollectionInput) ([]*GetContractWalletNFTsByCollectionData, *utils.ServiceError)
	GetCollectionNFTsForWallet(input *GetContractCollectionNFTsForWalletInput) (*GetContractCollectionNFTsForWalletData, *utils.ServiceError)
}

type RpcUseCase interface {
	Get(queryRower postgresql.Query, input *GetContractInput) (*Contract, *utils.ServiceError)
	GetId(queryRower postgresql.Query, input *GetContractIdInput) (int, *utils.ServiceError)
	GetType(queryRower postgresql.Query, input *GetContractTypeInput) (int, *utils.ServiceError)
	Create(queryRower postgresql.Query, input *CreateContractInput) (*Contract, *utils.ServiceError)
	UpdateType(queryRower postgresql.Query, input *UpdateContractTypeInput) *utils.ServiceError
	UpdateAllVolume(queryRower postgresql.Query, input *UpdateContractAllVolumeInput) *utils.ServiceError
}

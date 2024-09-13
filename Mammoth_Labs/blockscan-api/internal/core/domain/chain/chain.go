package chain

import "blockscan-go/internal/core/common/utils"

type Chain struct {
	ID              int    `json:"id"`
	ChainId         int    `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        int    `json:"decimals"`
	ImageUrl        string `json:"image_url"`
	Site            string `json:"site"`
	ScanSite        string `json:"scan_site"`
}

type Reader interface {
	Get(chainId int) (*Chain, error)
	GetToken(input *GetTokenChainInput) (*Chain, error)
	GetTokens(input *GetTokensChainInput) ([]*Chain, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(chainId int) (*Chain, *utils.ServiceError)
	GetToken(input *GetTokenChainInput) (*Chain, *utils.ServiceError)
	GetTokens(input *GetTokensChainInput) ([]*Chain, *utils.ServiceError)
}

type RpcUseCase interface {
}

package contract

import "math/big"

type GetContractInput struct {
	ChainId int    `json:"chain_id" validate:"required"`
	Hash    string `json:"hash" validate:"required"`
}

type CreateContractInput struct {
	ChainId             int     `json:"chain_id" validate:"required"`
	TransactionCreateId int     `json:"transaction_create_id" validate:"required"`
	Timestamp           int     `json:"timestamp" validate:"required"`
	Hash                string  `json:"hash" validate:"required"`
	Name                string  `json:"name"`
	Symbol              string  `json:"symbol"`
	Decimals            int     `json:"decimals"`
	IsErc20             bool    `json:"is_erc_20"`
	IsErc721            bool    `json:"is_erc_721"`
	IsErc1155           bool    `json:"is_erc_1155"`
	TotalSupply         big.Int `json:"total_supply,omitempty"`
	Creator             string  `json:"creator" validate:"required"`
}

type GetContractIdInput struct {
	ChainId         int    `json:"chain_id" validate:"required"`
	ContractAddress string `json:"contract_address" validate:"required"`
}

type GetContractTypeInput struct {
	ChainId int    `json:"chain_id" validate:"required"`
	Hash    string `json:"hash" validate:"required"`
}

type GetContractTypeInfo struct {
	IsErc20   bool `json:"is_erc_20"`
	IsErc721  bool `json:"is_erc_721"`
	IsErc1155 bool `json:"is_erc_1155"`
}

type UpdateContractTypeInput struct {
	ChainId int    `json:"chain_id" validate:"required"`
	Hash    string `json:"hash" validate:"required"`
	Type    int    `json:"type" validate:"required"`
}

type UpdateContractAllVolumeInput struct {
	ChainId int     `json:"chain_id" validate:"required"`
	Hash    string  `json:"hash" validate:"required"`
	Volume  big.Int `json:"volume" validate:"required"`
}

type GetContractWalletNFTsByCollectionInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
	Take          int    `query:"take" validate:"required"`
	Skip          int    `query:"skip"`
}

type GetContractWalletNFTsByCollectionData struct {
	ID                int       `json:"id"`
	ChainId           int       `json:"chain_id"`
	Hash              string    `json:"hash"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	AwsLogoImage      string    `json:"aws_logo_image"`
	AwsBannerImage    string    `json:"aws_banner_image"`
	Description       string    `json:"description"`
	ErcType           string    `json:"erc_type"`
	WalletAddress     string    `json:"wallet_address"`
	WalletNickname    string    `json:"wallet_nickname"`
	WalletProfile     string    `json:"wallet_profile"`
	ErcList           []ErcList `json:"erc_list,omitempty"`
	UniqueTokensCount string    `json:"unique_tokens_count"`
	TokensAmount      string    `json:"tokens_amount,omitempty"`
}

// ErcList 최대 10개 까지만 제공
type ErcList struct {
	TokenId     string `json:"token_id"`
	Amount      string `json:"amount"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AwsImageUrl string `json:"aws_image_url"`
}

type GetContractCollectionNFTsForWalletInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	Hash          string `param:"hash" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
	Take          int    `query:"take" validate:"required"`
	Skip          int    `query:"skip"`
}

type GetContractCollectionNFTsForWalletData struct {
	ID             int       `json:"id"`
	ChainId        int       `json:"chain_id"`
	Hash           string    `json:"hash"`
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	AwsLogoImage   string    `json:"aws_logo_image"`
	AwsBannerImage string    `json:"aws_banner_image"`
	Description    string    `json:"description"`
	ErcType        string    `json:"erc_type"`
	WalletAddress  string    `json:"wallet_address"`
	WalletNickname string    `json:"wallet_nickname"`
	WalletProfile  string    `json:"wallet_profile"`
	ErcList        []ErcList `json:"erc_list,omitempty"`
}

package nft

type GetNftInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
	Take          int    `query:"take" validate:"required"`
	Skip          int    `query:"skip"`
}

type GetNftData struct {
	ID                  int          `json:"id"`
	ChainId             int          `json:"chain_id"`
	ContractId          int          `json:"contract_id"`
	ContractHash        string       `json:"contract_hash"`
	ContractName        string       `json:"contract_name"`
	ContractSymbol      string       `json:"contract_symbol"`
	ContractDescription string       `json:"contract_description"`
	AwsLogoImage        string       `json:"aws_logo_image"`
	AwsBannerImage      string       `json:"aws_banner_image"`
	TokenId             string       `json:"token_id"`
	Amount              string       `json:"amount"`
	Type                string       `json:"type"`
	Url                 string       `json:"url"`
	ImageUrl            string       `json:"image_url"`
	NftName             string       `json:"nft_name"`
	NftDescription      string       `json:"nft_description"`
	AttributesArray     []Attributes `json:"attributes_array,omitempty"`
}

type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type GetNftDetailInput struct {
	Hash    string `param:"hash" validate:"required"`
	TokenId string `param:"token_id" validate:"required"`
	ChainId int    `query:"chain_id" validate:"required"`
}

type GetNFtDetailData struct {
	ID                  int                 `json:"id"`
	ChainId             int                 `json:"chain_id"`
	ContractId          int                 `json:"contract_id"`
	ContractHash        string              `json:"contract_hash"`
	ContractName        string              `json:"contract_name"`
	ContractSymbol      string              `json:"contract_symbol"`
	ContractDescription string              `json:"contract_description"`
	AwsLogoImage        string              `json:"aws_logo_image"`
	AwsBannerImage      string              `json:"aws_banner_image"`
	TokenId             string              `json:"token_id"`
	Amount              string              `json:"amount"`
	Type                string              `json:"type"`
	Url                 string              `json:"url"`
	ImageUrl            string              `json:"image_url"`
	NftName             string              `json:"nft_name"`
	NftDescription      string              `json:"nft_description"`
	AttributesArray     []AttributesPercent `json:"attributes_array,omitempty"`
}
type AttributesPercent struct {
	TraitType     string  `json:"trait_type"`
	Value         string  `json:"value"`
	TotalCount    int     `json:"total_count"`
	SpecificCount int     `json:"specific_count"`
	Percentage    float64 `json:"percentage"`
}

type GetNftDetailOfWalletInput struct {
	Hash          string `param:"hash" validate:"required"`
	TokenId       string `param:"token_id" validate:"required"`
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
}

package chain

type GetChainInput struct {
	ChainId int `param:"chain_id" validate:"required"`
}

type GetTokenChainInput struct {
	ChainId         int    `param:"chain_id" validate:"required"`
	ContractAddress string `param:"contract_address" validate:"required"`
}

type GetTokensChainInput struct {
	ChainId int `param:"chain_id" validate:"required"`
}

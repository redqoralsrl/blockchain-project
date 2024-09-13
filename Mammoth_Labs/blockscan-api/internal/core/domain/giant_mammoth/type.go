package giant_mammoth

type GetStakingByAccountInput struct {
	ChainId       int    `param:"chain_id" validate:"required"`
	WalletAddress string `param:"wallet_address" validate:"required"`
}

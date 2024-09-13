package market_info

import "math/big"

type CreateMarketInfoInput struct {
	ChainId               int     `json:"chain_id" validate:"required"`
	LogId                 int     `json:"log_id" validate:"required"`
	TransactionHash       string  `json:"transaction_hash" validate:"required"`
	Collection            string  `json:"collection" validate:"required"`
	Seller                string  `json:"seller" validate:"required"`
	Buyer                 string  `json:"buyer" validate:"required"`
	Volume                big.Int `json:"volume" validate:"required"`
	VolumeSymbol          string  `json:"volume_symbol"`
	VolumeContractAddress string  `json:"volume_contract_address"`
}

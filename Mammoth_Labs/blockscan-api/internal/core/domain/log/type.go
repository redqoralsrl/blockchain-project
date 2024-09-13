package log

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type CreateLogInput struct {
	ChainId                int           `json:"chain_id" validate:"required"`
	TransactionId          int           `json:"transaction_id" validate:"required"`
	Address                string        `json:"address,omitempty"`
	BlockHash              string        `json:"block_hash" validate:"required"`
	BlockNumberHex         string        `json:"block_number_hex" validate:"required"`
	BlockNumber            big.Int       `json:"block_number" validate:"required"`
	Data                   string        `json:"data"`
	LogIndex               string        `json:"log_index" validate:"required"`
	Removed                bool          `json:"removed"`
	TransactionHash        string        `json:"transaction_hash" validate:"required"`
	TransactionIndex       string        `json:"transaction_index" validate:"required"`
	Timestamp              int           `json:"timestamp" validate:"required"`
	Function               string        `json:"function,omitempty"`
	Type                   int           `json:"type,omitempty"`
	Dapp                   string        `json:"dapp,omitempty"`
	FromAddress            string        `json:"from_address,omitempty"`
	ToAddress              string        `json:"to_address,omitempty"`
	Value                  big.Int       `json:"value,omitempty"`
	TokenId                big.Int       `json:"token_id,omitempty"`
	Url                    string        `json:"url,omitempty"`
	Name                   string        `json:"name,omitempty"`
	Symbol                 string        `json:"symbol,omitempty"`
	Decimals               uint8         `json:"decimals,omitempty"`
	Erc1155Value           []*big.Int    `json:"erc_1155_value,omitempty"`
	Erc1155TokenId         []*big.Int    `json:"erc_1155_token_id,omitempty"`
	Erc1155Url             []string      `json:"erc_1155_url,omitempty"`
	TradeNftVolume         big.Int       `json:"trade_nft_volume,omitempty"`
	TradeNftVolumeSymbol   string        `json:"trade_nft_volume_symbol,omitempty"`
	TradeNftVolumeContract string        `json:"trade_nft_volume_contract,omitempty"`
	Topics                 []common.Hash `json:"topics,omitempty"`
}

type GetLogNFTsByWalletInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `query:"chain_id" validate:"required"`
	Take          int    `query:"take" validate:"required"`
	Skip          int    `query:"skip"`
}

type GetLogNFTsByWalletData struct {
	ChainId                int             `json:"chain_id"`
	Address                string          `json:"address,omitempty"`
	BlockHash              string          `json:"block_hash"`
	BlockNumber            string          `json:"block_number"`
	TransactionHash        string          `json:"transaction_hash"`
	Timestamp              int             `json:"timestamp"`
	Function               string          `json:"function"`
	Type                   int             `json:"type"`
	DApp                   string          `json:"dapp"`
	FromAddress            string          `json:"from_address"`
	ToAddress              string          `json:"to_address"`
	Value                  string          `json:"value"`
	TokenId                string          `json:"token_id"`
	Url                    string          `json:"url"`
	Name                   string          `json:"name"`
	Symbol                 string          `json:"symbol"`
	TradeNftVolume         string          `json:"trade_nft_volume"`
	TradeNftVolumeSymbol   string          `json:"trade_nft_volume_symbol"`
	TradeNftVolumeContract string          `json:"trade_nft_volume_contract"`
	Topics                 []string        `json:"topics"`
	Erc1155Value           []string        `json:"erc1155_value"`
	Erc1155TokenId         []string        `json:"erc1155_token_id"`
	Erc1155Url             []string        `json:"erc1155_url"`
	Gas                    string          `json:"gas"`
	GasPrice               string          `json:"gas_price"`
	GasUsed                string          `json:"gas_used"`
	NftVolumeArray         []NftVolumeData `json:"nft_volume_array"`
	MarketInfoArray        []MarketInfo    `json:"market_info_array"`
}
type MarketInfo struct {
	ChainId               int    `json:"chain_id"`
	TransactionHash       string `json:"transaction_hash"`
	Collection            string `json:"collection"`
	Seller                string `json:"seller"`
	Buyer                 string `json:"buyer"`
	Volume                string `json:"volume"`
	VolumeSymbol          string `json:"volume_symbol"`
	VolumeContractAddress string `json:"volume_contract_address"`
}
type NftVolumeData struct {
	ChainId         int    `json:"chain_id"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Value           string `json:"value"`
	Contract        string `json:"contract"`
	Symbol          string `json:"symbol"`
	Timestamp       int    `json:"timestamp"`
	TransactionHash string `json:"transaction_hash"`
	Event           string `json:"event"`
}
type MarketInfoTemp struct {
	ChainId               int    `json:"chain_id"`
	TransactionHash       string `json:"transaction_hash"`
	Collection            string `json:"collection"`
	Seller                string `json:"seller"`
	Buyer                 string `json:"buyer"`
	Volume                string `json:"volume"`
	VolumeSymbol          string `json:"volume_symbol"`
	VolumeContractAddress string `json:"volume_contract_address"`
}

type NftVolumeDataTemp struct {
	ChainId         int    `json:"chain_id"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Value           string `json:"value"`
	Contract        string `json:"contract"`
	Symbol          string `json:"symbol"`
	Timestamp       int    `json:"timestamp"`
	TransactionHash string `json:"transaction_hash"`
	Event           string `json:"event"`
}

type GetSearchNFTsByWalletInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `json:"chain_id" validate:"required"`
	Date          string `json:"date" validate:"required,oneof=all 7D 30D 1Y"`
	Type          string `json:"type" validate:"required,oneof=all mint transfer sale burn"`
	Take          int    `json:"take" validate:"required"`
	Skip          int    `json:"skip"`
}

type GetSearchByWalletInput struct {
	WalletAddress string `param:"wallet_address" validate:"required"`
	ChainId       int    `json:"chain_id" validate:"required"`
	Date          string `json:"date" validate:"required,oneof=all 7D 30D 1Y"`
	Type          string `json:"type" validate:"required,oneof=all mint transfer sale burn"`
	Take          int    `json:"take" validate:"required"`
	Skip          int    `json:"skip"`
}

type GetSearchByWalletData struct {
	ID              int    `json:"id"`
	ChainId         int    `json:"chain_id"`
	TransactionId   int    `json:"transaction_id"`
	TransactionHash string `json:"transaction_hash"`
	ContractId      int    `json:"contract_id"`
	ContractHash    string `json:"contract_hash"`
	Timestamp       int    `json:"timestamp"`
	Function        string `json:"function"`
	Type            int    `json:"type"`
	Dapp            string `json:"dapp"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Value           string `json:"value"`
	TokenId         string `json:"token_id"`
	// Nullable
	Url string `json:"url,omitempty"`
	// Nullable
	ContractName string `json:"contract_name,omitempty"`
	// Nullable
	ContractSymbol string `json:"contract_symbol,omitempty"`
	// Nullable
	NftImageUrl string `json:"nft_image_url,omitempty"`
	// Nullable
	NftName string `json:"nft_name,omitempty"`
	// Default : 0x0000...
	Seller string `json:"seller"`
	// Default : 0x0000...
	Buyer string `json:"buyer"`
	// Nullable
	VolumeSymbol string `json:"volume_symbol,omitempty"`
	// Nullable
	VolumeContract string `json:"volume_contract,omitempty"`
	// Nullable
	TradeNftVolume string `json:"trade_nft_volume,omitempty"`
	Gas            string `json:"gas"`
	GasPrice       string `json:"gas_price"`
	GasUsed        string `json:"gas_used"`
}

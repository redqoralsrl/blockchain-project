package log

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"math/big"
)

type Log struct {
	ID                     int        `json:"id"`
	ChainId                int        `json:"chain_id"`
	TransactionId          int        `json:"transaction_id"`
	Address                string     `json:"address,omitempty"`
	BlockHash              string     `json:"block_hash"`
	BlockNumberHex         string     `json:"block_number_hex"`
	BlockNumber            big.Int    `json:"block_number"`
	Data                   string     `json:"data,omitempty"`
	LogIndex               string     `json:"log_index"`
	Removed                bool       `json:"removed"`
	TransactionHash        string     `json:"transaction_hash"`
	TransactionIndex       string     `json:"transaction_index"`
	Timestamp              int        `json:"timestamp"`
	Function               string     `json:"function,omitempty"`
	Type                   int        `json:"type,omitempty"`
	Dapp                   string     `json:"dapp,omitempty"`
	FromAddress            string     `json:"from_address,omitempty"`
	ToAddress              string     `json:"to_address,omitempty"`
	Value                  big.Int    `json:"value,omitempty"`
	TokenId                big.Int    `json:"token_id,omitempty"`
	Url                    string     `json:"url,omitempty"`
	Name                   string     `json:"name,omitempty"`
	Symbol                 string     `json:"symbol,omitempty"`
	Decimals               uint8      `json:"decimals,omitempty"`
	Erc1155Value           []*big.Int `json:"erc1155_value,omitempty"`
	Erc1155TokenId         []*big.Int `json:"erc1155_token_id,omitempty"`
	Erc1155Url             []string   `json:"erc1155_url,omitempty"`
	TradeNftVolume         big.Int    `json:"trade_nft_volume,omitempty"`
	TradeNftVolumeSymbol   string     `json:"trade_nft_volume_symbol,omitempty"`
	TradeNftVolumeContract string     `json:"trade_nft_volume_contract,omitempty"`
	Topics                 []string   `json:"topics,omitempty"`
}

type Reader interface {
	GetNFTsByWallet(input *GetLogNFTsByWalletInput) ([]*GetLogNFTsByWalletData, error)
	GetSearchNFTsByWallet(input *GetSearchNFTsByWalletInput) ([]*GetLogNFTsByWalletData, error)
}

type Writer interface {
	Create(queryRower postgresql.Query, input *CreateLogInput) (*Log, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetNFTsByWallet(input *GetLogNFTsByWalletInput) ([]*GetLogNFTsByWalletData, *utils.ServiceError)
	GetSearchNFTsByWallet(input *GetSearchNFTsByWalletInput) ([]*GetLogNFTsByWalletData, *utils.ServiceError)
}

type RpcUseCase interface {
	Create(queryRower postgresql.Query, input *CreateLogInput) (*Log, *utils.ServiceError)
}

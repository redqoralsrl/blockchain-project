package erc721

import "math/big"

type MoveErc721Input struct {
	ChainId    int     `json:"chain_id" validate:"required"`
	ContractId int     `json:"contract_id" validate:"required"`
	From       string  `json:"from" validate:"required"`
	To         string  `json:"to" validate:"required"`
	TokenId    big.Int `json:"token_id" validate:"required"`
	Amount     big.Int `json:"amount" validate:"required"`
	Url        string  `json:"url"`
}

type GetEmptyUrlErc721 struct {
	Erc721ID   int    `json:"erc_721_id"`
	ContractId int    `json:"contract_id"`
	Url        string `json:"url"`
	ChainId    int    `json:"chain_id"`
}

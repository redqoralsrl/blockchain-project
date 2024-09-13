package erc1155

import "math/big"

type MoveErc1155Input struct {
	ChainId        int        `json:"chain_id" validate:"required"`
	ContractId     int        `json:"contract_id" validate:"required"`
	From           string     `json:"from" validate:"required"`
	To             string     `json:"to" validate:"required"`
	Erc1155TokenId []*big.Int `json:"erc1155_token_id" validate:"required"`
	Erc1155Value   []*big.Int `json:"erc1155_value" validate:"required"`
	Erc1155Url     []string   `json:"erc1155_url" validate:"required"`
}

type GetEmptyUrlErc1155 struct {
	Erc1155ID  int    `json:"erc_1155_id"`
	ContractId int    `json:"contract_id"`
	Url        string `json:"url"`
	ChainId    int    `json:"chain_id"`
}

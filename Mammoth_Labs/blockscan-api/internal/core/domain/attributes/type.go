package attributes

type CreateErc721AttributesInput struct {
	Erc721Id      int                     `json:"erc721_id"`
	ImageUrl      string                  `json:"image_url,omitempty"`
	Name          string                  `json:"name,omitempty"`
	Description   string                  `json:"description,omitempty"`
	ExternalUrl   string                  `json:"external_url,omitempty"`
	AttributeList []Erc721AttributesInput `json:"attribute_list,omitempty"`
}

type Erc721AttributesInput struct {
	ChainId     int    `json:"chain_id"`
	ContractId  int    `json:"contract_id"`
	Erc721Id    int    `json:"erc721_id"`
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
	DisplayType string `json:"display_type,omitempty"`
}

type CreateErc1155AttributesInput struct {
	Erc1155Id     int                      `json:"erc1155_id"`
	ImageUrl      string                   `json:"image_url,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Description   string                   `json:"description,omitempty"`
	ExternalUrl   string                   `json:"external_url,omitempty"`
	AttributeList []Erc1155AttributesInput `json:"attribute_list,omitempty"`
}

type Erc1155AttributesInput struct {
	ChainId     int    `json:"chain_id"`
	ContractId  int    `json:"contract_id"`
	Erc1155Id   int    `json:"erc1155_id"`
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
	DisplayType string `json:"display_type,omitempty"`
}

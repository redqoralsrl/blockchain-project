package crypto_currency

type CreateCryptoCurrencyInput struct {
	Timestamp    int     `json:"timestamp" validate:"ne=0"`
	MorningDate  int     `json:"morning_date" validate:"ne=0"`
	MidnightDate int     `json:"midnight_date" validate:"ne=0"`
	LastUpdated  int     `json:"last_updated" validate:"ne=0"`
	EthPrice     float64 `json:"eth_price" validate:"ne=0"`
	MmtPrice     float64 `json:"mmt_price" validate:"ne=0"`
	GmmtPrice    float64 `json:"gmmt_price" validate:"ne=0"`
	MaticPrice   float64 `json:"matic_price" validate:"ne=0"`
	BnbPrice     float64 `json:"bnb_price" validate:"ne=0"`
}

type GetCryptoCurrencyData struct {
	Timestamp  int    `json:"timestamp"`
	EthPrice   string `json:"eth_price"`
	MmtPrice   string `json:"mmt_price"`
	GmmtPrice  string `json:"gmmt_price"`
	MaticPrice string `json:"matic_price"`
	BnbPrice   string `json:"bnb_price"`
}

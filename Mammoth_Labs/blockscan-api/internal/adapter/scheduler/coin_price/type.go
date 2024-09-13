package coin_price

import "time"

type Response struct {
	Data   map[string]CurrencyData `json:"data"`
	Status Status                  `json:"status"`
}

type Status struct {
	CreditCount  int       `json:"credit_count"`
	Elapsed      int       `json:"elapsed"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage *string   `json:"error_message"`
	Notice       *string   `json:"notice"`
	Timestamp    time.Time `json:"timestamp"`
}

type CurrencyData struct {
	CirculatingSupply             float64          `json:"circulating_supply"`
	CMCRank                       *int             `json:"cmc_rank"`
	DateAdded                     time.Time        `json:"date_added"`
	ID                            int              `json:"id"`
	InfiniteSupply                bool             `json:"infinite_supply"`
	IsActive                      int              `json:"is_active"`
	IsFiat                        int              `json:"is_fiat"`
	LastUpdated                   time.Time        `json:"last_updated"`
	MaxSupply                     *float64         `json:"max_supply"`
	Name                          string           `json:"name"`
	NumMarketPairs                int              `json:"num_market_pairs"`
	Platform                      interface{}      `json:"platform"`
	Quote                         map[string]Quote `json:"quote"`
	SelfReportedCirculatingSupply *float64         `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64         `json:"self_reported_market_cap"`
	Slug                          string           `json:"slug"`
	Symbol                        string           `json:"symbol"`
	Tags                          interface{}      `json:"tags"`
	TotalSupply                   float64          `json:"total_supply"`
	TVLRatio                      *float64         `json:"tvl_ratio"`
}

type Quote struct {
	FullyDilutedMarketCap float64     `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time   `json:"last_updated"`
	MarketCap             *float64    `json:"market_cap"`
	MarketCapDominance    float64     `json:"market_cap_dominance"`
	PercentChange1h       float64     `json:"percent_change_1h"`
	PercentChange24h      float64     `json:"percent_change_24h"`
	PercentChange30d      float64     `json:"percent_change_30d"`
	PercentChange60d      float64     `json:"percent_change_60d"`
	PercentChange7d       float64     `json:"percent_change_7d"`
	PercentChange90d      float64     `json:"percent_change_90d"`
	Price                 *float64    `json:"price"`
	Tvl                   interface{} `json:"tvl"`
	Volume24h             float64     `json:"volume_24h"`
	VolumeChange24h       float64     `json:"volume_change_24h"`
}

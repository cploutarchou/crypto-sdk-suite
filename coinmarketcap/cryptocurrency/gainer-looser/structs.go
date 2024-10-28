package gainer_looser

import "time"

type CryptocurrencyData struct {
	ID                CryptocurrencyID `json:"id"`
	Name              string           `json:"name"`
	Symbol            string           `json:"symbol"`
	Slug              string           `json:"slug"`
	CmcRank           Rank             `json:"cmc_rank,omitempty"`
	NumMarketPairs    int              `json:"num_market_pairs"`
	CirculatingSupply int              `json:"circulating_supply"`
	TotalSupply       int              `json:"total_supply"`
	MaxSupply         int              `json:"max_supply"`
	LastUpdated       time.Time        `json:"last_updated"`
	DateAdded         time.Time        `json:"date_added"`
	Tags              []string         `json:"tags"`
	Platform          any              `json:"platform"`
	Quote             Quote            `json:"quote"`
}

type Response struct {
	Data struct {
		Data   []CryptocurrencyData `json:"data"`
		Status Status               `json:"status"`
	} `json:"data"`
	Status Status `json:"status"`
}

type Quote map[string]CurrencyQuote

type CurrencyQuote struct {
	Price            float64   `json:"price"`
	Volume24H        float64   `json:"volume_24h"`
	PercentChange1H  float64   `json:"percent_change_1h"`
	PercentChange24H float64   `json:"percent_change_24h"`
	PercentChange7D  float64   `json:"percent_change_7d"`
	MarketCap        float64   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

// Params represents the query parameters for fetching data.
type Params struct {
	Limit      *int        `json:"limit,omitempty"`
	TimePeriod *TimePeriod `json:"time_period,omitempty"`
	Convert    *string     `json:"convert,omitempty"`
	ConvertID  *int        `json:"convert_id,omitempty"`
	Sort       *Sort       `json:"sort,omitempty"`
	SortDir    *SortDir    `json:"sort_dir,omitempty"`
}

type CryptocurrencyID int

type Rank int

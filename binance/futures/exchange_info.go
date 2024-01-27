package binance

type ExchangeInfo struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ServerTime      int64         `json:"serverTime"`
	Assets          []Asset       `json:"assets"`
	Symbols         []Symbol      `json:"symbols"`
	Timezone        string        `json:"timezone"`
}

type RateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type Asset struct {
	Asset             string `json:"asset"`
	MarginAvailable   bool   `json:"marginAvailable"`
	AutoAssetExchange int    `json:"autoAssetExchange"`
}

type Symbol struct {
	Symbol string `json:"symbol"`
	Pair   string `json:"pair"`
	// Add more fields as needed
}

// OrderBookResponse represents the response for the order book.
type OrderBookResponse struct {
	LastUpdateID    int64      `json:"lastUpdateId"`
	EventTime       int64      `json:"E"` // Message output time
	TransactionTime int64      `json:"T"` // Transaction time
	Bids            [][]string `json:"bids"`
	Asks            [][]string `json:"asks"`
}

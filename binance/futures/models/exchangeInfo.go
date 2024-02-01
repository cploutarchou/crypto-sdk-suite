package models

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
}

// OrderBookResponse represents the response for the order book.
type OrderBookResponse struct {
	LastUpdateID    int64      `json:"lastUpdateId"`
	EventTime       int64      `json:"E"` // Message output time
	TransactionTime int64      `json:"T"` // Transaction time
	Bids            [][]string `json:"bids"`
	Asks            [][]string `json:"asks"`
}

// Trade represents a single trade from the recent trades list.
type Trade struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

// AggregateTrade represents a single aggregate trade.
type AggregateTrade struct {
	AggregateTradeID int64  `json:"a"` // Aggregate tradeId
	Price            string `json:"p"` // Price
	Quantity         string `json:"q"` // Quantity
	FirstTradeID     int64  `json:"f"` // First tradeId
	LastTradeID      int64  `json:"l"` // Last tradeId
	Timestamp        int64  `json:"T"` // Timestamp
	WasBuyerMaker    bool   `json:"m"` // Was the buyer the maker?
}
type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"` // UNIX timestamp in milliseconds
}

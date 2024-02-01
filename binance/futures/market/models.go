package market

// ExchangeInfo represents information about the exchange, including rate limits, server time, available assets, symbols, and timezone.
type ExchangeInfo struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"` // Filters applied to the exchange.
	RateLimits      []RateLimit   `json:"rateLimits"`      // Rate limits for API requests.
	ServerTime      int64         `json:"serverTime"`      // Current server time in milliseconds.
	Assets          []Asset       `json:"assets"`          // List of available assets on the exchange.
	Symbols         []Symbol      `json:"symbols"`         // List of available trading pairs (symbols).
	Timezone        string        `json:"timezone"`        // Timezone of the exchange.
}

// RateLimit represents a rate limit for API requests.
type RateLimit struct {
	Interval      string `json:"interval"`      // Time interval (e.g., "SECOND", "MINUTE").
	IntervalNum   int    `json:"intervalNum"`   // Number of intervals allowed.
	Limit         int    `json:"limit"`         // Maximum number of requests allowed in the given interval.
	RateLimitType string `json:"rateLimitType"` // Type of rate limit (e.g., "REQUESTS", "ORDERS").
}

// Asset represents information about an asset available on the exchange.
type Asset struct {
	Asset             string `json:"asset"`             // Asset symbol (e.g., "BTC").
	MarginAvailable   bool   `json:"marginAvailable"`   // Whether margin trading is available for this asset.
	AutoAssetExchange int    `json:"autoAssetExchange"` // Auto asset exchange status.
}

// Symbol represents a trading pair (symbol) available on the exchange.
type Symbol struct {
	Symbol string `json:"symbol"` // Symbol (e.g., "BTCUSDT").
	Pair   string `json:"pair"`   // Pair (e.g., "BTC/USDT").
}

// OrderBookResponse represents the response for the order book.
type OrderBookResponse struct {
	LastUpdateID    int64      `json:"lastUpdateId"` // Last update ID.
	EventTime       int64      `json:"E"`            // Message output time.
	TransactionTime int64      `json:"T"`            // Transaction time.
	Bids            [][]string `json:"bids"`         // List of bid prices and quantities.
	Asks            [][]string `json:"asks"`         // List of ask prices and quantities.
}

// Trade represents a single trade from the recent trades list.
type Trade struct {
	ID           int64  `json:"id"`           // Trade ID.
	Price        string `json:"price"`        // Trade price.
	Qty          string `json:"qty"`          // Trade quantity.
	QuoteQty     string `json:"quoteQty"`     // Quote quantity.
	Time         int64  `json:"time"`         // Trade execution time.
	IsBuyerMaker bool   `json:"isBuyerMaker"` // Whether the buyer is the maker.
}

// AggregateTrade represents a single aggregate trade.
type AggregateTrade struct {
	AggregateTradeID int64  `json:"a"` // Aggregate trade ID.
	Price            string `json:"p"` // Price of the aggregate trade.
	Quantity         string `json:"q"` // Quantity of the aggregate trade.
	FirstTradeID     int64  `json:"f"` // First trade ID in the aggregate trade.
	LastTradeID      int64  `json:"l"` // Last trade ID in the aggregate trade.
	Timestamp        int64  `json:"T"` // Timestamp of the aggregate trade.
	WasBuyerMaker    bool   `json:"m"` // Whether the buyer was the maker.
}

// ServerTimeResponse represents the server time response.
type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"` // Current server time in milliseconds.
}

package asset

// GetCoinExchangeRecordsRequest represents the query parameters for fetching coin exchange records.
type GetCoinExchangeRecordsRequest struct {
	FromCoin *string `json:"fromCoin,omitempty"` // Optional: The currency to convert from
	ToCoin   *string `json:"toCoin,omitempty"`   // Optional: The currency to convert to
	Limit    *int    `json:"limit,omitempty"`    // Optional: Limit for data size per page
	Cursor   *string `json:"cursor,omitempty"`   // Optional: Cursor for pagination
}

// CoinExchangeRecord represents a single record of a coin exchange.
type CoinExchangeRecord struct {
	FromCoin     string `json:"fromCoin"`
	FromAmount   string `json:"fromAmount"`
	ToCoin       string `json:"toCoin"`
	ToAmount     string `json:"toAmount"`
	ExchangeRate string `json:"exchangeRate"`
	CreatedTime  string `json:"createdTime"`
	ExchangeTxId string `json:"exchangeTxId"`
}

// GetCoinExchangeRecordsResponse represents the response from fetching coin exchange records.
type GetCoinExchangeRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderBody      []CoinExchangeRecord `json:"orderBody"`
		NextPageCursor string               `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

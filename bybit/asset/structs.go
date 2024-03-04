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

// GetDeliveryRecordRequest represents the query parameters for fetching delivery records.
type GetDeliveryRecordRequest struct {
	Category  string  `json:"category"`            // Required: Product type. option, linear
	Symbol    *string `json:"symbol,omitempty"`    // Optional: Symbol name
	StartTime *int64  `json:"startTime,omitempty"` // Optional: Start timestamp (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: End time. timestamp (ms)
	ExpDate   *string `json:"expDate,omitempty"`   // Optional: Expiry date. 25MAR22
	Limit     *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Cursor for pagination
}

// DeliveryRecordEntry represents a single entry in the delivery record list.
type DeliveryRecordEntry struct {
	DeliveryTime  int64  `json:"deliveryTime"`  // Delivery time (ms)
	Symbol        string `json:"symbol"`        // Symbol name
	Side          string `json:"side"`          // Buy, Sell
	Position      string `json:"position"`      // Executed size
	DeliveryPrice string `json:"deliveryPrice"` // Delivery price
	Strike        string `json:"strike"`        // Exercise price
	Fee           string `json:"fee"`           // Trading fee
	DeliveryRpl   string `json:"deliveryRpl"`   // Realized PnL of the delivery
}

// GetDeliveryRecordResponse represents the response from fetching delivery records.
type GetDeliveryRecordResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		NextPageCursor string                `json:"nextPageCursor"`
		Category       string                `json:"category"`
		List           []DeliveryRecordEntry `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

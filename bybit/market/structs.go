package market

type ApiResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result,omitempty"`
	RetExtInfo interface{} `json:"retExtInfo,omitempty"`
	Time       int64       `json:"time"`
}

// KlineRequest represents a request for querying historical klines
type KlineRequest struct {
	Category string `json:"category,omitempty"` // Optional: 'spot', 'linear', 'inverse'. Defaults to 'linear' if not specified.
	Symbol   string `json:"symbol"`             // Required: Symbol name.
	Interval string `json:"interval"`           // Required: Kline interval. Accepts '1', '3', '5', '15', '30', '60', '120', '240', '360', '720', 'D', 'M', 'W'.
	Start    *int64 `json:"start,omitempty"`    // Optional: The start timestamp in milliseconds.
	End      *int64 `json:"end,omitempty"`      // Optional: The end timestamp in milliseconds.
	Limit    *int   `json:"limit,omitempty"`    // Optional: Limit the number of klines returned.
}

type KlineResult struct {
	Symbol   string     `json:"symbol"`
	Category string     `json:"category"`
	List     [][]string `json:"list"`
}

type ServerTimeResult struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

type OrderBookResult struct {
	S  string     `json:"s"`
	A  [][]string `json:"a"`
	B  [][]string `json:"b"`
	Ts int64      `json:"ts"`
	U  int        `json:"u"`
}

type RiskLimitResult struct {
	Category string        `json:"category"`
	List     []interface{} `json:"list"`
}

type ResendTradeItem struct {
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	Time         string `json:"time"`
	ExecId       string `json:"execId"`
	IsBlockTrade bool   `json:"isBlockTrade"`
}

type DeliveryPriceItem struct {
	Symbol        string `json:"symbol"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryTime  string `json:"deliveryTime"`
}

type HistoricalVolatilityItem struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

type InsuranceItem struct {
	Coin    string `json:"coin"`
	Balance string `json:"balance"`
	Value   string `json:"value"`
}

type OpenHistoryItem struct {
	OpenInterest string `json:"openInterest"`
	Timestamp    string `json:"timestamp"`
}

type FundingRateHistoryItem struct {
	Symbol               string `json:"symbol"`
	FundingRate          string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}

type KlineResponse struct {
	ApiResponse
	Result KlineResult `json:"result"`
}

type ServerTimeResponse struct {
	ApiResponse
	Result ServerTimeResult `json:"result"`
}

type OrderBook struct {
	ApiResponse
	Result OrderBookResult `json:"result"`
}

type RiskLimit struct {
	ApiResponse
	Result RiskLimitResult `json:"result"`
}

type ResendTrade struct {
	ApiResponse
	Result struct {
		Category string            `json:"category"`
		List     []ResendTradeItem `json:"list"`
	} `json:"result"`
}

type DeliveryPrice struct {
	ApiResponse
	Result struct {
		Category       string              `json:"category"`
		NextPageCursor string              `json:"nextPageCursor"`
		List           []DeliveryPriceItem `json:"list"`
	} `json:"result"`
}

type HistoricalVolatility struct {
	ApiResponse
	Result []HistoricalVolatilityItem `json:"result"`
}

type Insurance struct {
	ApiResponse
	Result struct {
		UpdatedTime string          `json:"updatedTime"`
		List        []InsuranceItem `json:"list"`
	} `json:"result"`
}

type OpenHistory struct {
	ApiResponse
	Result struct {
		Symbol         string            `json:"symbol"`
		Category       string            `json:"category"`
		List           []OpenHistoryItem `json:"list"`
		NextPageCursor string            `json:"nextPageCursor"`
	} `json:"result"`
}

type FundingRateHistory struct {
	ApiResponse
	Result struct {
		Category string                   `json:"category"`
		List     []FundingRateHistoryItem `json:"list"`
	} `json:"result"`
}
type APIBaseResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Time       int64       `json:"time"`
	RetExtInfo interface{} `json:"retExtInfo,omitempty"` // Using omitempty since sometimes the field is empty
}

type Announcement struct {
	Total int `json:"total"`
	List  []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Type        struct {
			Title string `json:"title"`
			Key   string `json:"key"`
		} `json:"type"`
		Tags               []string `json:"tags"`
		Url                string   `json:"url"`
		DateTimestamp      int64    `json:"dateTimestamp"`
		StartDateTimestamp int64    `json:"startDateTimestamp"`
		EndDateTimestamp   int64    `json:"endDateTimestamp"`
	} `json:"list"`
}

type AnnouncementsResponse struct {
	APIBaseResponse
	Result Announcement `json:"result"`
}

type Instrument struct {
	Symbol          string `json:"symbol"`
	ContractType    string `json:"contractType"`
	Status          string `json:"status"`
	BaseCoin        string `json:"baseCoin"`
	QuoteCoin       string `json:"quoteCoin"`
	LaunchTime      string `json:"launchTime"`
	DeliveryTime    string `json:"deliveryTime"`
	DeliveryFeeRate string `json:"deliveryFeeRate"`
	PriceScale      string `json:"priceScale"`
	LeverageFilter  struct {
		MinLeverage  string `json:"minLeverage"`
		MaxLeverage  string `json:"maxLeverage"`
		LeverageStep string `json:"leverageStep"`
	} `json:"leverageFilter"`
	PriceFilter struct {
		MinPrice string `json:"minPrice"`
		MaxPrice string `json:"maxPrice"`
		TickSize string `json:"tickSize"`
	} `json:"priceFilter"`
	LotSizeFilter struct {
		MaxOrderQty         string `json:"maxOrderQty"`
		MinOrderQty         string `json:"minOrderQty"`
		QtyStep             string `json:"qtyStep"`
		PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
	} `json:"lotSizeFilter"`
	UnifiedMarginTrade bool   `json:"unifiedMarginTrade"`
	FundingInterval    int    `json:"fundingInterval"`
	SettleCoin         string `json:"settleCoin"`
}

type InstrumentsInfoResponse struct {
	APIBaseResponse
	Result struct {
		Category       string       `json:"category"`
		List           []Instrument `json:"list"`
		NextPageCursor string       `json:"nextPageCursor"`
	} `json:"result"`
}

type TickerInfo struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24H           string `json:"prevPrice24h"`
	Price24HPcnt           string `json:"price24hPcnt"`
	HighPrice24H           string `json:"highPrice24h"`
	LowPrice24H            string `json:"lowPrice24h"`
	PrevPrice1H            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24H            string `json:"turnover24h"`
	Volume24H              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
	Basis                  string `json:"basis"`
}

type TickerResponse struct {
	APIBaseResponse
	Result struct {
		Category string       `json:"category"`
		List     []TickerInfo `json:"list"`
	} `json:"result"`
}

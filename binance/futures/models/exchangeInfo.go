package models

import "time"

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

type AssetInfo struct {
	Asset             string `json:"asset"`
	MarginAvailable   bool   `json:"marginAvailable"`
	AutoAssetExchange string `json:"autoAssetExchange"`
}

type SymbolFilter struct {
	TickSize          string `json:"tickSize,omitempty"`
	MinPrice          string `json:"minPrice,omitempty"`
	MaxPrice          string `json:"maxPrice,omitempty"`
	FilterType        string `json:"filterType"`
	StepSize          string `json:"stepSize,omitempty"`
	MaxQty            string `json:"maxQty,omitempty"`
	MinQty            string `json:"minQty,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Notional          string `json:"notional,omitempty"`
	MultiplierDecimal string `json:"multiplierDecimal,omitempty"`
	MultiplierDown    string `json:"multiplierDown,omitempty"`
	MultiplierUp      string `json:"multiplierUp,omitempty"`
}

type SymbolInfo struct {
	Symbol                string         `json:"symbol"`
	Pair                  string         `json:"pair"`
	ContractType          string         `json:"contractType"`
	DeliveryDate          int64          `json:"deliveryDate"`
	OnboardDate           int64          `json:"onboardDate"`
	Status                string         `json:"status"`
	MaintMarginPercent    string         `json:"maintMarginPercent"`
	RequiredMarginPercent string         `json:"requiredMarginPercent"`
	BaseAsset             string         `json:"baseAsset"`
	QuoteAsset            string         `json:"quoteAsset"`
	MarginAsset           string         `json:"marginAsset"`
	PricePrecision        int            `json:"pricePrecision"`
	QuantityPrecision     int            `json:"quantityPrecision"`
	BaseAssetPrecision    int            `json:"baseAssetPrecision"`
	QuotePrecision        int            `json:"quotePrecision"`
	UnderlyingType        string         `json:"underlyingType"`
	UnderlyingSubType     []interface{}  `json:"underlyingSubType"`
	SettlePlan            int            `json:"settlePlan"`
	TriggerProtect        string         `json:"triggerProtect"`
	LiquidationFee        string         `json:"liquidationFee"`
	MarketTakeBound       string         `json:"marketTakeBound"`
	MaxMoveOrderLimit     int            `json:"maxMoveOrderLimit"`
	Filters               []SymbolFilter `json:"filters"`
	OrderTypes            []string       `json:"orderTypes"`
	TimeInForce           []string       `json:"timeInForce"`
}

type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	FuturesType     string        `json:"futuresType"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Assets          []AssetInfo   `json:"assets"`
	Symbols         []SymbolInfo  `json:"symbols"`
}

type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

func (r ServerTimeResponse) Format(layout string) string {
	return time.Unix(r.ServerTime/1000, 0).Format(layout)
}

package position

// PositionRequestParams represents the query parameters for fetching position information.
type PositionRequestParams struct {
	Category   *string `json:"category"`
	Symbol     *string `json:"symbol"`
	BaseCoin   *string `json:"baseCoin"`
	SettleCoin *string `json:"settleCoin"`
	Limit      *int    `json:"limit"`
	Cursor     *string `json:"cursor"`
}

// PositionResponse represents the response structure for position information.
type PositionResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []PositionDetails `json:"list"`
		NextPageCursor string            `json:"nextPageCursor"`
		Category       string            `json:"category"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// PositionDetails represents the details of a single position.
type PositionDetails struct {
	PositionIdx            int    `json:"positionIdx"`
	RiskID                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	AvgPrice               string `json:"avgPrice"`
	PositionValue          string `json:"positionValue"`
	TradeMode              int    `json:"tradeMode"`
	PositionStatus         string `json:"positionStatus"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	Leverage               string `json:"leverage"`
	PositionBalance        string `json:"positionBalance"`
	MarkPrice              string `json:"markPrice"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	PositionMM             string `json:"positionMM"`
	PositionIM             string `json:"positionIM"`
	TpslMode               string `json:"tpslMode"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	Seq                    int64  `json:"seq"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	MmrSysUpdateTime       string `json:"mmrSysUpdateTime"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
}

// SetLeverageRequest represents the payload for setting leverage.
type SetLeverageRequest struct {
	Category     *string `json:"category"`
	Symbol       *string `json:"symbol"`
	BuyLeverage  *string `json:"buyLeverage"`
	SellLeverage *string `json:"sellLeverage"`
}

type SwitchMarginModeRequest struct {
	Category     *string `json:"category"`
	Symbol       *string `json:"symbol"`
	TradeMode    *int    `json:"tradeMode"`
	BuyLeverage  *string `json:"buyLeverage"`
	SellLeverage *string `json:"sellLeverage"`
}

// SetTPSLModeRequest represents the payload for setting the TP/SL mode.
type SetTPSLModeRequest struct {
	Category *string `json:"category"`
	Symbol   *string `json:"symbol"`
	TPSLMode *string `json:"tpSlMode"` // "Full" or "Partial"
}

// SwitchPositionModeRequest represents the payload for switching the position mode.
type SwitchPositionModeRequest struct {
	Category string  `json:"category"`         // Required: "linear" for USDT Perp, "inverse" for Inverse Futures
	Symbol   *string `json:"symbol,omitempty"` // Optional: Symbol name; either symbol or coin is required
	Coin     *string `json:"coin,omitempty"`   // Optional: Coin; either symbol or coin is required
	Mode     *int    `json:"mode"`             // Required: 0 for Merged Single, 3 for Both Sides
}

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

// SetRiskLimitRequest represents the payload for setting the risk limit of a position.
type SetRiskLimitRequest struct {
	Category    string `json:"category"`    // Required: "linear" or "inverse"
	Symbol      string `json:"symbol"`      // Required: Symbol name
	RiskID      int    `json:"riskId"`      // Required: Risk limit ID
	PositionIdx *int   `json:"positionIdx"` // Optional: Position index (for hedge mode)
}

// SetTradingStopRequest represents the payload for setting trading stops (TP, SL, TS).
type SetTradingStopRequest struct {
	Category     string  `json:"category"`               // Required
	Symbol       string  `json:"symbol"`                 // Required
	TakeProfit   *string `json:"takeProfit,omitempty"`   // Optional, 0 to cancel
	StopLoss     *string `json:"stopLoss,omitempty"`     // Optional, 0 to cancel
	TrailingStop *string `json:"trailingStop,omitempty"` // Optional, 0 to cancel
	TpTriggerBy  *string `json:"tpTriggerBy,omitempty"`  // Optional
	SlTriggerBy  *string `json:"slTriggerBy,omitempty"`  // Optional
	ActivePrice  *string `json:"activePrice,omitempty"`  // Optional
	TPSLMode     string  `json:"tpslMode"`               // Required
	TpSize       *string `json:"tpSize,omitempty"`       // Optional
	SlSize       *string `json:"slSize,omitempty"`       // Optional
	TpLimitPrice *string `json:"tpLimitPrice,omitempty"` // Optional
	SlLimitPrice *string `json:"slLimitPrice,omitempty"` // Optional
	TpOrderType  *string `json:"tpOrderType,omitempty"`  // Optional
	SlOrderType  *string `json:"slOrderType,omitempty"`  // Optional
	PositionIdx  int     `json:"positionIdx"`            // Required
}

// SetAutoAddMarginRequest represents the payload for toggling auto-add-margin.
type SetAutoAddMarginRequest struct {
	Category      string `json:"category"`              // Required: "linear" or "inverse"
	Symbol        string `json:"symbol"`                // Required: Symbol name
	AutoAddMargin int    `json:"autoAddMargin"`         // Required: 0 for off, 1 for on
	PositionIdx   *int   `json:"positionIdx,omitempty"` // Optional: Position index for hedge mode
}

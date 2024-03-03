package position

// RequestParams represents the query parameters for fetching position information.
type RequestParams struct {
	Category   *string `json:"category"`
	Symbol     *string `json:"symbol"`
	BaseCoin   *string `json:"baseCoin"`
	SettleCoin *string `json:"settleCoin"`
	Limit      *int    `json:"limit"`
	Cursor     *string `json:"cursor"`
}

// Response represents the response structure for position information.
type Response struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []Details `json:"list"`
		NextPageCursor string    `json:"nextPageCursor"`
		Category       string    `json:"category"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// Details represent the details of a single position.
type Details struct {
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

// AddReduceMarginRequest represents the payload for adding or reducing margin.
type AddReduceMarginRequest struct {
	Category    string `json:"category"`    // Required: "linear" or "inverse"
	Symbol      string `json:"symbol"`      // Required: Symbol name
	Margin      string `json:"margin"`      // Required: Amount to add (positive) or reduce (negative)
	PositionIdx *int   `json:"positionIdx"` // Optional: Position index for hedge mode
}

// GetClosedPnLRequest represents the query parameters for fetching closed PnL records.
type GetClosedPnLRequest struct {
	Category  string  `json:"category"`            // Required: "linear" or "inverse"
	Symbol    *string `json:"symbol,omitempty"`    // Optional: Symbol name
	StartTime *int64  `json:"startTime,omitempty"` // Optional: The start timestamp (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: The end timestamp (ms)
	Limit     *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Cursor for pagination
}

// ClosedPnLResponse represents the response structure for closed PnL records.
type ClosedPnLResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		NextPageCursor string        `json:"nextPageCursor"`
		Category       string        `json:"category"`
		List           []interface{} `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// MovePositionRequestLeg represents a single leg of a move position request.
type MovePositionRequestLeg struct {
	Category string `json:"category"` // "linear", "spot", "option"
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Side     string `json:"side"` // "Buy" or "Sell"
	Qty      string `json:"qty"`
}

// MovePositionRequest encapsulates the payload for moving positions.
type MovePositionRequest struct {
	FromUID string                   `json:"fromUid"`
	ToUID   string                   `json:"toUid"`
	List    []MovePositionRequestLeg `json:"list"`
}

// MovePositionResponse represents the response from a move position request.
type MovePositionResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		BlockTradeId string `json:"blockTradeId"`
		Status       string `json:"status"`
		RejectParty  string `json:"rejectParty"`
	} `json:"result"`
}

// GetMovePositionHistoryRequest represents the query parameters for fetching move position history.
type GetMovePositionHistoryRequest struct {
	Category     *string `json:"category,omitempty"`     // Optional: Product type
	Symbol       *string `json:"symbol,omitempty"`       // Optional: Symbol name
	StartTime    *int64  `json:"startTime,omitempty"`    // Optional: Start timestamp
	EndTime      *int64  `json:"endTime,omitempty"`      // Optional: End timestamp
	Status       *string `json:"status,omitempty"`       // Optional: Order status
	BlockTradeId *string `json:"blockTradeId,omitempty"` // Optional: Block trade ID
	Limit        *int    `json:"limit,omitempty"`        // Optional: Data size limit per page
	Cursor       *string `json:"cursor,omitempty"`       // Optional: Pagination cursor
}

// MovePositionHistoryEntry represents a single entry in the move position history.
type MovePositionHistoryEntry struct {
	BlockTradeId  string `json:"blockTradeId"`
	Category      string `json:"category"`
	OrderId       string `json:"orderId"`
	UserId        int    `json:"userId"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	ExecFee       string `json:"execFee"`
	Status        string `json:"status"`
	ExecId        string `json:"execId"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
	RejectParty   string `json:"rejectParty"`
}

// GetMovePositionHistoryResponse represents the response from fetching move position history.
type GetMovePositionHistoryResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []MovePositionHistoryEntry `json:"list"`
		NextPageCursor string                     `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

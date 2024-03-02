package position

// PositionRequestParams represents the query parameters for fetching position information.
type PositionRequestParams struct {
	Category   string
	Symbol     string
	BaseCoin   string
	SettleCoin string
	Limit      int
	Cursor     string
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

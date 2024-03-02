package trade

type PlaceOrderRequest struct {
	Category         string  `json:"category"`
	Symbol           string  `json:"symbol"`
	IsLeverage       *int    `json:"isLeverage,omitempty"`
	Side             string  `json:"side"`
	OrderType        string  `json:"orderType"`
	Qty              string  `json:"qty"`
	Price            *string `json:"price,omitempty"`
	TriggerPrice     *string `json:"triggerPrice,omitempty"`
	TriggerDirection *int    `json:"triggerDirection,omitempty"`
	TriggerBy        *string `json:"triggerBy,omitempty"`
	OrderFilter      *string `json:"orderFilter,omitempty"`
	OrderIv          *string `json:"orderIv,omitempty"`
	TimeInForce      string  `json:"timeInForce"`
	PositionIdx      *int    `json:"positionIdx,omitempty"`
	OrderLinkId      string  `json:"orderLinkId"`
	TakeProfit       *string `json:"takeProfit,omitempty"`
	StopLoss         *string `json:"stopLoss,omitempty"`
	TpTriggerBy      *string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy      *string `json:"slTriggerBy,omitempty"`
	ReduceOnly       *bool   `json:"reduceOnly,omitempty"`
	CloseOnTrigger   *bool   `json:"closeOnTrigger,omitempty"`
	SmpType          *string `json:"smpType,omitempty"`
	Mmp              *bool   `json:"mmp,omitempty"`
	TpslMode         *string `json:"tpslMode,omitempty"`
	TpLimitPrice     *string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice     *string `json:"slLimitPrice,omitempty"`
	TpOrderType      *string `json:"tpOrderType,omitempty"`
	SlOrderType      *string `json:"slOrderType,omitempty"`
}

type PlaceOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderID     string `json:"orderId"`
		OrderLinkID string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type AmendOrderRequest struct {
	Category     string  `json:"category"`
	Symbol       string  `json:"symbol"`
	OrderId      *string `json:"orderId,omitempty"`
	OrderLinkId  *string `json:"orderLinkId,omitempty"`
	OrderIv      *string `json:"orderIv,omitempty"`
	TriggerPrice *string `json:"triggerPrice,omitempty"`
	Qty          *string `json:"qty,omitempty"`
	Price        *string `json:"price,omitempty"`
	TpslMode     *string `json:"tpslMode,omitempty"`
	TakeProfit   *string `json:"takeProfit,omitempty"`
	StopLoss     *string `json:"stopLoss,omitempty"`
	TpTriggerBy  *string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy  *string `json:"slTriggerBy,omitempty"`
	TriggerBy    *string `json:"triggerBy,omitempty"`
	TpLimitPrice *string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice *string `json:"slLimitPrice,omitempty"`
}
type AmendOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
type CancelOrderRequest struct {
	Category    string  `json:"category"`
	Symbol      string  `json:"symbol"`
	OrderId     *string `json:"orderId,omitempty"`
	OrderLinkId *string `json:"orderLinkId,omitempty"`
	OrderFilter *string `json:"orderFilter,omitempty"` // Valid for spot only
}
type CancelOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
type GetOpenOrdersRequest struct {
	Category    string
	Symbol      *string
	BaseCoin    *string
	SettleCoin  *string
	OrderId     *string
	OrderLinkId *string
	OpenOnly    *int
	OrderFilter *string
	Limit       *int
	Cursor      *string
}
type GetOpenOrdersResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []OrderDetails `json:"list"`
		NextPageCursor string         `json:"nextPageCursor"`
		Category       string         `json:"category"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}

type OrderDetails struct {
	OrderId            string `json:"orderId"`
	OrderLinkId        string `json:"orderLinkId"`
	BlockTradeId       string `json:"blockTradeId"`
	Symbol             string `json:"symbol"`
	Price              string `json:"price"`
	Qty                string `json:"qty"`
	Side               string `json:"side"`
	IsLeverage         string `json:"isLeverage"`
	PositionIdx        int    `json:"positionIdx"`
	OrderStatus        string `json:"orderStatus"`
	CreateType         string `json:"createType"`
	CancelType         string `json:"cancelType"`
	RejectReason       string `json:"rejectReason"`
	AvgPrice           string `json:"avgPrice"`
	LeavesQty          string `json:"leavesQty"`
	LeavesValue        string `json:"leavesValue"`
	CumExecQty         string `json:"cumExecQty"`
	CumExecValue       string `json:"cumExecValue"`
	CumExecFee         string `json:"cumExecFee"`
	TimeInForce        string `json:"timeInForce"`
	OrderType          string `json:"orderType"`
	StopOrderType      string `json:"stopOrderType"`
	OrderIv            string `json:"orderIv"`
	MarketUnit         string `json:"marketUnit"`
	TriggerPrice       string `json:"triggerPrice"`
	TakeProfit         string `json:"takeProfit"`
	StopLoss           string `json:"stopLoss"`
	TpslMode           string `json:"tpslMode"`
	OcoTriggerType     string `json:"ocoTriggerType"`
	TpLimitPrice       string `json:"tpLimitPrice"`
	SlLimitPrice       string `json:"slLimitPrice"`
	TpTriggerBy        string `json:"tpTriggerBy"`
	SlTriggerBy        string `json:"slTriggerBy"`
	TriggerDirection   int    `json:"triggerDirection"`
	TriggerBy          string `json:"triggerBy"`
	LastPriceOnCreated string `json:"lastPriceOnCreated"`
	ReduceOnly         bool   `json:"reduceOnly"`
	CloseOnTrigger     bool   `json:"closeOnTrigger"`
	PlaceType          string `json:"placeType"`
	SmpType            string `json:"smpType"`
	SmpGroup           int    `json:"smpGroup"`
	SmpOrderId         string `json:"smpOrderId"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
}
type CancelAllOrdersRequest struct {
	Category      string  `json:"category"`
	Symbol        *string `json:"symbol,omitempty"`
	BaseCoin      *string `json:"baseCoin,omitempty"`
	SettleCoin    *string `json:"settleCoin,omitempty"`
	OrderFilter   *string `json:"orderFilter,omitempty"`
	StopOrderType *string `json:"stopOrderType,omitempty"`
}
type CancelAllOrdersResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			OrderId     string `json:"orderId"`
			OrderLinkId string `json:"orderLinkId"`
		} `json:"list"`
		Success string `json:"success"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
type GetOrderHistoryRequest struct {
	Category    string
	Symbol      *string
	BaseCoin    *string
	SettleCoin  *string
	OrderId     *string
	OrderLinkId *string
	OrderFilter *string
	OrderStatus *string
	StartTime   *int64
	EndTime     *int64
	Limit       *int
	Cursor      *string
}
type GetOrderHistoryResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []OrderDetails `json:"list"`
		NextPageCursor string         `json:"nextPageCursor"`
		Category       string         `json:"category"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
type GetTradeHistoryRequest struct {
	Category    string
	Symbol      *string
	OrderId     *string
	OrderLinkId *string
	BaseCoin    *string
	StartTime   *int64
	EndTime     *int64
	ExecType    *string
	Limit       *int
	Cursor      *string
}
type GetTradeHistoryResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []TradeDetails `json:"list"`
		NextPageCursor string         `json:"nextPageCursor"`
		Category       string         `json:"category"`
	} `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}

type TradeDetails struct {
	Symbol          string `json:"symbol"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	CreateType      string `json:"createType"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecFee         string `json:"execFee"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	FeeCurrency     string `json:"feeCurrency"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int64  `json:"seq"`
}

type BatchPlaceOrderRequest struct {
	Category string         `json:"category"`
	Request  []OrderRequest `json:"request"`
}

type OrderRequest struct {
	Symbol           string  `json:"symbol"`
	Side             string  `json:"side"`
	OrderType        string  `json:"orderType"`
	Qty              string  `json:"qty"`
	Price            *string `json:"price,omitempty"`
	TriggerDirection *int    `json:"triggerDirection,omitempty"`
	TriggerPrice     *string `json:"triggerPrice,omitempty"`
	TriggerBy        *string `json:"triggerBy,omitempty"`
	OrderIv          *string `json:"orderIv,omitempty"`
	TimeInForce      *string `json:"timeInForce,omitempty"`
	PositionIdx      *int    `json:"positionIdx,omitempty"`
	OrderLinkId      *string `json:"orderLinkId,omitempty"`
	TakeProfit       *string `json:"takeProfit,omitempty"`
	StopLoss         *string `json:"stopLoss,omitempty"`
	TpTriggerBy      *string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy      *string `json:"slTriggerBy,omitempty"`
	ReduceOnly       *bool   `json:"reduceOnly,omitempty"`
	CloseOnTrigger   *bool   `json:"closeOnTrigger,omitempty"`
	SmpType          *string `json:"smpType,omitempty"`
	Mmp              *bool   `json:"mmp,omitempty"`
	TpslMode         *string `json:"tpslMode,omitempty"`
	TpLimitPrice     *string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice     *string `json:"slLimitPrice,omitempty"`
	TpOrderType      *string `json:"tpOrderType,omitempty"`
	SlOrderType      *string `json:"slOrderType,omitempty"`
}
type BatchPlaceOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Category    string `json:"category"`
			Symbol      string `json:"symbol"`
			OrderId     string `json:"orderId"`
			OrderLinkId string `json:"orderLinkId"`
			CreateAt    string `json:"createAt"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}
type BatchAmendOrderRequest struct {
	Category string              `json:"category"`
	Request  []AmendOrderRequest `json:"request"`
}

type BatchAmendOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Category    string `json:"category"`
			Symbol      string `json:"symbol"`
			OrderId     string `json:"orderId"`
			OrderLinkId string `json:"orderLinkId"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}
type BatchCancelOrderRequest struct {
	Category string               `json:"category"`
	Request  []CancelOrderRequest `json:"request"`
}
type BatchCancelOrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Category    string `json:"category"`
			Symbol      string `json:"symbol"`
			OrderId     string `json:"orderId"`
			OrderLinkId string `json:"orderLinkId"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

// Response struct for parsing the API response
type BorrowQuotaResponse struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     BorrowQuotaResult `json:"result"`
	RetExtInfo interface{}       `json:"retExtInfo"`
	Time       int64             `json:"time"`
}

// BorrowQuotaResult holds the specific data of interest from the response
type BorrowQuotaResult struct {
	Symbol             string `json:"symbol"`
	Side               string `json:"side"`
	MaxTradeQty        string `json:"maxTradeQty"`
	MaxTradeAmount     string `json:"maxTradeAmount"`
	SpotMaxTradeQty    string `json:"spotMaxTradeQty"`
	SpotMaxTradeAmount string `json:"spotMaxTradeAmount"`
	BorrowCoin         string `json:"borrowCoin"`
}

package trade

import (
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

func ConvertPlaceOrderRequestToParams(req *PlaceOrderRequest) client.Params {
	params := client.Params{}

	params["category"] = req.Category
	params["symbol"] = req.Symbol
	if req.IsLeverage != nil {
		params["isLeverage"] = strconv.Itoa(*req.IsLeverage)
	}
	params["side"] = req.Side
	params["orderType"] = req.OrderType
	params["qty"] = req.Qty
	if req.Price != nil {
		params["price"] = *req.Price
	}
	if req.TriggerPrice != nil {
		params["triggerPrice"] = *req.TriggerPrice
	}
	if req.TriggerDirection != nil {
		params["triggerDirection"] = strconv.Itoa(*req.TriggerDirection)
	}
	if req.TriggerBy != nil {
		params["triggerBy"] = *req.TriggerBy
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}
	if req.OrderIv != nil {
		params["orderIv"] = *req.OrderIv
	}
	params["timeInForce"] = req.TimeInForce
	if req.PositionIdx != nil {
		params["positionIdx"] = strconv.Itoa(*req.PositionIdx)
	}
	params["orderLinkId"] = req.OrderLinkId
	if req.TakeProfit != nil {
		params["takeProfit"] = *req.TakeProfit
	}
	if req.StopLoss != nil {
		params["stopLoss"] = *req.StopLoss
	}
	if req.TpTriggerBy != nil {
		params["tpTriggerBy"] = *req.TpTriggerBy
	}
	if req.SlTriggerBy != nil {
		params["slTriggerBy"] = *req.SlTriggerBy
	}
	if req.ReduceOnly != nil {
		params["reduceOnly"] = strconv.FormatBool(*req.ReduceOnly)
	}
	if req.CloseOnTrigger != nil {
		params["closeOnTrigger"] = strconv.FormatBool(*req.CloseOnTrigger)
	}
	if req.SmpType != nil {
		params["smpType"] = *req.SmpType
	}
	if req.Mmp != nil {
		params["mmp"] = strconv.FormatBool(*req.Mmp)
	}
	if req.TpslMode != nil {
		params["tpslMode"] = *req.TpslMode
	}
	if req.TpLimitPrice != nil {
		params["tpLimitPrice"] = *req.TpLimitPrice
	}
	if req.SlLimitPrice != nil {
		params["slLimitPrice"] = *req.SlLimitPrice
	}
	if req.TpOrderType != nil {
		params["tpOrderType"] = *req.TpOrderType
	}
	if req.SlOrderType != nil {
		params["slOrderType"] = *req.SlOrderType
	}

	return params
}

// ConvertAmendOrderRequestToParams converts an AmendOrderRequest to a client.Params map.
func ConvertAmendOrderRequestToParams(req *AmendOrderRequest) client.Params {
	params := client.Params{}

	// Mandatory fields
	params["category"] = req.Category
	params["symbol"] = req.Symbol

	// Optional fields
	if req.OrderId != nil {
		params["orderId"] = *req.OrderId
	}
	if req.OrderLinkId != nil {
		params["orderLinkId"] = *req.OrderLinkId
	}
	if req.OrderIv != nil {
		params["orderIv"] = *req.OrderIv
	}
	if req.TriggerPrice != nil {
		params["triggerPrice"] = *req.TriggerPrice
	}
	if req.Qty != nil {
		params["qty"] = *req.Qty
	}
	if req.Price != nil {
		params["price"] = *req.Price
	}
	if req.TpslMode != nil {
		params["tpslMode"] = *req.TpslMode
	}
	if req.TakeProfit != nil {
		params["takeProfit"] = *req.TakeProfit
	}
	if req.StopLoss != nil {
		params["stopLoss"] = *req.StopLoss
	}
	if req.TpTriggerBy != nil {
		params["tpTriggerBy"] = *req.TpTriggerBy
	}
	if req.SlTriggerBy != nil {
		params["slTriggerBy"] = *req.SlTriggerBy
	}
	if req.TriggerBy != nil {
		params["triggerBy"] = *req.TriggerBy
	}
	if req.TpLimitPrice != nil {
		params["tpLimitPrice"] = *req.TpLimitPrice
	}
	if req.SlLimitPrice != nil {
		params["slLimitPrice"] = *req.SlLimitPrice
	}

	return params
}

func ConvertCancelOrderRequestToParams(req *CancelOrderRequest) client.Params {
	params := client.Params{
		"category": req.Category,
		"symbol":   req.Symbol,
	}

	if req.OrderId != nil {
		params["orderId"] = *req.OrderId
	}
	if req.OrderLinkId != nil {
		params["orderLinkId"] = *req.OrderLinkId
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}

	return params
}

func ConvertGetOpenOrdersRequestToParams(req *GetOpenOrdersRequest) client.Params {
	params := client.Params{
		"category": req.Category, // Required field
	}

	// Check and add optional fields only if they are not nil
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.BaseCoin != nil {
		params["baseCoin"] = *req.BaseCoin
	}
	if req.SettleCoin != nil {
		params["settleCoin"] = *req.SettleCoin
	}
	if req.OrderId != nil {
		params["orderId"] = *req.OrderId
	}
	if req.OrderLinkId != nil {
		params["orderLinkId"] = *req.OrderLinkId
	}
	if req.OpenOnly != nil {
		params["openOnly"] = strconv.Itoa(*req.OpenOnly) // Convert integer to string
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}
	if req.Limit != nil {
		params["limit"] = strconv.Itoa(*req.Limit) // Convert integer to string for limit
	}
	if req.Cursor != nil {
		params["cursor"] = *req.Cursor
	}

	return params
}
func ConvertCancelAllOrdersRequestToParams(req *CancelAllOrdersRequest) client.Params {
	params := client.Params{
		"category": req.Category,
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.BaseCoin != nil {
		params["baseCoin"] = *req.BaseCoin
	}
	if req.SettleCoin != nil {
		params["settleCoin"] = *req.SettleCoin
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}
	if req.StopOrderType != nil {
		params["stopOrderType"] = *req.StopOrderType
	}
	return params
}
func ConvertGetOrderHistoryRequestToParams(req *GetOrderHistoryRequest) client.Params {
	params := client.Params{
		"category": req.Category,
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.BaseCoin != nil {
		params["baseCoin"] = *req.BaseCoin
	}
	if req.SettleCoin != nil {
		params["settleCoin"] = *req.SettleCoin
	}
	if req.OrderId != nil {
		params["orderId"] = *req.OrderId
	}
	if req.OrderLinkId != nil {
		params["orderLinkId"] = *req.OrderLinkId
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}
	if req.OrderStatus != nil {
		params["orderStatus"] = *req.OrderStatus
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.Limit != nil {
		params["limit"] = strconv.Itoa(*req.Limit)
	}
	if req.Cursor != nil {
		params["cursor"] = *req.Cursor
	}
	return params
}
func ConvertGetTradeHistoryRequestToParams(req *GetTradeHistoryRequest) client.Params {
	params := client.Params{
		"category": req.Category,
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.OrderId != nil {
		params["orderId"] = *req.OrderId
	}
	if req.OrderLinkId != nil {
		params["orderLinkId"] = *req.OrderLinkId
	}
	if req.BaseCoin != nil {
		params["baseCoin"] = *req.BaseCoin
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.ExecType != nil {
		params["execType"] = *req.ExecType
	}
	if req.Limit != nil {
		params["limit"] = strconv.Itoa(*req.Limit)
	}
	if req.Cursor != nil {
		params["cursor"] = *req.Cursor
	}
	return params
}

package trade

import (
	"fmt"
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

// ConvertAmendOrderRequestToParams converts an AmendOrderRequest to a client.Params map.
func ConvertAmendOrderRequestToParams(req *AmendOrderRequest) client.Params {
	params := client.Params{}

	// Mandatory fields
	params["category"] = req.Category
	params["symbol"] = req.Symbol

	// Optional fields
	if req.OrderID != nil {
		params["orderId"] = *req.OrderID
	}
	if req.OrderLinkID != nil {
		params["orderLinkId"] = *req.OrderLinkID
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

	if req.OrderID != nil {
		params["orderId"] = *req.OrderID
	}
	if req.OrderLinkID != nil {
		params["orderLinkId"] = *req.OrderLinkID
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
	if req.OrderID != nil {
		params["orderId"] = *req.OrderID
	}
	if req.OrderLinkID != nil {
		params["orderLinkId"] = *req.OrderLinkID
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
	if req.OrderID != nil {
		params["orderId"] = *req.OrderID
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
	if req.OrderID != nil {
		params["orderId"] = *req.OrderID
	}
	if req.OrderLinkID != nil {
		params["orderLinkId"] = *req.OrderLinkID
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

func ConvertBatchPlaceOrderRequestToParams(req *BatchPlaceOrderRequest) client.Params {
	params := client.Params{}
	params["category"] = req.Category

	for i, order := range req.Request {
		prefix := fmt.Sprintf("request[%d].", i)
		params[prefix+"symbol"] = order.Symbol
		params[prefix+"side"] = order.Side
		params[prefix+"orderType"] = order.OrderType
		params[prefix+"qty"] = order.Qty

		if order.Price != nil {
			params[prefix+"price"] = *order.Price
		}
		if order.TriggerDirection != nil {
			params[prefix+"triggerDirection"] = strconv.Itoa(*order.TriggerDirection)
		}
		if order.TriggerPrice != nil {
			params[prefix+"triggerPrice"] = *order.TriggerPrice
		}
		if order.TriggerBy != nil {
			params[prefix+"triggerBy"] = *order.TriggerBy
		}
		if order.OrderIv != nil {
			params[prefix+"orderIv"] = *order.OrderIv
		}
		if order.TimeInForce != nil {
			params[prefix+"timeInForce"] = *order.TimeInForce
		}
		if order.PositionIdx != nil {
			params[prefix+"positionIdx"] = strconv.Itoa(*order.PositionIdx)
		}
		if order.OrderLinkID != nil {
			params[prefix+"orderLinkId"] = *order.OrderLinkID
		}
		if order.TakeProfit != nil {
			params[prefix+"takeProfit"] = *order.TakeProfit
		}
		if order.StopLoss != nil {
			params[prefix+"stopLoss"] = *order.StopLoss
		}
		if order.TpTriggerBy != nil {
			params[prefix+"tpTriggerBy"] = *order.TpTriggerBy
		}
		if order.SlTriggerBy != nil {
			params[prefix+"slTriggerBy"] = *order.SlTriggerBy
		}
		if order.ReduceOnly != nil {
			params[prefix+"reduceOnly"] = strconv.FormatBool(*order.ReduceOnly)
		}
		if order.CloseOnTrigger != nil {
			params[prefix+"closeOnTrigger"] = strconv.FormatBool(*order.CloseOnTrigger)
		}
		if order.SmpType != nil {
			params[prefix+"smpType"] = *order.SmpType
		}
		if order.Mmp != nil {
			params[prefix+"mmp"] = strconv.FormatBool(*order.Mmp)
		}
		if order.TpslMode != nil {
			params[prefix+"tpslMode"] = *order.TpslMode
		}
		if order.TpLimitPrice != nil {
			params[prefix+"tpLimitPrice"] = *order.TpLimitPrice
		}
		if order.SlLimitPrice != nil {
			params[prefix+"slLimitPrice"] = *order.SlLimitPrice
		}
		if order.TpOrderType != nil {
			params[prefix+"tpOrderType"] = *order.TpOrderType
		}
		if order.SlOrderType != nil {
			params[prefix+"slOrderType"] = *order.SlOrderType
		}
	}

	return params
}
func ConvertBatchAmendOrderRequestToParams(req *BatchAmendOrderRequest) client.Params {
	params := client.Params{}
	params["category"] = req.Category

	for i, order := range req.Request {
		base := fmt.Sprintf("request[%d].", i)
		params[base+"symbol"] = order.Symbol

		if order.OrderID != nil {
			params[base+"orderId"] = *order.OrderID
		}
		if order.OrderLinkID != nil {
			params[base+"orderLinkId"] = *order.OrderLinkID
		}
		if order.OrderIv != nil {
			params[base+"orderIv"] = *order.OrderIv
		}
		if order.TriggerPrice != nil {
			params[base+"triggerPrice"] = *order.TriggerPrice
		}
		if order.Qty != nil {
			params[base+"qty"] = *order.Qty
		}
		if order.Price != nil {
			params[base+"price"] = *order.Price
		}
		if order.TpslMode != nil {
			params[base+"tpslMode"] = *order.TpslMode
		}
		if order.TakeProfit != nil {
			params[base+"takeProfit"] = *order.TakeProfit
		}
		if order.StopLoss != nil {
			params[base+"stopLoss"] = *order.StopLoss
		}
		if order.TpTriggerBy != nil {
			params[base+"tpTriggerBy"] = *order.TpTriggerBy
		}
		if order.SlTriggerBy != nil {
			params[base+"slTriggerBy"] = *order.SlTriggerBy
		}
		if order.TriggerBy != nil {
			params[base+"triggerBy"] = *order.TriggerBy
		}
		if order.TpLimitPrice != nil {
			params[base+"tpLimitPrice"] = *order.TpLimitPrice
		}
		if order.SlLimitPrice != nil {
			params[base+"slLimitPrice"] = *order.SlLimitPrice
		}
	}

	return params
}

// ConvertBatchCancelOrderRequestToParams creates a Params representation for logging/debugging.
func ConvertBatchCancelOrderRequestToParams(req *BatchCancelOrderRequest) client.Params {
	params := client.Params{}
	params["category"] = req.Category

	for i, orderReq := range req.Request {
		index := fmt.Sprintf("request[%d]", i)
		params[index+".symbol"] = orderReq.Symbol

		if orderReq.OrderID != nil {
			params[index+".orderId"] = *orderReq.OrderID
		}
		if orderReq.OrderLinkID != nil {
			params[index+".orderLinkId"] = *orderReq.OrderLinkID
		}
		if orderReq.OrderFilter != nil {
			params[index+".orderFilter"] = *orderReq.OrderFilter
		}
	}

	return params
}

// NewDCPParams creates a new Params map for setting the DCP time window.
func NewDCPParams(timeWindow int) client.Params {
	params := make(client.Params)
	params["timeWindow"] = strconv.Itoa(timeWindow) // Convert int to string
	return params
}

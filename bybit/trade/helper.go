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

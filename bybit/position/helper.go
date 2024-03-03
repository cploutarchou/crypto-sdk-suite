package position

import (
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

// ConvertPositionRequestParams prepares the request parameters for fetching position info.
func ConvertPositionRequestParams(params *RequestParams) client.Params {
	paramsMap := make(client.Params)
	if params.Category != nil {
		paramsMap["category"] = *params.Category
	}
	if params.Symbol != nil {
		paramsMap["symbol"] = *params.Symbol
	}
	if params.BaseCoin != nil {
		paramsMap["baseCoin"] = *params.BaseCoin
	}
	if params.SettleCoin != nil {
		paramsMap["settleCoin"] = *params.SettleCoin
	}
	if params.Limit != nil {
		paramsMap["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.Cursor != nil {
		paramsMap["cursor"] = *params.Cursor
	}

	return paramsMap
}

func ConvertSetLeverageRequestToParams(req *SetLeverageRequest) client.Params {
	params := make(client.Params)

	if req.Category != nil {
		params["category"] = *req.Category
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.BuyLeverage != nil {
		params["buyLeverage"] = *req.BuyLeverage
	}
	if req.SellLeverage != nil {
		params["sellLeverage"] = *req.SellLeverage
	}
	return params
}

func ConvertSwitchMarginModeRequestToParams(req *SwitchMarginModeRequest) client.Params {
	params := make(client.Params)

	if req.Category != nil {
		params["category"] = *req.Category
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.TradeMode != nil {
		params["tradeMode"] = strconv.Itoa(*req.TradeMode)
	}
	if req.BuyLeverage != nil {
		params["buyLeverage"] = *req.BuyLeverage
	}
	if req.SellLeverage != nil {
		params["sellLeverage"] = *req.SellLeverage
	}
	return params
}
func ConvertSetTPSLModeRequestToParams(req *SetTPSLModeRequest) client.Params {
	params := make(client.Params)
	if req.Category != nil {
		params["category"] = *req.Category
	}

	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}

	if req.TPSLMode != nil {
		params["tpslMode"] = *req.TPSLMode
	}
	return params
}
func ConvertSwitchPositionModeRequestToParams(req *SwitchPositionModeRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.Coin != nil {
		params["coin"] = *req.Coin
	}
	if req.Mode != nil {
		params["positionMode"] = strconv.Itoa(*req.Mode)
	}
	return params
}
func ConvertSetRiskLimitRequestToParams(req *SetRiskLimitRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != "" {
		params["symbol"] = req.Symbol
	}

	if req.RiskID > 0 {
		params["riskId"] = strconv.Itoa(req.RiskID)
	}
	if req.PositionIdx != nil {
		params["positionIdx"] = strconv.Itoa(*req.PositionIdx)
	}
	return params
}

func ConvertSetTradingStopRequestToParams(req *SetTradingStopRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != "" {
		params["symbol"] = req.Symbol
	}

	if req.TakeProfit != nil {
		params["takeProfit"] = *req.TakeProfit
	}

	if req.StopLoss != nil {
		params["stopLoss"] = *req.StopLoss
	}

	if req.TrailingStop != nil {
		params["trailingStop"] = *req.TrailingStop
	}

	if req.TpTriggerBy != nil {
		params["tpTriggerBy"] = *req.TpTriggerBy
	}

	if req.SlTriggerBy != nil {
		params["slTriggerBy"] = *req.SlTriggerBy
	}

	if req.ActivePrice != nil {
		params["activePrice"] = *req.ActivePrice
	}

	if req.TPSLMode != "" {
		params["tpslMode"] = req.TPSLMode
	}

	if req.TpSize != nil {
		params["tpSize"] = *req.TpSize
	}
	if req.SlSize != nil {
		params["slSize"] = *req.SlSize
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
	if req.PositionIdx > 0 {
		params["positionIdx"] = strconv.Itoa(req.PositionIdx)
	}
	return params
}
func ConvertSetAutoAddMarginRequestToParams(req *SetAutoAddMarginRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != "" {
		params["symbol"] = req.Symbol
	}
	if req.AutoAddMargin > 0 {
		params["autoAddMargin"] = strconv.Itoa(req.AutoAddMargin)
	}
	if req.PositionIdx != nil {
		params["positionIdx"] = strconv.Itoa(*req.PositionIdx)
	}
	return params

}
func ConvertAddReduceMarginRequestToParams(req *AddReduceMarginRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != "" {
		params["symbol"] = req.Symbol
	}
	if req.Margin != "" {
		params["autoAddMargin"] = req.Margin
	}
	if req.PositionIdx != nil {
		params["positionIdx"] = strconv.Itoa(*req.PositionIdx)
	}
	return params
}
func ConvertGetClosedPnLRequestToParams(req *GetClosedPnLRequest) client.Params {
	params := make(client.Params)
	if req.Category != "" {
		params["category"] = req.Category
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
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
func ConvertMovePositionRequestToParams(req *MovePositionRequest) client.Params {
	params := make(client.Params)
	if req.FromUID != "" {
		params["fromUid"] = req.FromUID
	}

	if req.ToUID != "" {
		params["toUid"] = req.ToUID
	}
	if len(req.List) > 0 {
		params["list"] = req.List
	}
	return params

}
func ConvertGetMovePositionHistoryRequestToParams(req *GetMovePositionHistoryRequest) client.Params {
	params := make(client.Params)
	if req.Category != nil {
		params["category"] = *req.Category
	}
	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.Status != nil {
		params["status"] = *req.Status
	}
	if req.BlockTradeId != nil {
		params["blockTradeId"] = *req.BlockTradeId
	}
	if req.Limit != nil {
		params["limit"] = strconv.Itoa(*req.Limit)
	}
	if req.Cursor != nil {
		params["cursor"] = *req.Cursor
	}
	return params
}

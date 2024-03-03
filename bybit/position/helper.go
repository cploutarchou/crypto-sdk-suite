package position

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"strconv"
)

// ConvertPositionRequestParams prepares the request parameters for fetching position info.
func ConvertPositionRequestParams(params *PositionRequestParams) client.Params {
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

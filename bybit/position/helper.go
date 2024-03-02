package position

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"strconv"
)

// PreparePositionRequestParams prepares the request parameters for fetching position info.
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

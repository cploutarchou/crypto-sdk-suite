package trade

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Trade interface {
	PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error)
	AmendOrder(req *AmendOrderRequest) (*AmendOrderResponse, error)
	CancelOrder(req *CancelOrderRequest) (*CancelOrderResponse, error)
	GetOpenOrders(req *GetOpenOrdersRequest) (*GetOpenOrdersResponse, error)
	CancelAllOrders(req *CancelAllOrdersRequest) (*CancelAllOrdersResponse, error)
	GetOrderHistory(req *GetOrderHistoryRequest) (*GetOrderHistoryResponse, error)
	GetTradeHistory(req *GetTradeHistoryRequest) (*GetTradeHistoryResponse, error)
	BatchPlaceOrder(req *BatchPlaceOrderRequest) (*BatchPlaceOrderResponse, error)
	GetBorrowQuotaSpot(symbol, side string) (*BorrowQuotaResponse, error)
}

type tradeImpl struct {
	client *client.Client
}

func New(c *client.Client) Trade {
	return &tradeImpl{client: c}
}

func (t *tradeImpl) PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	params := ConvertPlaceOrderRequestToParams(req)
	res, err := t.client.Post("/v5/order/create", params)
	if err != nil {
		return nil, err
	}

	var placeOrderResponse PlaceOrderResponse
	err = res.Unmarshal(&placeOrderResponse)
	if err != nil {
		return nil, err
	}
	if placeOrderResponse.RetCode != 0 {
		return &placeOrderResponse, fmt.Errorf("API returned error: %s", placeOrderResponse.RetMsg)
	}
	return &placeOrderResponse, nil
}

func ConvertPlaceOrderRequestToParams(req *PlaceOrderRequest) client.Params {
	params := client.Params{
		"category":    req.Category,
		"symbol":      req.Symbol,
		"side":        req.Side,
		"orderType":   req.OrderType,
		"qty":         req.Qty,
		"orderLinkId": req.OrderLinkId,
	}

	if req.Price != "" {
		params["price"] = req.Price
	}

	optionalParams := map[string]interface{}{
		"isLeverage":       req.IsLeverage,
		"triggerPrice":     req.TriggerPrice,
		"triggerDirection": req.TriggerDirection,
		"triggerBy":        req.TriggerBy,
		"orderFilter":      req.OrderFilter,
		"orderIv":          req.OrderIv,
		"timeInForce":      req.TimeInForce,
		"positionIdx":      req.PositionIdx,
		"takeProfit":       req.TakeProfit,
		"stopLoss":         req.StopLoss,
		"tpTriggerBy":      req.TpTriggerBy,
		"slTriggerBy":      req.SlTriggerBy,
		"reduceOnly":       req.ReduceOnly,
		"closeOnTrigger":   req.CloseOnTrigger,
		"smpType":          req.SmpType,
		"mmp":              req.Mmp,
		"tpslMode":         req.TpslMode,
		"tpLimitPrice":     req.TpLimitPrice,
		"slLimitPrice":     req.SlLimitPrice,
		"tpOrderType":      req.TpOrderType,
		"slOrderType":      req.SlOrderType,
	}

	for k, v := range optionalParams {
		if v != nil {
			switch value := v.(type) {
			case *string:
				params[k] = *value
			case *int:
				params[k] = strconv.Itoa(*value)
			case *bool:
				params[k] = strconv.FormatBool(*value)
			case int:
				params[k] = strconv.Itoa(value)
			case bool:
				params[k] = strconv.FormatBool(value)
			}
		}
	}

	return params
}

func (t *tradeImpl) AmendOrder(req *AmendOrderRequest) (*AmendOrderResponse, error) {
	params := ConvertAmendOrderRequestToParams(req)
	res, err := t.client.Post("/v5/order/amend", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	var response AmendOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) CancelOrder(req *CancelOrderRequest) (*CancelOrderResponse, error) {
	params := ConvertCancelOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response CancelOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetOpenOrders(req *GetOpenOrdersRequest) (*GetOpenOrdersResponse, error) {
	queryParams := ConvertGetOpenOrdersRequestToParams(req)

	// Assuming the client.Get method constructs the query string from the provided params and sends a GET request.
	resBytes, err := t.client.Get("/v5/order/realtime", queryParams)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response GetOpenOrdersResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) CancelAllOrders(req *CancelAllOrdersRequest) (*CancelAllOrdersResponse, error) {
	params := ConvertCancelAllOrdersRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel-all", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response CancelAllOrdersResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetOrderHistory(req *GetOrderHistoryRequest) (*GetOrderHistoryResponse, error) {
	queryParams := ConvertGetOrderHistoryRequestToParams(req)

	// Assuming the client.Get method constructs the query string from the provided params and sends a GET request.
	resBytes, err := t.client.Get("/v5/order/history", queryParams)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response GetOrderHistoryResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetTradeHistory(req *GetTradeHistoryRequest) (*GetTradeHistoryResponse, error) {
	queryParams := ConvertGetTradeHistoryRequestToParams(req)

	// Assuming the client.Get method constructs the query string from the provided params and sends a GET request.
	resBytes, err := t.client.Get("/v5/execution/list", queryParams)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response GetTradeHistoryResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) BatchPlaceOrder(req *BatchPlaceOrderRequest) (*BatchPlaceOrderResponse, error) {
	params := ConvertBatchPlaceOrderRequestToParams(req)
	resBytes, err := t.client.Post("/v5/order/create-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchPlaceOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}

func (t *tradeImpl) BatchAmendOrder(req *BatchAmendOrderRequest) (*BatchAmendOrderResponse, error) {
	params := ConvertBatchAmendOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/amend-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchAmendOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) BatchCancelOrder(req *BatchCancelOrderRequest) (*BatchCancelOrderResponse, error) {
	params := ConvertBatchCancelOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchCancelOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetBorrowQuotaSpot(symbol, side string) (*BorrowQuotaResponse, error) {
	params := client.Params{
		"category": "spot",
		"symbol":   symbol,
		"side":     side,
	}
	resBytes, err := t.client.Get("/v5/order/spot-borrow-check", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var response BorrowQuotaResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Check for API error
	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) SetDisconnectCancelAll(req *SetDisconnectCancelAllRequest) (*APIResponse, error) {
	dcpRequest := NewDCPParams(req.TimeWindow)

	// Send POST request to the Bybit API
	responseBody, err := t.client.Post("/v5/order/disconnected-cancel-all", dcpRequest)
	if err != nil {
		return nil, fmt.Errorf("error sending request to API: %w", err)
	}
	data, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var response APIResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Check for API error
	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}

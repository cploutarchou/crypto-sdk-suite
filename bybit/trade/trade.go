package trade

import (
	"encoding/json"
	"fmt"

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
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	var placeOrderResponse PlaceOrderResponse
	err = json.Unmarshal(data, &placeOrderResponse)
	if err != nil {
		return nil, err
	}
	if placeOrderResponse.RetCode != 0 {
		return &placeOrderResponse, fmt.Errorf("API returned error: %s", placeOrderResponse.RetMsg)
	}
	return &placeOrderResponse, nil
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

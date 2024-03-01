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

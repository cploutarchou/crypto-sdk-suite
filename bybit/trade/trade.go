package trade

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Trade interface {
	PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error)
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

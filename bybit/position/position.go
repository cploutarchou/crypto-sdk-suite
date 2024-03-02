package position

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Position interface {
	GetPositionInfo(params client.Params) (*PositionResponse, error)
}
type impl struct {
	client *client.Client
}

func New(c *client.Client) Position {
	return &impl{client: c}
}

// GetPositionInfo fetches position information from Bybit.
func (i *impl) GetPositionInfo(params client.Params) (*PositionResponse, error) {
	response, err := i.client.Get("/v5/position/list", params)
	if err != nil {
		return nil, fmt.Errorf("error fetching position info: %v", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse PositionResponse
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing position info response: %v", err)
	}

	return &positionResponse, nil
}

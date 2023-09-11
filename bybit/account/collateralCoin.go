package account

import (
	"errors"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type CollateralCoin struct {
	client *client.Client
}

func NewSetCollateralCoin(c *client.Client) *CollateralCoin {
	return &CollateralCoin{client: c}
}

func (s *CollateralCoin) Set(coin string, collateralSwitch CollateralSwitch) (*CollateralInfoResponse, error) {
	if coin == "USDT" || coin == "USDC" {
		return nil, errors.New("USDT and USDC cannot be switched off")
	}
	params := client.Params{
		"coin": coin,
	}
	switch collateralSwitch {
	case ON:
		params["collateralSwitch"] = "ON"
	case OFF:
		params["collateralSwitch"] = "OFF"

	}

	response, err := s.client.Post(Endpoints.Collateral, params)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status())
	}

	var resp CollateralInfoResponse
	err = response.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CollateralCoin) GetInfo(currency string) (*CollateralInfoResponse, error) {
	params := client.Params{}
	if currency != "" {
		params["currency"] = currency
	}

	response, err := s.client.Get("/v5/account/collateral-info", params)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status())
	}

	var resp CollateralInfoResponse
	err = response.Unmarshal(&resp)

	if resp.RetCode != 0 {
		return nil, fmt.Errorf("API error: %s", resp.RetMsg)
	}

	return &resp, nil
}

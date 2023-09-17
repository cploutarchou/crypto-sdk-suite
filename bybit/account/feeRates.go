package account

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"net/http"
)

const FeeRatesEndpoint = "/v5/account/fee-rate"

type FeeRates struct {
	client *client.Client
}

func NewFeeRates(client *client.Client) *FeeRates {
	return &FeeRates{client: client}
}

func (fr *FeeRates) GetFeeRate(category string, symbol, baseCoin string) (*FeeRatesResponse, error) {
	// Construct parameters
	params := client.Params{
		"category": category,
	}

	// Set optional parameters
	if symbol != "" {
		params["symbol"] = symbol
	}
	if baseCoin != "" {
		params["baseCoin"] = baseCoin
	}

	response, err := fr.client.Get(FeeRatesEndpoint, params)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status code: %d", response.StatusCode())
	}

	var feeRatesResponse FeeRatesResponse
	err = response.Unmarshal(&feeRatesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &feeRatesResponse, nil

}

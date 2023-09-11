package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type CoinGreeks struct {
	client *client.Client
}

func NewCoinGreeks(client *client.Client) *CoinGreeks {
	return &CoinGreeks{client: client}
}

func (cg *CoinGreeks) Get(coin string) (*CoinGreekRes, error) {
	params := client.Params{}

	// Only add baseCoin to parameters if provided.
	if coin != "" {
		params["baseCoin"] = coin
	}

	response, err := cg.client.Get(Endpoints.CoinGreek, params)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, response.Error()
	}
	var coinGreekRes CoinGreekRes
	err = response.Unmarshal(&coinGreekRes)

	return &coinGreekRes, nil
}

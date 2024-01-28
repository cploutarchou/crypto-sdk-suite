package binance

import (
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/requests"
)

type Generic struct {
	*requests.Client
}

func NewGeneric(client *requests.Client) *Generic {
	return &Generic{
		Client: client,
	}
}
func (g *Generic) Ping() error {
	var responseData struct{}
	return g.MakeRequest(http.MethodGet, constants.PingEndpoint, &responseData)
}

func (g *Generic) CheckServerTime() (int64, error) {
	var responseData ServerTimeResponse

	if err := g.MakeRequest(http.MethodGet, constants.ServerTimeEndpoint, &responseData); err != nil {
		return 0, err
	}

	return responseData.ServerTime, nil
}

func (g *Generic) GetExchangeInfo() (*ExchangeInfo, error) {
	var response ExchangeInfo
	if err := g.MakeRequest(http.MethodGet, constants.ExchangeInfoEndpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

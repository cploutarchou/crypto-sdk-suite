package futures

import (
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"
)

// Generic defines an interface for generic API operations.
type Generic interface {
	Ping() (interface{}, error)
	CheckServerTime() (int64, error)
	GetExchangeInfo() (*models.ExchangeInfo, error)
}

// genericServiceImpl implements GenericService using a Binance futures client.
type genericServiceImpl struct {
	*client.Client
}

// NewGeneric creates a new instance of GenericService.
func NewGeneric(client *client.Client) Generic {
	return &genericServiceImpl{
		client,
	}
}

// Ping checks the connectivity to the Binance API server.
func (g *genericServiceImpl) Ping() (interface{}, error) {
	var responseData struct{}
	return responseData, g.MakeRequest(http.MethodGet, constants.PingEndpoint, &responseData)
}

// CheckServerTime retrieves the server time from the Binance API.
func (g *genericServiceImpl) CheckServerTime() (int64, error) {
	var responseData models.ServerTimeResponse

	if err := g.MakeRequest(http.MethodGet, constants.ServerTimeEndpoint, &responseData); err != nil {
		return 0, err
	}

	return responseData.ServerTime, nil
}

// GetExchangeInfo fetches exchange information from the Binance API.
func (g *genericServiceImpl) GetExchangeInfo() (*models.ExchangeInfo, error) {
	var response models.ExchangeInfo
	if err := g.MakeRequest(http.MethodGet, constants.ExchangeInfoEndpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

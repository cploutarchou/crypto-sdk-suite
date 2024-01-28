package binance

import (
	"net/http"
)

func (b *Client) Ping() error {
	var responseData struct{}
	return b.makeRequest(http.MethodGet, pingEndpoint, &responseData)
}

func (b *Client) CheckServerTime() (int64, error) {
	var responseData ServerTimeResponse

	if err := b.makeRequest(http.MethodGet, serverTimeEndpoint, &responseData); err != nil {
		return 0, err
	}

	return responseData.ServerTime, nil
}

func (b *Client) GetExchangeInfo() (*ExchangeInfo, error) {
	var response ExchangeInfo
	if err := b.makeRequest(http.MethodGet, exchangeInfoEndpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

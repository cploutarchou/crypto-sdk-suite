package market

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/requests"
	"net/http"
)

// Client wraps the requests.Client for local method definitions.
type Client struct {
	*requests.Client
}

// NewBinanceClient creates a new Binance futures client.
func NewBinanceClient(apiKey, apiSecret string, isTestnet bool) *Client {
	client := requests.NewFuturesClient(apiKey, apiSecret, isTestnet)
	return &Client{client}
}

// GetOrderBook retrieves the order book for a specific symbol.
func (c *Client) GetOrderBook(symbol string, limit int) (*OrderBookResponse, error) {
	endpoint := fmt.Sprintf("/fapi/v1/depth?symbol=%s&limit=%d", symbol, limit)

	var response OrderBookResponse
	if err := c.MakeRequest(http.MethodGet, endpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRecentTrades retrieves the recent trades for a specific symbol.
func (c *Client) GetRecentTrades(symbol string, limit int) ([]Trade, error) {
	endpoint := fmt.Sprintf("/fapi/v1/trades?symbol=%s&limit=%d", symbol, limit)

	var trades []Trade
	if err := c.MakeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

// GetHistoricalTrades retrieves older market historical trades for a specific symbol.
func (c *Client) GetHistoricalTrades(symbol string, limit int, fromId int64) ([]Trade, error) {
	endpoint := fmt.Sprintf("/fapi/v1/historicalTrades?symbol=%s", symbol)

	// Adding limit and fromId to the endpoint if they are provided
	if limit > 0 {
		endpoint += fmt.Sprintf("&limit=%d", limit)
	}
	if fromId > 0 {
		endpoint += fmt.Sprintf("&fromId=%d", fromId)
	}

	var trades []Trade
	if err := c.MakeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

// GetAggregateTrades retrieves compressed, aggregate market trades for a specific symbol.
func (c *Client) GetAggregateTrades(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error) {
	endpoint := fmt.Sprintf("/fapi/v1/aggTrades?symbol=%s", symbol)

	// Add additional parameters if provided
	if fromId > 0 {
		endpoint += fmt.Sprintf("&fromId=%d", fromId)
	}
	if startTime > 0 {
		endpoint += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime > 0 {
		endpoint += fmt.Sprintf("&endTime=%d", endTime)
	}
	if limit > 0 {
		endpoint += fmt.Sprintf("&limit=%d", limit)
	}

	var aggTrades []AggregateTrade
	if err := c.MakeRequest(http.MethodGet, endpoint, &aggTrades); err != nil {
		return nil, err
	}

	return aggTrades, nil
}

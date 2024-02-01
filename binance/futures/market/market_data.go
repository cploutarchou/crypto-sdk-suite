package market

import (
	"fmt"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
)

// Market defines the interface for market operations.
type Market interface {
	// GetOrderBook retrieves the order book for a specific symbol.
	GetOrderBook(symbol string, limit int) (*OrderBookResponse, error)

	// GetRecentTrades retrieves the recent trades for a specific symbol.
	GetRecentTrades(symbol string, limit int) ([]Trade, error)

	// GetHistoricalTrades retrieves older market historical trades for a specific symbol.
	GetHistoricalTrades(symbol string, limit int, fromId int64) ([]Trade, error)

	// GetAggregateTrades retrieves compressed, aggregate market trades for a specific symbol.
	GetAggregateTrades(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error)
}

type marketImpl struct {
	*client.Client
}

// NewMarket creates a new Market instance.
func NewMarket(client *client.Client) Market {
	return &marketImpl{client}
}

// buildEndpoint creates a formatted API endpoint string.
func buildEndpoint(base string, symbol string, params ...interface{}) string {
	endpoint := fmt.Sprintf(base, symbol)
	for _, param := range params {
		endpoint += fmt.Sprintf("&%s", param)
	}
	return endpoint
}

// GetOrderBook retrieves the order book for a specific symbol.
func (m *marketImpl) GetOrderBook(symbol string, limit int) (*OrderBookResponse, error) {
	endpoint := buildEndpoint("/fapi/v1/depth?symbol=%s&limit=%d", symbol, fmt.Sprintf("limit=%d", limit))
	response := new(OrderBookResponse)
	if err := m.MakeRequest(http.MethodGet, endpoint, response); err != nil {
		return nil, fmt.Errorf("failed to get order book: %w", err)
	}
	return response, nil
}

// GetRecentTrades retrieves the recent trades for a specific symbol.
func (m *marketImpl) GetRecentTrades(symbol string, limit int) ([]Trade, error) {
	endpoint := buildEndpoint("/fapi/v1/trades?symbol=%s&limit=%d", symbol, fmt.Sprintf("limit=%d", limit))
	var trades []Trade
	if err := m.MakeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, fmt.Errorf("failed to get recent trades: %w", err)
	}
	return trades, nil
}

// GetHistoricalTrades retrieves older market historical trades for a specific symbol.
func (m *marketImpl) GetHistoricalTrades(symbol string, limit int, fromId int64) ([]Trade, error) {
	endpoint := buildEndpoint("/fapi/v1/historicalTrades?symbol=%s", symbol,
		fmt.Sprintf("limit=%d", limit),
		fmt.Sprintf("fromId=%d", fromId))
	var trades []Trade
	if err := m.MakeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, fmt.Errorf("failed to get historical trades: %w", err)
	}
	return trades, nil
}

// GetAggregateTrades retrieves compressed, aggregate market trades for a specific symbol.
func (m *marketImpl) GetAggregateTrades(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error) {
	endpoint := buildEndpoint("/fapi/v1/aggTrades?symbol=%s", symbol,
		fmt.Sprintf("fromId=%d", fromId),
		fmt.Sprintf("startTime=%d", startTime),
		fmt.Sprintf("endTime=%d", endTime),
		fmt.Sprintf("limit=%d", limit))
	var aggTrades []AggregateTrade
	if err := m.MakeRequest(http.MethodGet, endpoint, &aggTrades); err != nil {
		return nil, fmt.Errorf("failed to get aggregate trades: %w", err)
	}
	return aggTrades, nil
}

package market

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
)

// Market defines the interface for market operations.
type Market interface {
	Ping() (interface{}, error)
	CheckServerTime() (int64, error)
	GetExchangeInfo() (*models.ExchangeInfo, error)
	// OrderBook retrieves the order book for a specific symbol.
	OrderBook(symbol string, limit int) (*OrderBookResponse, error)

	// RecentTradesList retrieves the recent trades for a specific symbol.
	RecentTradesList(symbol string, limit int) ([]Trade, error)

	// OldTradesLookup retrieves older market historical trades for a specific symbol.
	OldTradesLookup(symbol string, limit int, fromId int64) ([]Trade, error)

	// CompressedAggregateTradesList retrieves compressed, aggregate market trades for a specific symbol.
	CompressedAggregateTradesList(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error)
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

// Ping checks the connectivity to the Binance API server.
func (g *marketImpl) Ping() (interface{}, error) {
	var responseData struct{}
	return responseData, g.MakeRequestWithoutSignature(http.MethodGet, constants.PingEndpoint, &responseData)
}

// CheckServerTime retrieves the server time from the Binance API.
func (g *marketImpl) CheckServerTime() (int64, error) {
	var responseData models.ServerTimeResponse

	if err := g.MakeRequestWithoutSignature(http.MethodGet, constants.ServerTimeEndpoint, &responseData); err != nil {
		return 0, err
	}

	return responseData.ServerTime, nil
}

// GetExchangeInfo fetches exchange information from the Binance API.
func (g *marketImpl) GetExchangeInfo() (*models.ExchangeInfo, error) {
	var response models.ExchangeInfo
	if err := g.MakeRequestWithoutSignature(http.MethodGet, constants.ExchangeInfoEndpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// OrderBook retrieves the order book for a specific symbol.
func (m *marketImpl) OrderBook(symbol string, limit int) (*OrderBookResponse, error) {
	endpoint := buildEndpoint("/fapi/v1/depth?symbol=%s", symbol, fmt.Sprintf("limit=%d", limit))
	response := new(OrderBookResponse)
	if err := m.MakeRequestWithoutSignature(http.MethodGet, endpoint, response); err != nil {
		return nil, fmt.Errorf("failed to get order book: %w", err)
	}
	return response, nil
}

// RecentTradesList retrieves the recent trades for a specific symbol.
func (m *marketImpl) RecentTradesList(symbol string, limit int) ([]Trade, error) {
	endpoint := buildEndpoint("/fapi/v1/trades?symbol=%s&limit=%d", symbol, fmt.Sprintf("limit=%d", limit))
	var trades []Trade
	if err := m.MakeRequestWithoutSignature(http.MethodGet, endpoint, &trades); err != nil {
		return nil, fmt.Errorf("failed to get recent trades: %w", err)
	}
	return trades, nil
}

// OldTradesLookup retrieves older market historical trades for a specific symbol.
func (m *marketImpl) OldTradesLookup(symbol string, limit int, fromId int64) ([]Trade, error) {
	var params []interface{}

	if limit != -1 {
		params = append(params, fmt.Sprintf("limit=%d", limit))
	}
	if fromId != -1 {
		params = append(params, fmt.Sprintf("fromId=%d", fromId))
	}

	endpoint := buildEndpoint("/fapi/v1/historicalTrades?symbol=%s", symbol, params...)
	var trades []Trade // Change this line

	if err := m.MakeAuthenticatedRequest(http.MethodGet, endpoint, "", &trades); err != nil {
		return nil, fmt.Errorf("failed to get historical trades: %w", err)
	}

	return trades, nil
}

// CompressedAggregateTradesList retrieves compressed, aggregate market trades for a specific symbol.
func (m *marketImpl) CompressedAggregateTradesList(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error) {
	var params []interface{}

	if fromId != -1 {
		params = append(params, fmt.Sprintf("fromId=%d", fromId))
	}
	if startTime != -1 {
		params = append(params, fmt.Sprintf("startTime=%d", startTime))
	}
	if endTime != -1 {
		params = append(params, fmt.Sprintf("endTime=%d", endTime))
	}
	if limit != -1 {
		params = append(params, fmt.Sprintf("limit=%d", limit))
	}

	endpoint := buildEndpoint("/fapi/v1/aggTrades?symbol=%s", symbol, params...)
	var aggTrades []AggregateTrade

	if err := m.MakeRequestWithoutSignature(http.MethodGet, endpoint, &aggTrades); err != nil {
		return nil, fmt.Errorf("failed to get aggregate trades: %w", err)
	}

	return aggTrades, nil
}

package market

import (
	"fmt"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"
)

// Market defines the interface for market operations.
type Market interface {
	Ping() (any, error)
	CheckServerTime() (int64, error)
	GetExchangeInfo() (*models.ExchangeInfo, error)
	OrderBook(symbol string, limit int) (*OrderBookResponse, error)
	RecentTradesList(symbol string, limit int) ([]Trade, error)
	OldTradesLookup(symbol string, limit int, fromId int64) ([]Trade, error)
	CompressedAggregateTradesList(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error)
	KlineCandlestickData(symbol string, interval Interval, startTime, endTime int64, limit int) ([][]any, error)
}

type marketImpl struct {
	*client.Client
}

// NewMarket creates a new Market instance.
func NewMarket(client *client.Client) Market {
	return &marketImpl{client}
}

// buildEndpoint creates a formatted API endpoint string.
func buildEndpoint(base string, symbol string, params ...any) string {
	endpoint := fmt.Sprintf(base, symbol)
	for _, param := range params {
		endpoint += fmt.Sprintf("&%s", param)
	}
	return endpoint
}

// Ping checks the connectivity to the Binance API server.
func (m *marketImpl) Ping() (any, error) {
	var responseData struct{}
	return responseData, m.MakeRequestWithoutSignature(http.MethodGet, constants.PingEndpoint, &responseData)
}

// CheckServerTime retrieves the server time from the Binance API.
func (m *marketImpl) CheckServerTime() (int64, error) {
	var responseData models.ServerTimeResponse

	if err := m.MakeRequestWithoutSignature(http.MethodGet, constants.ServerTimeEndpoint, &responseData); err != nil {
		return 0, err
	}

	return responseData.ServerTime, nil
}

// GetExchangeInfo fetches exchange information from the Binance API.
func (m *marketImpl) GetExchangeInfo() (*models.ExchangeInfo, error) {
	var response models.ExchangeInfo
	if err := m.MakeRequestWithoutSignature(http.MethodGet, constants.ExchangeInfoEndpoint, &response); err != nil {
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
	var params []any

	if limit != -1 {
		params = append(params, fmt.Sprintf("limit=%d", limit))
	}
	if fromId != -1 {
		params = append(params, fmt.Sprintf("fromId=%d", fromId))
	}

	endpoint := buildEndpoint("/fapi/v1/historicalTrades?symbol=%s", symbol, params...)
	var trades []Trade

	if err := m.MakeAuthenticatedRequest(http.MethodGet, endpoint, "", &trades); err != nil {
		return nil, fmt.Errorf("failed to get historical trades: %w", err)
	}

	return trades, nil
}

// CompressedAggregateTradesList retrieves compressed, aggregate market trades for a specific symbol.
func (m *marketImpl) CompressedAggregateTradesList(symbol string, fromID, startTime, endTime int64, limit int) ([]AggregateTrade, error) {
	var params []any

	if fromID != -1 {
		params = append(params, fmt.Sprintf("fromId=%d", fromID))
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

// KlineCandlestickData retrieves kline candlestick data for a specific symbol.
func (m *marketImpl) KlineCandlestickData(symbol string, interval Interval, startTime, endTime int64, limit int) ([][]any, error) {
	var params []any

	if interval != "" {
		params = append(params, fmt.Sprintf("interval=%s", interval))
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

	endpoint := buildEndpoint("/fapi/v1/klines?symbol=%s", symbol, params...)
	var klines [][]any

	if err := m.MakeRequestWithoutSignature(http.MethodGet, endpoint, &klines); err != nil {
		return nil, fmt.Errorf("failed to get kline candlestick data: %w", err)
	}

	return klines, nil
}

package binance

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	productionBaseURL = "https://fapi.binance.com"
	testnetBaseURL    = "https://testnet.binancefuture.com"
	productionWSURL   = "wss://fstream.binancefuture.com"
	testnetWSURL      = "wss://fstream.testnet.binancefuture.com"
)

// API endpoints
const (
	pingEndpoint         = "/fapi/v1/ping"
	serverTimeEndpoint   = "/fapi/v1/time"
	exchangeInfoEndpoint = "/fapi/v1/exchangeInfo"
	orderBookEndpoint    = "/fapi/v1/depth"
	recentTradesEndpoint = "/fapi/v1/trades"
)

type BinanceClient struct {
	APIKey    string
	APISecret string
	BaseURL   string
	WSBaseURL string
}

func NewBinanceClient(apiKey, apiSecret string, isTestnet bool) *BinanceClient {
	var baseURL, wsBaseURL string
	if isTestnet {
		baseURL, wsBaseURL = testnetBaseURL, testnetWSURL
	} else {
		baseURL, wsBaseURL = productionBaseURL, productionWSURL
	}

	return &BinanceClient{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
		WSBaseURL: wsBaseURL,
	}
}

func (b *BinanceClient) createSignature(data string) string {
	h := hmac.New(sha256.New, []byte(b.APISecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *BinanceClient) makeRequest(method, endpoint string, responseData interface{}) error {
	resp, err := b.makeAuthenticatedRequest(method, endpoint, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &responseData)
}

func (b *BinanceClient) makeAuthenticatedRequest(method, endpoint, bodyData string) (*http.Response, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	data := fmt.Sprintf("%s&timestamp=%d", bodyData, timestamp)
	signature := b.createSignature(data)

	reqURL := fmt.Sprintf("%s%s?%s&signature=%s", b.BaseURL, endpoint, data, signature)

	req, err := http.NewRequest(method, reqURL, bytes.NewBufferString(bodyData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", b.APIKey)
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

func (b *BinanceClient) Ping() error {
	var responseData struct{}
	return b.makeRequest(http.MethodGet, pingEndpoint, &responseData)
}

func (b *BinanceClient) CheckServerTime() (int64, error) {
	var response struct {
		ServerTime int64 `json:"serverTime"`
	}

	if err := b.makeRequest(http.MethodGet, serverTimeEndpoint, &response); err != nil {
		return 0, err
	}

	return response.ServerTime, nil
}

func (b *BinanceClient) GetExchangeInfo() (*ExchangeInfo, error) {
	var response ExchangeInfo
	if err := b.makeRequest(http.MethodGet, exchangeInfoEndpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetOrderBook retrieves the order book for a specific symbol.
func (b *BinanceClient) GetOrderBook(symbol string, limit int) (*OrderBookResponse, error) {
	endpoint := fmt.Sprintf("/fapi/v1/depth?symbol=%s&limit=%d", symbol, limit)

	var response OrderBookResponse
	if err := b.makeRequest(http.MethodGet, endpoint, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRecentTrades retrieves the recent trades for a specific symbol.
func (b *BinanceClient) GetRecentTrades(symbol string, limit int) ([]Trade, error) {
	endpoint := fmt.Sprintf("/fapi/v1/trades?symbol=%s&limit=%d", symbol, limit)

	var trades []Trade
	if err := b.makeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

// GetHistoricalTrades retrieves older market historical trades for a specific symbol.
func (b *BinanceClient) GetHistoricalTrades(symbol string, limit int, fromId int64) ([]Trade, error) {
	endpoint := fmt.Sprintf("/fapi/v1/historicalTrades?symbol=%s", symbol)

	// Adding limit and fromId to the endpoint if they are provided
	if limit > 0 {
		endpoint += fmt.Sprintf("&limit=%d", limit)
	}
	if fromId > 0 {
		endpoint += fmt.Sprintf("&fromId=%d", fromId)
	}

	var trades []Trade
	if err := b.makeRequest(http.MethodGet, endpoint, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

// GetAggregateTrades retrieves compressed, aggregate market trades for a specific symbol.
func (b *BinanceClient) GetAggregateTrades(symbol string, fromId, startTime, endTime int64, limit int) ([]AggregateTrade, error) {
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
	if err := b.makeRequest(http.MethodGet, endpoint, &aggTrades); err != nil {
		return nil, err
	}

	return aggTrades, nil
}

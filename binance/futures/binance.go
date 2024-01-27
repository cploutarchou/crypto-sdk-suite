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

// Endpoints
const (
	pingEndpoint         = "/fapi/v1/ping"
	serverTimeEndpoint   = "/fapi/v1/time"
	exchangeInfoEndpoint = "/fapi/v1/exchangeInfo"
)

// BinanceClient struct holds API credentials and URLs.
type BinanceClient struct {
	APIKey    string
	APISecret string
	BaseURL   string
	WSBaseURL string
}

// NewBinanceClient creates a new Binance client instance.
func NewBinanceClient(apiKey, apiSecret string, isTestnet bool) *BinanceClient {
	baseURL := productionBaseURL
	wsBaseURL := productionWSURL

	if isTestnet {
		baseURL = testnetBaseURL
		wsBaseURL = testnetWSURL
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

	client := &http.Client{}
	return client.Do(req)
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

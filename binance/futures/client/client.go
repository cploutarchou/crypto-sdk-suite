package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
)

// Config stores configuration for the API client.
type Config struct {
	APIKey    string
	APISecret string
	BaseURL   string
	WSBaseURL string
}

// Client represents a client for Binance's futures trading.
type Client struct {
	sync.Mutex
	config Config
}

// NewFuturesClient creates a new client instance.
func NewFuturesClient(apiKey, apiSecret string, isTestnet bool) *Client {
	var baseURL, wsBaseURL string
	if isTestnet {
		baseURL, wsBaseURL = constants.TestnetBaseURL, constants.TestnetWSURL
	} else {
		baseURL, wsBaseURL = constants.ProductionBaseURL, constants.ProductionWSURL
	}

	config := Config{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
		WSBaseURL: wsBaseURL,
	}

	return &Client{
		config: config,
	}
}

// MakeRequest handles making an API request.
func (c *Client) MakeRequest(method, endpoint string, responseData interface{}) error {
	c.Lock()
	defer c.Unlock()

	resp, err := c.MakeAuthenticatedRequest(method, endpoint, "")
	if err != nil {
		log.Printf("Error making authenticated request: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	return json.Unmarshal(body, &responseData)
}

// MakeAuthenticatedRequest creates an authenticated request to the API.
func (c *Client) MakeAuthenticatedRequest(method, endpoint, bodyData string) (*http.Response, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	data := fmt.Sprintf("%s&timestamp=%d", bodyData, timestamp)
	signature := c.createSignature(data)

	reqURL := fmt.Sprintf("%s%s?%s&signature=%s", c.config.BaseURL, endpoint, data, signature)

	req, err := http.NewRequest(method, reqURL, bytes.NewBufferString(bodyData))
	if err != nil {
		log.Printf("Error creating new request: %v", err)
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

// createSignature generates the HMAC SHA256 signature for the request.
func (c *Client) createSignature(data string) string {
	h := hmac.New(sha256.New, []byte(c.config.APISecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

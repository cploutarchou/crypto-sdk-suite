package requests

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/constants"
	"io"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex
	APIKey    string
	APISecret string
	BaseURL   string
	WSBaseURL string
}

func NewFuturesClient(apiKey, apiSecret string, isTestnet bool) *Client {
	var baseURL, wsBaseURL string
	if isTestnet {
		baseURL, wsBaseURL = constants.TestnetBaseURL, constants.TestnetWSURL
	} else {
		baseURL, wsBaseURL = constants.ProductionBaseURL, constants.ProductionWSURL
	}

	return &Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
		WSBaseURL: wsBaseURL,
	}
}

func (b *Client) MakeRequest(method, endpoint string, responseData interface{}) error {
	b.Lock()
	defer b.Unlock()

	resp, err := b.MakeAuthenticatedRequest(method, endpoint, "")
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

func (b *Client) MakeAuthenticatedRequest(method, endpoint, bodyData string) (*http.Response, error) {
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

func (b *Client) createSignature(data string) string {
	h := hmac.New(sha256.New, []byte(b.APISecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

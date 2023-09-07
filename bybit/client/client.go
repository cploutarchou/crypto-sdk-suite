package client

import "C"
import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"
)

const (
	POST Method = "POST"
	GET  Method = "GET"
)

type Requester interface {
	Get(path string, params Params) (Response, error)
	Post(path string, params Params) (Response, error)
}

type Client struct {
	key        string
	secretKey  string
	httpClient *http.Client
	IsTestNet  bool
	lock       sync.RWMutex // Might be used for thread safety in the future
}

type Method string
type Params map[string]string

type Request struct {
	method Method
	path   string
	params Params
}

func NewClient(key, secretKey string, isTestnet bool) *Client {
	return &Client{
		key:        key,
		secretKey:  secretKey,
		httpClient: &http.Client{},
		IsTestNet:  isTestnet,
	}
}

func (c *Client) Get(path string, params Params) (Response, error) {
	return c.doRequest(GET, path, params)
}

func (c *Client) Post(path string, params Params) (Response, error) {
	return c.doRequest(POST, path, params)
}

func (c *Client) doRequest(method Method, path string, params Params) (Response, error) {
	req := &Request{
		method: method,
		path:   path,
		params: params,
	}
	return c.do(req)
}

func (c *Client) do(req *Request) (Response, error) {
	baseURL := BaseURL
	if c.IsTestNet {
		baseURL = TestnetBaseURL
	}

	var (
		httpReq *http.Request
		err     error
	)

	switch req.method {
	case GET:
		httpReq, err = c.newGETRequest(baseURL, req)
	case POST:
		httpReq, err = c.newPOSTRequest(baseURL, req)
	default:
		err = errors.New("unsupported method")
	}

	if err != nil {
		return nil, err
	}

	c.setCommonHeaders(httpReq, req.params)

	resp, err := c.httpClient.Do(httpReq)

	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	return NewResponse(resp), nil
}

func (c *Client) newGETRequest(baseURL string, req *Request) (*http.Request, error) {
	queryParams := url.Values{}
	for k, v := range req.params {
		queryParams.Set(k, v)
	}

	return http.NewRequest(string(GET), baseURL+req.path+"?"+queryParams.Encode(), nil)
}

func (c *Client) newPOSTRequest(baseURL string, req *Request) (*http.Request, error) {
	jsonData, err := json.Marshal(req.params)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(string(POST), baseURL+req.path, bytes.NewBuffer(jsonData))
}

func (c *Client) setCommonHeaders(req *http.Request, params Params) {
	timestamp := fmt.Sprintf("%d", time.Now().UTC().UnixMilli())

	req.Header.Add("X-BAPI-SIGN-TYPE", "2")
	req.Header.Add("X-BAPI-SIGN", GenerateSignature(c.secretKey, c.key, params, timestamp))
	req.Header.Add("X-BAPI-API-KEY", c.key)
	req.Header.Add("X-BAPI-TIMESTAMP", timestamp)
	req.Header.Add("X-BAPI-RECV-WINDOW", "50000")
	req.Header.Add("Content-Type", "application/json")
}

func GenerateSignature(secretKey, apiKey string, params Params, timestamp string) string {
	// Start with the timestamp
	str := timestamp

	// Append the API Key
	str += apiKey

	// Append recv window value
	str += "50000" // adjust as needed

	// Sort keys and append sorted query parameters
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	queryParams := url.Values{}
	for _, k := range keys {
		queryParams.Set(k, params[k])
	}
	str += fmt.Sprintf("%s", queryParams.Encode())
	// Generate HMAC-SHA256 signature
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(str))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const (
	recvWindow = "5000"

	// BaseURL is the base URL for the Bybit API
	BaseURL = "https://api.bybit.com"

	// TestnetBaseURL is the base URL for the Bybit Testnet API
	TestnetBaseURL = "https://api-testnet.bybit.com"

	// ApiVersion is the version of the Bybit API
	ApiVersion        = "v5"
	GET        Method = "GET"
	POST       Method = "POST"
	// Globals
	timestampKey  = "X-BAPI-TIMESTAMP"
	signatureKey  = "X-BAPI-SIGN"
	apiRequestKey = "X-BAPI-API-KEY"
	recvWindowKey = "X-BAPI-RECV-WINDOW"
	signTypeKey   = "X-BAPI-SIGN-TYPE"
)

type Requester interface {
	Get(path string, params Params) (Response, error)
	Post(path string, params Params) (Response, error)
}

type Client struct {
	key             string
	secretKey       string
	httpClient      *http.Client
	IsTestNet       bool
	params          []byte
	QueryParams     url.Values
	endpointLimiter *EndpointRateLimiter
}

type Method string
type Params map[string]interface{}

type Request struct {
	method Method
	path   string
	params Params
}

// Initializes endpoint limiters based on predefined rates and sets a sensible burst size.
func (c *Client) initializeEndpointLimiters() {
	wg := sync.WaitGroup{}

	for endpoint, limit := range endpointLimits {
		wg.Add(1)
		go func(endpoint string, limit rate.Limit) {
			defer wg.Done()
			limiter := rate.NewLimiter(limit, 5)
			c.endpointLimiter.SetLimiter(endpoint, limiter)
		}(endpoint, limit)
	}

	wg.Wait()
}

// NewClient function to create a new client instance.
func NewClient(key, secretKey string, isTestnet bool) *Client {
	return &Client{
		key:             key,
		secretKey:       secretKey,
		httpClient:      &http.Client{},
		IsTestNet:       isTestnet,
		endpointLimiter: NewEndpointRateLimiter(),
	}
}
func (c *Client) Get(path string, params Params) (Response, error) {
	return c.doRequest(GET, path, params)
}

func (c *Client) Post(path string, params Params) (Response, error) {
	return c.doRequest(POST, path, params)
}

func (c *Client) doRequest(method Method, path string, params Params) (Response, error) {
	// Construct the endpoint key using the method and path
	endpointKey := fmt.Sprintf("%s %s", method, path)
	// Retrieve the existing limiter for the endpoint
	limiter := c.endpointLimiter.GetLimiter(endpointKey)
	if limiter == nil {
		log.Printf("Warning: No rate limiter found for %s. Requests for this endpoint may not be rate-limited.", endpointKey)
		// You might choose to handle this situation differently, such as by setting a default limiter.
	} else {
		// Wait for permission to proceed from the rate limiter
		ctx := context.Background() // Consider passing a context from higher-level methods
		if err := limiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limiter error: %w", err)
		}
	}

	// Create and execute the HTTP request as before
	req := &Request{
		method: method,
		path:   path,
		params: params,
	}
	return c.do(req)
}

func (c *Client) do(req *Request) (Response, error) {
	c.QueryParams = make(url.Values)
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

	c.setCommonHeaders(httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	c.params = nil
	return NewResponse(resp), nil
}

func (c *Client) newGETRequest(baseURL string, req *Request) (*http.Request, error) {
	c.QueryParams = url.Values{}
	for k, v := range req.params {
		c.QueryParams.Set(k, fmt.Sprintf("%v", v))
	}

	return http.NewRequest(string(GET), baseURL+req.path+"?"+c.QueryParams.Encode(), nil)
}

func (c *Client) newPOSTRequest(baseURL string, req *Request) (*http.Request, error) {
	jsonData, err := json.Marshal(req.params)
	if err != nil {
		return nil, err
	}
	c.params = nil
	c.params = jsonData
	return http.NewRequest(string(POST), baseURL+req.path, bytes.NewBuffer(jsonData))
}

func (c *Client) setCommonHeaders(req *http.Request) {

	timestamp := strconv.FormatInt(GetCurrentTime(), 10)
	req.Header.Set(signTypeKey, "2")
	req.Header.Set(apiRequestKey, c.key)
	req.Header.Set(timestampKey, strconv.FormatInt(GetCurrentTime(), 10))
	req.Header.Set(recvWindowKey, recvWindow)
	var signatureBase []byte
	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/json")
		signatureBase = []byte(timestamp + c.key + recvWindow + string(c.params[:]))
	} else {
		queryString := c.QueryParams.Encode()
		signatureBase = []byte(timestamp + c.key + recvWindow + queryString)
	}
	hmac256 := hmac.New(sha256.New, []byte(c.secretKey))
	hmac256.Write(signatureBase)
	signature := hex.EncodeToString(hmac256.Sum(nil))
	req.Header.Set(signatureKey, signature)

}

func GetCurrentTime() int64 {
	now := time.Now()
	unixNano := now.UnixNano()
	timeStamp := unixNano / int64(time.Millisecond)
	return timeStamp
}

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
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	recvWindow            = "5000"
	BaseURL               = "https://api.bybit.com"
	TestnetBaseURL        = "https://api-testnet.bybit.com"
	ApiVersion            = "v5"
	GET            Method = "GET"
	POST           Method = "POST"

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
	endpointKey := fmt.Sprintf("%s %s", method, path)
	limiter := c.endpointLimiter.GetLimiter(endpointKey)
	if limiter == nil {
		log.Printf("Warning: No rate limiter found for %s. Requests for this endpoint may not be rate-limited.", endpointKey)

		limiter = rate.NewLimiter(rate.Inf, 5)
		log.Println("Warning: Using an unlimited rate limiter for this endpoint.")
	}

	ctx := context.Background()
	if err := limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

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
	req.Header.Set(timestampKey, timestamp)
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

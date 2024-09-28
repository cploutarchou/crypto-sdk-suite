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
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

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

// Requester interface defines methods for making HTTP GET and POST requests
type Requester interface {
	Get(path string, params Params) (Response, error)
	Post(path string, params Params) (Response, error)
}

// Client struct holds information needed for API interaction
type Client struct {
	key             string
	secretKey       string
	httpClient      *http.Client
	IsTestNet       bool
	params          []byte
	QueryParams     url.Values
	endpointLimiter *EndpointRateLimiter
}

// Define HTTP method types as strings
type Method string

// Params represents parameters for the API request
type Params map[string]interface{}

// Request struct represents an HTTP request with method, path, and params
type Request struct {
	method Method
	path   string
	params Params
}

func (c *Client) initializeEndpointLimiters() {
	// Make sure the limiter map isn't nil
	if c.endpointLimiter == nil {
		c.endpointLimiter = NewEndpointRateLimiter()
	}

	// Set the limiters for each endpoint
	for endpoint, limit := range endpointLimits {
		// log.Printf("Setting rate limiter for endpoint: %s with limit %v requests/sec", endpoint, limit)
		limiter := rate.NewLimiter(limit, 5) // Burst size 5
		c.endpointLimiter.SetLimiter(endpoint, limiter)
	}
}

// NewClient creates a new client instance with API key, secret key, and testnet setting
func NewClient(key, secretKey string, isTestnet bool) *Client {
	client := &Client{
		key:             key,
		secretKey:       secretKey,
		httpClient:      &http.Client{},
		IsTestNet:       isTestnet,
		endpointLimiter: NewEndpointRateLimiter(),
	}

	// Initialize the rate limiters for all endpoints
	client.initializeEndpointLimiters()
	// fmt.Printf("Rate limiter initialized for endpoint: %+v", client.endpointLimiter)
	return client
}

// Get method performs a GET request to the specified API path with params
func (c *Client) Get(path string, params Params) (Response, error) {
	return c.doRequest(GET, path, params)
}

// Post method performs a POST request to the specified API path with params
func (c *Client) Post(path string, params Params) (Response, error) {
	return c.doRequest(POST, path, params)
}

// doRequest handles both GET and POST requests, applying rate limiting and signing
func (c *Client) doRequest(method Method, path string, params Params) (Response, error) {
	// Ensure the endpointLimiter is initialized
	if c.endpointLimiter == nil {
		return nil, fmt.Errorf("endpointLimiter is not initialized")
	}

	// Generate the endpoint key
	endpointKey := fmt.Sprintf("%s %s", method, path)

	// Get the rate limiter for this endpoint
	limiter := c.endpointLimiter.GetLimiter(endpointKey)
	if limiter == nil {
		// log.Printf("No rate limiter found for endpoint: %s. Applying default rate limit.", endpointKey)
		limiter = rate.NewLimiter(rate.Limit(30.0/60.0), 1) // Default to 30 requests per minute
	}

	// Wait for the rate limiter to allow the request
	ctx := context.Background()
	if err := limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	// Continue with request processing
	req := &Request{
		method: method,
		path:   path,
		params: params,
	}
	return c.do(req)
}

// do handles the actual execution of the HTTP request
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

	// Prepare the GET or POST request based on the method
	switch req.method {
	case GET:
		httpReq, err = c.newGETRequest(baseURL, req)
	case POST:
		httpReq, err = c.newPOSTRequest(baseURL, req)
	default:
		return nil, errors.New("unsupported method")
	}

	if err != nil {
		return nil, err
	}

	// Set common headers for the request
	c.setCommonHeaders(httpReq)

	// Execute the request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Process and return the response
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
	c.params = jsonData
	return http.NewRequest(string(POST), baseURL+req.path, bytes.NewBuffer(jsonData))
}
func (c *Client) setCommonHeaders(req *http.Request) {
	timestamp := strconv.FormatInt(GetCurrentTime(), 10) // Get the current timestamp in milliseconds
	req.Header.Set(signTypeKey, "2")
	req.Header.Set(apiRequestKey, c.key)
	req.Header.Set(timestampKey, timestamp)
	req.Header.Set(recvWindowKey, "5000") // Match Bybit's recvWindow of 5000 ms

	var signatureBase []byte
	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/json")
		// Concatenate timestamp, API key, recvWindow, and the request body for POST requests
		signatureBase = []byte(timestamp + c.key + "5000" + string(c.params[:]))
	} else {
		// Alphabetically sort query parameters and concatenate them with other fields for GET requests
		queryString := c.QueryParams.Encode() // Automatically sorts the parameters alphabetically
		signatureBase = []byte(timestamp + c.key + "5000" + queryString)
	}

	// Generate the HMAC-SHA256 signature
	hmac256 := hmac.New(sha256.New, []byte(c.secretKey))
	hmac256.Write(signatureBase)
	signature := hex.EncodeToString(hmac256.Sum(nil))

	// Set the signature in the headers
	req.Header.Set(signatureKey, signature)

	// Debug logging for troubleshooting
	log.Printf("Signature Base String: %s", string(signatureBase))
	log.Printf("Generated Signature: %s", signature)
	log.Printf("Headers: X-BAPI-API-KEY=%s, X-BAPI-TIMESTAMP=%s, X-BAPI-SIGN=%s", c.key, timestamp, signature)
}

func GetCurrentTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

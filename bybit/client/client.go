package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	key         string
	secretKey   string
	httpClient  *http.Client
	IsTestNet   bool
	params      []byte
	QueryParams url.Values
	rateLimiter *rate.Limiter
}

type Method string
type Params map[string]interface{}

type Request struct {
	method Method
	path   string
	params Params
}

func NewClient(key, secretKey string, isTestnet bool) *Client {
	limiter := rate.NewLimiter(rate.Every(time.Second/120), 120) // 120 requests per second
	return &Client{
		key:         key,
		secretKey:   secretKey,
		httpClient:  &http.Client{},
		IsTestNet:   isTestnet,
		rateLimiter: limiter,
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
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

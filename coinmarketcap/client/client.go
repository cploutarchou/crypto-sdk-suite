package client

import "C"
import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sync"
)

const (
	POST           Method = "POST"
	GET            Method = "GET"
	BaseURL               = "https://pro-api.coinmarketcap.com"
	TestnetBaseURL        = "https://sandbox-api.coinmarketcap.com"
)

type Requester interface {
	Get(path string, params Params) (Response, error)
	Post(path string, params Params) (Response, error)
}

type Client struct {
	apiKey     string
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

func NewClient(key string, isTestnet bool, httpClient ...*http.Client) *Client {
	client := &http.Client{}
	if len(httpClient) > 0 && httpClient[0] != nil {
		client = httpClient[0]
	}
	return &Client{
		apiKey:     key,
		httpClient: client,
		IsTestNet:  isTestnet,
	}
}

func (c *Client) SetIsTestNet(isTestNet bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.IsTestNet = isTestNet
}

func (c *Client) getBaseURL() string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.IsTestNet {
		return TestnetBaseURL
	}
	return BaseURL
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
	baseURL := c.getBaseURL()
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
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			logrus.Error(closeErr)
		}
	}()

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

func (c *Client) setCommonHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.apiKey)
	req.Header.Add("Accept", "application/json")
}

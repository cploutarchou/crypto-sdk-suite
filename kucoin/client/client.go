package client

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
	"strings"
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

	lock     sync.RWMutex // Might be used for thread safety in the future
	isFuture bool
}

type Method string
type Params map[string]interface{}

type Request struct {
	method Method
	path   string
	params Params
}

func NewClient(key, secretKey string) *Client {
	return &Client{
		key:        key,
		secretKey:  secretKey,
		httpClient: &http.Client{},
	}
}

func (c *Client) Get(path string, marketType MarketType, params Params) (Response, error) {
	c.isFuture = true
	if marketType == SpotMargin {
		c.isFuture = false
	}

	return c.doRequest(GET, path, params)
}

func (c *Client) Post(path string, marketType MarketType, params Params) (Response, error) {
	c.isFuture = true
	if marketType == SpotMargin {
		c.isFuture = false
	}
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
	baseURL := BaseURLFuture
	if !c.isFuture {
		baseURL = BaseURLSpotMargin
	}

	baseURL = baseURL + ApiVersion + "/"
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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed request: " + string(body))
	}

	return NewResponse(body), nil
}

func (c *Client) newGETRequest(baseURL string, req *Request) (*http.Request, error) {
	queryParams := url.Values{}
	for k, v := range req.params {
		queryParams.Add(k, v.(string))
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-SIGN-TYPE", "2")
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", fmt.Sprintf("%d", time.Now().UnixNano()/1000000))
	req.Header.Set("X-BAPI-SIGN", GenerateSignature(c.secretKey, params))
	req.Header.Set("X-BAPI-RECV-WINDOW", "50000")
}

func GenerateSignature(secretKey string, params Params) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	// Sort keys in alphabetical order
	sort.Strings(keys)

	var paramString string
	for _, k := range keys {
		paramString += k + "=" + params[k].(string) + "&"
	}

	// Trim the last '&'
	paramString = strings.TrimSuffix(paramString, "&")

	// Step 2: Use HMAC-SHA256 to hash the concatenated string, using secretKey as the key.
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(paramString))

	// Step 3: Convert the hash into a hexadecimal representation.
	signature := hex.EncodeToString(mac.Sum(nil))
	return signature
}

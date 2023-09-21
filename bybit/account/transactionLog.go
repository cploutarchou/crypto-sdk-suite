package account

import (
	"errors"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"net/http"
	"net/url"
)

// TransactionLog holds a client instance
type TransactionLog struct {
	client *client.Client
}

// NewTransactionLog initializes a new TransactionLog object with a client instance.
func NewTransactionLog(client *client.Client) *TransactionLog {
	if client == nil {
		panic("client should not be nil")
	}
	return &TransactionLog{
		client: client,
	}
}

// LogEntry represents a single log entry returned by the API
type LogEntry struct {
	ID              string `json:"id"`
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Type            string `json:"type"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeID         string `json:"tradeId"`
	OrderID         string `json:"orderId"`
	OrderLinkID     string `json:"orderLinkId"`
}

// LogResponse represents the response from the /v5/account/transaction-log endpoint
type LogResponse struct {
	List           []LogEntry `json:"list"`
	NextPageCursor string     `json:"nextPageCursor"`
}

// Get sends a GET request to the /v5/account/transaction-log endpoint to retrieve transaction logs.
func (tl *TransactionLog) Get(params map[string]string) (*LogResponse, error) {
	endpoint := "/v5/account/transaction-log"

	// Add the optional query parameters if provided
	if params != nil && len(params) > 0 {
		queryParams := url.Values{}
		for key, value := range params {
			queryParams.Add(key, value)
		}
		endpoint = endpoint + "?" + queryParams.Encode()
	}

	resp, err := tl.client.Get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to get transaction logs: non-200 status code received")
	}

	var logResponse LogResponse
	err = resp.Unmarshal(&logResponse)
	if err != nil {
		return nil, err
	}

	return &logResponse, nil
}

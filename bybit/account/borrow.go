package account

import (
	"errors"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

const borrowEndpoint = "/v5/account/borrow-history"

type BorrowRes struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		NextPageCursor string `json:"nextPageCursor"`
		List           []struct {
			CreatedTime               int64  `json:"createdTime"`
			CostExemption             string `json:"costExemption"`
			InterestBearingBorrowSize string `json:"InterestBearingBorrowSize"`
			Currency                  string `json:"currency"`
			HourlyBorrowRate          string `json:"hourlyBorrowRate"`
			BorrowCost                string `json:"borrowCost"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}
type Borrow struct {
	client *client.Client
}

func (b *Borrow) GetHistory(currency string, startTime, endTime, limit int, cursor string) (*BorrowRes, error) {

	params := client.Params{}

	if currency != "" {
		params["currency"] = currency
	}
	if startTime > 0 {
		params["startTime"] = fmt.Sprintf("%d", startTime)
	}
	if endTime > 0 {
		params["endTime"] = fmt.Sprintf("%d", endTime)
	}
	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if cursor != "" {
		params["cursor"] = cursor
	}

	response, err := b.client.Get(borrowEndpoint, params)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, errors.New("received non-200 response")
	}
	var borrowRes BorrowRes
	err = response.Unmarshal(&borrowRes)
	if err != nil {
		return nil, err
	}

	return &borrowRes, nil
}
func NewBorrow(client *client.Client) *Borrow {
	if client == nil {
		panic("client should not be nil")
	}
	return &Borrow{
		client: client,
	}
}

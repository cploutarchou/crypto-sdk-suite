package account

import (
	"errors"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

const (
	CollateralSwitchOn  = "ON"
	CollateralSwitchOff = "OFF"

	collateralEndpoint = "/v5/account/set-collateral-switch"
)

type CollateralCoin struct {
	client *client.Client
}

func NewSetCollateralCoin(c *client.Client) *CollateralCoin {
	return &CollateralCoin{client: c}
}

type CollateralSwitchResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

func (s *CollateralCoin) SwitchCollateral(coin string, collateralSwitch string) (*CollateralSwitchResponse, error) {
	if coin == "USDT" || coin == "USDC" {
		return nil, errors.New("USDT and USDC cannot be switched off")
	}

	params := client.Params{
		"coin":             coin,
		"collateralSwitch": collateralSwitch,
	}

	response, err := s.client.Post(collateralEndpoint, params)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status())
	}

	var resp CollateralSwitchResponse
	err = response.Unmarshal(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CollateralInfoResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     CollateralResult       `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type CollateralResult struct {
	List []CollateralData `json:"list"`
}

type CollateralData struct {
	CollateralSwitch    bool   `json:"collateralSwitch"`
	BorrowAmount        string `json:"borrowAmount"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	FreeBorrowingAmount string `json:"freeBorrowingAmount"`
	Borrowable          bool   `json:"borrowable"`
	Currency            string `json:"currency"`
	MaxBorrowingAmount  string `json:"maxBorrowingAmount"`
	HourlyBorrowRate    string `json:"hourlyBorrowRate"`
	BorrowUsageRate     string `json:"borrowUsageRate"`
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralRatio     string `json:"collateralRatio"`
}

func (s *CollateralCoin) GetCollateralInfo(currency string) (*CollateralInfoResponse, error) {
	params := client.Params{}
	if currency != "" {
		params["currency"] = currency
	}

	response, err := s.client.Get("/v5/account/collateral-info", params)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status())
	}

	var resp CollateralInfoResponse
	err = response.Unmarshal(&resp)

	if resp.RetCode != 0 {
		return nil, fmt.Errorf("API error: %s", resp.RetMsg)
	}

	return &resp, nil
}

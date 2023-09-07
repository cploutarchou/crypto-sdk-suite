package account

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Wallet struct {
	client *client.Client
}

func NewWallet(client *client.Client) *Wallet {
	if client == nil {
		panic("client should not be nil")
	}
	return &Wallet{
		client: client,
	}
}

const (
	endpoint string = "/v5/account/wallet-balance"
)

func (w Wallet) GetUnifiedWalletBalance(coins ...string) (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Unified)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}
func (w Wallet) GetAllUnifiedWalletBalance() (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Unified)

	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

func (w Wallet) GetAllSpotWalletBalance() (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Spot)

	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

func (w Wallet) GetSpotWalletBalance(coins ...string) (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Spot)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

func (w Wallet) GetAllContractWalletBalance() (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Contract)

	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

func (w Wallet) GetContractWalletBalance(coins ...string) (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = fmt.Sprintf("%s", Contract)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(endpoint, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

package account

import (
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Wallet struct {
	client *client.Client
}

func NewWallet(client_ *client.Client) *Wallet {
	if client_ == nil {
		panic("client should not be nil")
	}
	return &Wallet{
		client: client_,
	}
}

func (w Wallet) GetUnifiedWalletBalance(coins ...string) (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = string(Unified)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(Endpoints.Wallet, params)
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
	params["accountType"] = string(Unified)

	resp, err := w.client.Get(Endpoints.Wallet, params)
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
	params["accountType"] = string(Spot)

	resp, err := w.client.Get(Endpoints.Wallet, params)
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
	params["accountType"] = string(Spot)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(Endpoints.Wallet, params)
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
	params["accountType"] = string(Contract)

	resp, err := w.client.Get(Endpoints.Wallet, params)
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
	params["accountType"] = string(Contract)
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	if coinStr != "" {
		coinStr = coinStr[:len(coinStr)-1]
		params["coin"] = coinStr
	}
	resp, err := w.client.Get(Endpoints.Wallet, params)
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

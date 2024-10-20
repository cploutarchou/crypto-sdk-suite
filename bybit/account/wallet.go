package account

import (
	"fmt"
	"strings"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Wallet struct {
	*client.Client
}

func NewWallet(client_ *client.Client) *Wallet {
	if client_ == nil {
		panic("client should not be nil")
	}
	return &Wallet{
		client_,
	}
}

func (w Wallet) GetUnifiedWalletBalance(coins ...string) (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = string(Unified)

	// Construct the coin parameter string
	if len(coins) > 0 {
		params["coin"] = joinCoins(coins)
	}

	// Make the GET request
	resp, err := w.Get(Endpoints.Wallet, params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallet balance: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	// Unmarshal response
	var balanceResp WalletBalance
	if err := resp.Unmarshal(&balanceResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &balanceResp, nil
}

func joinCoins(coins []string) string {
	coinStr := ""
	for _, coin := range coins {
		coinStr += coin + ","
	}
	return strings.TrimRight(coinStr, ",") // Remove trailing comma
}

func (w Wallet) GetAllUnifiedWalletBalance() (*WalletBalance, error) {
	params := client.Params{}
	params["accountType"] = string(Unified)
	resp, err := w.Get(Endpoints.Wallet, params)
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

	resp, err := w.Get(Endpoints.Wallet, params)
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
	resp, err := w.Get(Endpoints.Wallet, params)
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

	resp, err := w.Get(Endpoints.Wallet, params)
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
	resp, err := w.Get(Endpoints.Wallet, params)
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

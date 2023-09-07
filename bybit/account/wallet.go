package account

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Wallet struct {
	client *client.Client
}

func NewWallet(client *client.Client) (Wallet, error) {
	if client == nil {
		return Wallet{}, fmt.Errorf("client is nil")
	}
	return Wallet{
		client: client,
	}, nil
}

const (
	testNetEndpoint string = "https://api-testnet.bybit.com/v5/account/wallet-balance"
	endpoint        string = "https://api.bybit.com/v2/private/wallet/balance"
)

func (w Wallet) GetWalletBalance(accountType AccountType, coins ...string) (*WalletBalance, error) {
	var endpoint_ string
	if w.client.IsTestNet {
		endpoint_ = testNetEndpoint
	} else {
		endpoint_ = endpoint
	}
	url := fmt.Sprintf("%s?accountType=%s", endpoint_, accountType)

	if len(coins) > 0 {
		coinParam := fmt.Sprintf("&coin=%s", coins[0])
		for _, coin := range coins[1:] {
			coinParam = fmt.Sprintf("%s,%s", coinParam, coin)
		}
		url += coinParam
	}

	resp, err := w.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status should be 200, got: %d, body: %s", resp.StatusCode(), resp.Status())
	}

	var balanceResp WalletBalance
	err = resp.Unmarshal(&balanceResp)
	if err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

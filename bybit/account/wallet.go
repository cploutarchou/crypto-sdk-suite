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

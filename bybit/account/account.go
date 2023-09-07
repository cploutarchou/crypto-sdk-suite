package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type Account interface {
	Wallet() (Wallet, error)
}

type account struct {
	client *client.Client
}

func (a *account) Wallet() (Wallet, error) {
	return NewWallet(a.client)
}
func New(client2 *client.Client) Account {
	return &account{client: client2}

}

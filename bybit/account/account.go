package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type Account interface {
	Wallet() *Wallet
	UpgradeToUnified() *UpgradeToUnified
}

type account struct {
	client *client.Client
}

func (a *account) Wallet() *Wallet {
	return NewWallet(a.client)
}
func (a *account) UpgradeToUnified() *UpgradeToUnified {
	return NewUpgradeToUnifiedRequest(a.client)
}

func New(client *client.Client) Account {
	return &account{client: client}
}

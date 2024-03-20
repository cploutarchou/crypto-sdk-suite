package wallet

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Wallet struct {
	*client.Client
}

func New(cli *client.Client) Wallet {
	return Wallet{cli}
}

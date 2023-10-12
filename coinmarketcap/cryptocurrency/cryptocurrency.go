package cryptocurrency

import "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"

type Cryptocurrency interface {
}

type cryptocurrencyImpl struct {
	client *client.Client
}

func New(client *client.Client) Cryptocurrency {
	return &cryptocurrencyImpl{
		client:client,
	}
}

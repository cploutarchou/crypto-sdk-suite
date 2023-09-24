package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type Margin struct {
	client *client.Client
}

func NewMargin(client *client.Client) *Margin {
	return &Margin{client: client}
}

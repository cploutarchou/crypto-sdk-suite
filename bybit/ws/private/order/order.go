package order

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Order struct {
	*client.Client
}

func New(cli *client.Client) Order {
	return Order{cli}
}

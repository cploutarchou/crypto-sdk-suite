package orderbook

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type OrderBook struct {
	*client.Client
}

func New(cli *client.Client) OrderBook {
	return OrderBook{cli}
}

package lt_ticker

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type LtTicker struct {
	*client.Client
}

func New(cli *client.Client) LtTicker {
	return LtTicker{cli}
}

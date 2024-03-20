package trade

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Trade struct {
	*client.Client
}

func New(cli *client.Client) Trade {
	return Trade{cli}
}

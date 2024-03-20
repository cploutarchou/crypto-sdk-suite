package lt_kline

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type LtKline struct {
	*client.Client
}

func New(cli *client.Client) LtKline {
	return LtKline{cli}
}

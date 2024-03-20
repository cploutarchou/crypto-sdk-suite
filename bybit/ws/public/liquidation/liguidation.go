package liquidation

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Liquidation struct {
	*client.Client
}

func New(cli *client.Client) Liquidation {
	return Liquidation{cli}
}

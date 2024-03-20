package greek

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Greek struct {
	*client.Client
}

func New(cli *client.Client) Greek {
	return Greek{cli}
}

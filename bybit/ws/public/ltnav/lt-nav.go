package ltnav

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type LtNav struct {
	*client.Client
}

func New(cli *client.Client) LtNav {
	return LtNav{cli}
}

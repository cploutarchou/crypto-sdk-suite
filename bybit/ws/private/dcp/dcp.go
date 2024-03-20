package dcp

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Dcp struct {
	*client.Client
}

func New(cli *client.Client) Dcp {
	return Dcp{cli}
}

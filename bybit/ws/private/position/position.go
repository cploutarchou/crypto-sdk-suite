package position

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Position struct {
	*client.Client
}

func New(cli *client.Client) Position {
	return Position{cli}
}

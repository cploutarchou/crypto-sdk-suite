package position

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type Position interface {
}
type impl struct {
	client *client.Client
}

func New(c *client.Client) Position {
	return &impl{client: c}
}

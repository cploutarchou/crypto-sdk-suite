package cryptocurrency

import (
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	idmap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/map"
)

type Cryptocurrency interface {
	GetMapID(params *idmap.Params) ([]idmap.Data, error)
}

type cryptocurrencyImpl struct {
	client *client.Client
}

func New(client *client.Client) Cryptocurrency {
	return &cryptocurrencyImpl{
		client: client,
	}
}

func (c *cryptocurrencyImpl) GetMapID(params *idmap.Params) ([]idmap.Data, error) {
	idMap := idmap.New(c.client)
	return idMap.GetID(params)
}

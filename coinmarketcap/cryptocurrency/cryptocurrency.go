package cryptocurrency

import (
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	gainer "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/gainer-looser"
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/info"
	idmap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/map"
)

type Cryptocurrency interface {
	Map() *idmap.Map
	GainersAndLosers() *gainer.GainersAndLosers
	Info() *info.Metadata
}

type cryptocurrencyImpl struct {
	client *client.Client
}

func New(client *client.Client) Cryptocurrency {
	return &cryptocurrencyImpl{
		client: client,
	}
}

func (c *cryptocurrencyImpl) Map() *idmap.Map {
	return idmap.New(c.client)

}

func (c *cryptocurrencyImpl) GainersAndLosers() *gainer.GainersAndLosers {
	return gainer.New(c.client)

}

func (c *cryptocurrencyImpl) Info() *info.Metadata {
	return info.New(c.client)

}

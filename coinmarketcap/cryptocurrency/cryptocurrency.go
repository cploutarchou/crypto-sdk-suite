package cryptocurrency

import (
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	gainer_looser "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/gainer-looser"
	idmap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/map"
)

type Cryptocurrency interface {
	GetMapID(params *idmap.Params) ([]idmap.Data, error)
	FetchGainersLosers(params *idmap.Params) ([]gainer_looser.CryptocurrencyData, error)
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
func (c *cryptocurrencyImpl) FetchGainersLosers(params *gainer_looser.Params) ([]gainer_looser.CryptocurrencyData, error) {
	idMap := gainer_looser.New(&c.client)
	return idMap.FetchGainersLosers(params)
}

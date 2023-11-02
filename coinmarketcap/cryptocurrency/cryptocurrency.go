package cryptocurrency

import (
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	gainer "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/gainer-looser"
	idmap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/map"
)

type Cryptocurrency interface {
	GetMapID(params *idmap.Params) ([]idmap.Data, error)
	FetchGainersLosers(params *gainer.Params) ([]gainer.CryptocurrencyData, error)
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
func (c *cryptocurrencyImpl) FetchGainersLosers(params *gainer.Params) ([]gainer.CryptocurrencyData, error) {
	gainer := gainer.New(c.client)
	return gainer.FetchGainersLosers(params)
}

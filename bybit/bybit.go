package bybit

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/market"
)

type Bybit interface {
	Market() market.Market
}

type bybitImpl struct {
	market    market.Market
	client    *client.Client
	isTestNet bool
}

func (b bybitImpl) Market() market.Market {
	return b.market
}

func NewBybit(key, secretKey string, isTestNet bool) Bybit {
	c := client.NewClient(key, secretKey)
	c.IsTestNet = isTestNet
	return &bybitImpl{
		isTestNet: isTestNet,
		client:    c,
		market:    market.NewMarket(c),
	}

}

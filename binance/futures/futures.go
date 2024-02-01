package futures

import (
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/market"
)

type Futures interface {
	Market() market.Market
	Account() Account
	Generic() Generic
}

type futureImpl struct {
	client *client.Client
}

func New(apiKey, apiSecret string, isTestnet bool) Futures {
	return &futureImpl{
		client: client.NewFuturesClient(apiKey, apiSecret, isTestnet),
	}
}

func (f *futureImpl) Account() Account {
	return NewAccount(f.client)
}

func (f *futureImpl) Generic() Generic {
	return NewGeneric(f.client)
}

func (f *futureImpl) Market() market.Market {
	return market.NewMarket(f.client)
}

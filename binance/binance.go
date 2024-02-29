package binance

import "github.com/cploutarchou/crypto-sdk-suite/binance/futures"

type Binance interface {
	Futures() futures.Futures
}

type binanceImpl struct {
	apiKey, apiSecret string
	isTestnet         bool
}

func New(apiKey, apiSecret string, isTestnet bool) Binance {
	return &binanceImpl{
		apiKey,
		apiSecret,
		isTestnet,
	}
}

func (b *binanceImpl) Futures() futures.Futures {
	return futures.New(b.apiKey, b.apiSecret, b.isTestnet)
}

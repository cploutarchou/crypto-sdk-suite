package binance

import "github.com/cploutarchou/crypto-sdk-suite/bybit/market"

type Binance interface {
	Market() market.Market
}

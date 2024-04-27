package binance

import "github.com/cploutarchou/crypto-sdk-suite/binance/futures"

// Binance interface represents the operations available for the Binance API.
type Binance interface {
	// Futures returns the interface for Futures operations.
	Futures() futures.Futures
}

// binanceImpl represents the implementation of the Binance interface.
type binanceImpl struct {
	apiKey, apiSecret string
	isTestnet         bool
}

// New creates a new Binance instance with the provided API key, API secret, and testnet flag.
func New(apiKey, apiSecret string, isTestnet bool) Binance {
	return &binanceImpl{
		apiKey,
		apiSecret,
		isTestnet,
	}
}

// Futures returns the Futures interface implementation by creating a new Futures instance.
func (b *binanceImpl) Futures() futures.Futures {
	return futures.New(b.apiKey, b.apiSecret, b.isTestnet)
}

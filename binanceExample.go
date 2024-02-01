package main

import (
	"github.com/cploutarchou/crypto-sdk-suite/binance"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures"
)

var b futures.Futures

func init() {
	dsd := binance.New("API_KEY", "API_SECRET", true)
	b = dsd.Futures()
	err := b.Generic().Ping()
	if err != nil {
		return
	}
}

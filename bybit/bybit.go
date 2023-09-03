package bybit

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/market"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
)

type Bybit interface {
	Market() market.Market
	Ws() ws.WebSocket
}

type bybitImpl struct {
	market    market.Market
	client    *client.Client
	isTestNet bool
	webSocket ws.WebSocket
}

func (b bybitImpl) Market() market.Market {
	return b.market
}

func (b bybitImpl) Ws() ws.WebSocket {
	return b.webSocket
}

func NewBybit(key, secretKey string, isTestNet bool) Bybit {
	c := client.NewClient(key, secretKey)
	wsClient, _ := ws.NewWebSocket(key, secretKey)
	return &bybitImpl{
		market:    market.NewMarket(c),
		client:    c,
		isTestNet: isTestNet,
		webSocket: wsClient,
	}
}

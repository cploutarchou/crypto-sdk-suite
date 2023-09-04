package bybit

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/market"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
	client2 "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
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

func New(key, secretKey string, isTestNet bool) Bybit {
	c := client.NewClient(key, secretKey)
	wsClient, _ := client2.New(key, secretKey, isTestNet, true)
	return &bybitImpl{
		market:    market.New(c),
		client:    c,
		isTestNet: isTestNet,
		webSocket: ws.New().SetClient(wsClient),
	}
}

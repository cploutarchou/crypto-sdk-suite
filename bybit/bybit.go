package bybit

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/market"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
	client2 "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type Bybit interface {
	Market() market.Market
	// WebSocket returns the websocket package for bybit
	WebSocket() ws.WebSocket
}

type bybitImpl struct {
	market    market.Market
	client    *client.Client
	isTestNet bool
	webSocket ws.WebSocket
	apiKey    string
	secretKey string
}

func (b bybitImpl) Market() market.Market {
	return b.market
}

// WebSocket returns the websocket package for bybit
func (b bybitImpl) WebSocket() ws.WebSocket {
	return b.webSocket
}

func New(key, secretKey string, isTestNet bool) Bybit {
	c := client.NewClient(key, secretKey, isTestNet)
	wsClient, _ := client2.New(key, secretKey, isTestNet)
	by := &bybitImpl{
		market:    market.New(c),
		client:    c,
		isTestNet: isTestNet,
		apiKey:    key,
		secretKey: secretKey,
		webSocket: ws.New(wsClient),
	}
	return by
}

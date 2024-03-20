package bybit

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/account"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/asset"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/market"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/position"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/trade"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
	client2 "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type Bybit interface {
	Market() market.Market
	WebSocket() ws.WebSocket
	Account() account.Account
	Trade() trade.Trade
	Position() position.Position
	Asset() asset.Asset
}

type bybitImpl struct {
	market    market.Market
	client    *client.Client
	isTestNet bool
	webSocket ws.WebSocket
	apiKey    string
	secretKey string
	account   account.Account
	trade     trade.Trade
	position  position.Position
	asset     asset.Asset
}

func New(key, secretKey string, isTestNet bool) Bybit {
	c := client.NewClient(key, secretKey, isTestNet)
	wsClient, _ := client2.NewClient(key, secretKey, isTestNet)
	by := &bybitImpl{
		market:    market.New(c),
		account:   account.New(c),
		trade:     trade.New(c),
		position:  position.New(c),
		asset:     asset.New(c),
		client:    c,
		isTestNet: isTestNet,
		apiKey:    key,
		secretKey: secretKey,
		webSocket: ws.New(wsClient, isTestNet),
	}
	return by
}

func (b bybitImpl) Market() market.Market {
	return b.market
}

func (b bybitImpl) WebSocket() ws.WebSocket {
	return b.webSocket
}

func (b bybitImpl) Account() account.Account {
	return b.account
}

func (b bybitImpl) Trade() trade.Trade {
	return b.trade
}

func (b bybitImpl) Position() position.Position {
	return b.position
}

func (b bybitImpl) Asset() asset.Asset {
	return b.asset
}

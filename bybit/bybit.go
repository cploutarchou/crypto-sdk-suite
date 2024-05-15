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

// Market returns the market interface for Bybit operations.
//
// No parameters.
// Returns a market.Market interface.
func (b bybitImpl) Market() market.Market {
	return b.market
}

// WebSocket returns the WebSocket interface for Bybit operations.
//
// No parameters.
// Returns a ws.WebSocket interface.
func (b bybitImpl) WebSocket() ws.WebSocket {
	return b.webSocket
}

// Account returns the Account interface for Bybit operations.
//
// No parameters.
// Returns an account.Account interface.
func (b bybitImpl) Account() account.Account {
	return b.account
}

// Trade returns the Trade interface for Bybit operations.
//
// No parameters.
// Returns a trade.Trade interface.
func (b bybitImpl) Trade() trade.Trade {
	return b.trade
}

// Position returns the Position interface for Bybit operations.
//
// No parameters.
// Returns a position.Position interface.
func (b bybitImpl) Position() position.Position {
	return b.position
}

// Asset returns the Asset interface for Bybit operations.
//
// No parameters.
// Returns an asset.Asset interface.
func (b bybitImpl) Asset() asset.Asset {
	return b.asset
}

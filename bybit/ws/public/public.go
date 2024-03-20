package public

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/kline"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/liquidation"
	ltkline "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/lt-kline"
	ltnav "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/lt-nav"
	ltticker "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/lt-ticker"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/orderbook"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/ticker"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/trade"
)

type Public interface {
	Kline() kline.Kline
	Liquidation() liquidation.Liquidation
	LtKline() ltkline.LtKline
	LtNav() ltnav.LtNav
	LtTickers() ltticker.LtTicker
	OrderBook() orderbook.OrderBook
	Ticker() ticker.Ticker
	Trade() trade.Trade
}

type implPublic struct {
	client *client.Client
}

func (i *implPublic) Kline() kline.Kline {
	return kline.New(i.client)
}
func (i *implPublic) Liquidation() liquidation.Liquidation {
	return *liquidation.New()
}

func (i *implPublic) LtKline() ltkline.LtKline {
	return *ltkline.New()
}

func (i *implPublic) LtNav() ltnav.LtNav {
	return *ltnav.New()
}

func (i *implPublic) LtTickers() ltticker.LtTicker {
	return *ltticker.New()
}

func (i *implPublic) OrderBook() orderbook.OrderBook {
	return *orderbook.New()
}

func (i *implPublic) Ticker() ticker.Ticker {
	return *ticker.New(i.client)
}

func (i *implPublic) Trade() trade.Trade {
	return *trade.New()
}

func New(wsClient *client.Client, category string) Public {
	wsClient.IsPublic = true
	wsClient.Category = category
	return &implPublic{client: wsClient}
}

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
	Kline(category string) (kline.Kline, error)
	Liquidation(category string) liquidation.Liquidation
	LtKline(category string) ltkline.LtKline
	LtNav(category string) ltnav.LtNav
	LtTickers(category string) ltticker.LtTicker
	OrderBook(category string) orderbook.OrderBook
	Ticker(category string) ticker.Ticker
	Trade(category string) trade.Trade
}

type implPublic struct {
	client *client.Client
}

func (i *implPublic) Kline(category string) (kline.Kline, error) {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return kline.New(cli)
}
func (i *implPublic) Liquidation(category string) liquidation.Liquidation {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return liquidation.New(cli)
}

func (i *implPublic) LtKline(category string) ltkline.LtKline {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return ltkline.New(cli)
}

func (i *implPublic) LtNav(category string) ltnav.LtNav {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return ltnav.New(cli)
}

func (i *implPublic) LtTickers(category string) ltticker.LtTicker {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return ltticker.New(cli)
}

func (i *implPublic) OrderBook(category string) orderbook.OrderBook {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return orderbook.New(cli)
}

func (i *implPublic) Ticker(category string) ticker.Ticker {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return ticker.New(cli)
}

func (i *implPublic) Trade(category string) trade.Trade {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return trade.New(cli)
}

func New(wsClient *client.Client, isPPublic bool) Public {
	return &implPublic{client: wsClient}
}

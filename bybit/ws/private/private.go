package private

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/dcp"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/execution"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/greek"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/order"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/position"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private/wallet"
)

type Private interface {
	Dcp(category string) dcp.Dcp
	Execution(category string) execution.Execution
	Greek(category string) greek.Greek
	Order(category string) order.Order
	Position(category string) position.Position
	Wallet(category string) wallet.Wallet
}

type implPrivate struct {
	client *client.Client
	isTest bool
}

func (i *implPrivate) Dcp(category string) dcp.Dcp {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return dcp.New(cli)
}

func (i *implPrivate) Execution(category string) execution.Execution {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return execution.New(cli)
}

func (i *implPrivate) Greek(category string) greek.Greek {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return greek.New(cli)
}

func (i *implPrivate) Order(category string) order.Order {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return order.New(cli)
}

func (i *implPrivate) Position(category string) position.Position {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return position.New(cli)
}

func (i *implPrivate) Wallet(category string) wallet.Wallet {
	cli := new(client.Client)
	cli.Category = category
	cli.ApiKey = i.client.ApiKey
	cli.ApiSecret = i.client.ApiSecret
	return wallet.New(cli)
}

func (i *implPrivate) SetClient(client_ *client.Client) Private {
	if client_ != nil {
		return &implPrivate{
			client: client_,
			isTest: i.isTest,
		}
	} else {
		return nil
	}
}
func New(wsClient *client.Client, isTestnet bool) Private {
	return &implPrivate{client: wsClient, isTest: isTestnet}
}

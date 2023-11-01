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
	Dcp() *dcp.Dcp
	Execution() *execution.Execution
	Greek() *greek.Greek
	Order() *order.Order
	Position() *position.Position
	Wallet() *wallet.Wallet
	SetClient(client_ *client.WSClient) Private
}

type implPrivate struct {
	client *client.WSClient
}

func (i *implPrivate) Dcp() *dcp.Dcp {
	return dcp.New()
}

func (i *implPrivate) Execution() *execution.Execution {
	return execution.New()
}

func (i *implPrivate) Greek() *greek.Greek {
	return greek.New()
}

func (i *implPrivate) Order() *order.Order {
	return order.New()
}

func (i *implPrivate) Position() *position.Position {
	return position.New()
}

func (i *implPrivate) Wallet() *wallet.Wallet {
	return wallet.New()
}

func (i *implPrivate) SetClient(client_ *client.WSClient) Private {
	if client_ != nil {
		return &implPrivate{
			client: client_,
		}
	} else {
		return nil
	}
}
func New(wsClient *client.WSClient) Private {
	return &implPrivate{client: wsClient}
}

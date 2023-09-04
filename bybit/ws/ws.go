package ws

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public"
)

type WebSocket interface {
	Private() (private.Private, error)
	Public() (public.Public, error)
	SetClient(client_ *client.WSClient) WebSocket
}

type implWebSocket struct {
	client *client.WSClient
}

func (i *implWebSocket) SetClient(client_ *client.WSClient) WebSocket {
	if client_ != nil {
		return &implWebSocket{
			client: client_,
		}
	} else {
		return nil
	}
}

func (i *implWebSocket) Private() (private.Private, error) {
	return private.New().SetClient(i.client), nil
}

func (i *implWebSocket) Public() (public.Public, error) {
	cl, err := client.New("", "", false, true)
	if err != nil {
		return nil, err
	}
	i.client = cl
	return public.New().SetClient(cl), nil
}
func New() WebSocket {
	return &implWebSocket{}

}

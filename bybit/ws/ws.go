package ws

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/private"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public"
)

type WebSocket interface {
	Private() private.Private
	Public() public.Public
}

type implWebSocket struct {
}

func (i *implWebSocket) Private() private.Private {
	return private.New()
}

func (i *implWebSocket) Public() public.Public {
	return public.New()
}

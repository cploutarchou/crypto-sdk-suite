package ws

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public"

type WebSocket interface {
	Public() public.Public
}

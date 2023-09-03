package public

import "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"

type Public interface {
	// AddSymbols adds symbols to the subscription list
	AddSymbols(symbols ...string) error
	// Subscribe subscribes to the given topics
	Subscribe(topics ...string) error
	// Unsubscribe unsubscribes from the given topics
	Unsubscribe(topics ...string) error
	// Receive receives messages from the WebSocket connection
	Receive() (int, []byte, error)
	// Close closes the WebSocket connection
	Close()
}

type implPublic struct {
	client   *client.WSClient
	Messages chan []byte
	StopChan chan struct{}
}

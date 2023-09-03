package kline

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type TheKline struct {
	client   *client.WSClient
	Messages chan []byte
}

func (w *TheKline) SetClient(client *client.WSClient) *TheKline {
	w.client = client
	return w
}

func (w *TheKline) Subscribe(topics ...string) error {
	return w.client.Subscribe(topics...)

}

func (w *TheKline) Unsubscribe(topics ...string) error {
	return w.client.Unsubscribe(topics...)
}

func (w *TheKline) Receive() ([]byte, error) {
	return w.client.Receive()
}

func (w *TheKline) Close() {
	w.client.Close()

}

func (w *TheKline) AddSymbols(symbols ...string) error {
	// Prepare the kline subscription message for each symbol
	var subscriptions []string
	for _, symbol := range symbols {
		subscription := fmt.Sprintf("kline.%s", symbol) // Modify this format according to the actual expected format by the WebSocket server.
		subscriptions = append(subscriptions, subscription)
	}

	// Send the kline subscription request
	if err := w.Subscribe(subscriptions...); err != nil {
		return fmt.Errorf("failed to subscribe to kline channel: %v", err)
	}

	// Wait and process the incoming data continuously
	for {
		message, err := w.Receive()
		if err != nil {
			// You might want to handle this error, for example, by trying to reconnect.
			return fmt.Errorf("error receiving message: %v", err)
		}

		// stream data processing to channel
		w.Messages <- message

	}
}

func (w *TheKline) GetMessagesChan() <-chan []byte {
	return w.Messages
}

func New(client *client.WSClient, channel, environment, subChannel string) *TheKline {
	return &TheKline{
		client:   client,
		Messages: make(chan []byte),
	}
}

package kline

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type TheKline struct {
	client   *client.WSClient
	Messages chan []byte
	StopChan chan struct{}
}

func (w *TheKline) SetClient(client_ *client.WSClient) *TheKline {
	w.client = client_
	return w
}

func (w *TheKline) Subscribe(topics ...string) error {
	topicJSON, err := json.Marshal(topics)
	if err != nil {
		return fmt.Errorf("failed to marshal topics: %v", err)
	}

	message := fmt.Sprintf(`{"op":"subscribe","args":%s}`, string(topicJSON))
	return w.client.Conn.WriteMessage(client.WSMessageText, []byte(message))
}

func (w *TheKline) Unsubscribe(topics ...string) error {
	topicJSON, err := json.Marshal(topics)
	if err != nil {
		return fmt.Errorf("failed to marshal topics: %v", err)
	}

	message := fmt.Sprintf(`{"op":"unsubscribe","args":%s}`, string(topicJSON))
	return w.client.Conn.WriteMessage(client.WSMessageText, []byte(message))
}

func (w *TheKline) Receive() (int, []byte, error) {
	return w.client.Conn.ReadMessage()
}

func (w *TheKline) Close() {
	w.client.Close()
}

func (w *TheKline) AddSymbols(symbols ...string) error {
	var subscriptions []string
	for _, symbol := range symbols {
		subscription := fmt.Sprintf("kline.%s", symbol)
		subscriptions = append(subscriptions, subscription)
	}

	if err := w.Subscribe(subscriptions...); err != nil {
		return fmt.Errorf("failed to subscribe to kline channel: %v", err)
	}

	for {
		select {
		case <-w.StopChan:
			return nil
		default:
			_, message, err := w.Receive()
			if err != nil {
				return fmt.Errorf("error receiving message: %v", err)
			}
			w.Messages <- message
		}
	}
}

func (w *TheKline) GetMessagesChan() <-chan []byte {
	return w.Messages
}

func (w *TheKline) Stop() {
	w.StopChan <- struct{}{}
}

func New(client_ *client.WSClient, channel, environment, subChannel string) *TheKline {
	return &TheKline{
		client:   client_,
		Messages: make(chan []byte),
		StopChan: make(chan struct{}),
	}
}

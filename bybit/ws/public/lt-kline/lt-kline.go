package lt_kline

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

// LTKline represents the interface for the LT Kline functionality.
type LTKline interface {
	SetClient(client *client.Client) error
	Subscribe(interval string, symbol string, callback func(response LTKlineResponse)) error
	// Unsubscribe unsubscribes from the specified topics.
	Unsubscribe(topics ...string) error

	// Listen reads the next message from the kline channel.
	Listen() (int, []byte, error)

	// Close closes the connection to the kline channel.
	Close()

	// GetMessagesChan returns a channel that receives messages from the kline channel.
	GetMessagesChan() <-chan []byte

	// Stop stops the kline functionality.
	Stop()
}
type ltKlineImpl struct {
	client   *client.Client
	stopChan chan struct{}
	Messages <-chan []byte
	StopChan chan struct{}
}

func (l *ltKlineImpl) Stop() {
	l.StopChan <- struct{}{}
}

func New(cli *client.Client) LTKline {
	return &ltKlineImpl{
		client:   cli,
		stopChan: make(chan struct{}, 1),
	}
}

func (l *ltKlineImpl) SetClient(c *client.Client) error {
	l.client = c
	return nil
}

func (l *ltKlineImpl) Subscribe(interval string, symbol string, callback func(response LTKlineResponse)) error {
	return l.SubscribeLTKline(interval, symbol, callback)
}

func (l *ltKlineImpl) Close() {
	l.stopChan <- struct{}{}
	l.client.Close()
}
func (l *ltKlineImpl) Unsubscribe(topics ...string) error {
	unsubscription := map[string]any{
		"op":   "unsubscribe",
		"args": topics,
	}
	msg, err := json.Marshal(unsubscription)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	if err := l.client.Send(msg); err != nil {
		return fmt.Errorf("failed to unsubscribe from kline channel: %v", err)
	}

	return nil
}
func (l *ltKlineImpl) Listen() (int, []byte, error) {
	return l.client.Conn.ReadMessage()
}
func (l *ltKlineImpl) GetMessagesChan() <-chan []byte {
	return l.Messages
}

// SubscribeLTKline subscribes to the leveraged token kline stream for the specified interval and symbol.
func (l *ltKlineImpl) SubscribeLTKline(interval string, symbol string, callback func(response LTKlineResponse)) error {
	topic := fmt.Sprintf("kline_lt.%s.%s", interval, symbol)
	subscription := map[string]any{
		"op":   "subscribe",
		"args": []string{topic},
	}

	msg, err := json.Marshal(subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription message: %v", err)
	}

	if err := l.client.Send(msg); err != nil {
		return fmt.Errorf("failed to subscribe to LT kline stream: %v", err)
	}

	// Start a goroutine to listen for messages
	go func() {
		for {
			message, err := l.client.Receive()
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				continue
			}

			var resp LTKlineResponse
			if err := json.Unmarshal(message, &resp); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			if resp.Topic == topic {
				callback(resp)
			}
		}
	}()

	return nil
}

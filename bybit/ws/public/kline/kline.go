package kline

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

// Kline represents the interface for the kline functionality.
type Kline interface {
	// SetClient sets the client for the kline functionality.
	SetClient(client *client.Client) error

	// Subscribe subscribes to kline data for the specified symbols and interval.
	// It also stores the callback for each topic.
	Subscribe(symbols []string, interval string, callback func(response Data)) error

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

// Response struct represents the kline response from the server.
type Response struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  []Data `json:"data"`
	Ts    int64  `json:"ts"`
}

// Data struct represents individual kline data points.
type Data struct {
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	Turnover  string `json:"turnover"`
	Confirm   bool   `json:"confirm"`
	Timestamp int64  `json:"timestamp"`
}

// New creates a new instance of KlineImpl.
func New(c *client.Client) (Kline, error) {
	var k klineImpl
	k.client = c
	k.Messages = make(chan []byte, 100)
	k.StopChan = make(chan struct{}, 1)
	k.isTest = c.IsTestNet
	err := k.client.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	<-k.client.Connected
	fmt.Println("Connected to WS")

	go k.listenForMessages()

	return &k, nil
}

type topicCallback struct {
	callback func(data Data)
}

type klineImpl struct {
	client         *client.Client
	Messages       chan []byte
	StopChan       chan struct{}
	isTest         bool
	topicCallbacks map[string]topicCallback
}

func (k *klineImpl) SetClient(c *client.Client) error {
	k.client = c
	return nil
}

func (k *klineImpl) Subscribe(symbols []string, interval string, callback func(response Data)) error {
	if k.topicCallbacks == nil {
		k.topicCallbacks = make(map[string]topicCallback)
	}

	topics := make([]string, len(symbols))
	for i, symbol := range symbols {
		topic := fmt.Sprintf("kline.%s.%s", interval, symbol)
		topics[i] = topic
		k.topicCallbacks[topic] = topicCallback{callback: callback}
	}

	subscription := map[string]interface{}{
		"op":   "subscribe",
		"args": topics,
	}

	msg, err := json.Marshal(subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription message: %v", err)
	}

	if err := k.client.Send(msg); err != nil {
		return fmt.Errorf("failed to subscribe to kline channel: %v", err)
	}

	return nil
}

func (k *klineImpl) Unsubscribe(topics ...string) error {
	unsubscription := map[string]interface{}{
		"op":   "unsubscribe",
		"args": topics,
	}
	msg, err := json.Marshal(unsubscription)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	if err := k.client.Send(msg); err != nil {
		return fmt.Errorf("failed to unsubscribe from kline channel: %v", err)
	}

	return nil
}

func (k *klineImpl) Listen() (int, []byte, error) {
	return k.client.Conn.ReadMessage()
}

func (k *klineImpl) Close() {
	k.client.Close()
}

func (k *klineImpl) GetMessagesChan() <-chan []byte {
	return k.Messages
}

func (k *klineImpl) Stop() {
	k.StopChan <- struct{}{}
}

func (k *klineImpl) listenForMessages() {
	for {
		select {
		case <-k.StopChan:
			return
		default:
			_, msg, err := k.client.Conn.ReadMessage()
			if err != nil {
				// Handle error, possibly logging and breaking the loop or attempting to reconnect
				return
			}
			k.Messages <- msg

			var resp Response
			if err := json.Unmarshal(msg, &resp); err != nil {
				// Handle unmarshal error
				continue
			}

			if tc, exists := k.topicCallbacks[resp.Topic]; exists {
				for _, data := range resp.Data {
					tc.callback(data)
				}
			}
		}
	}
}

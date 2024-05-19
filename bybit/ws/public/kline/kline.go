package kline

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

// Kline represents the interface for the kline functionality.
type Kline interface {
	SetClient(client *client.Client) error
	Subscribe(symbols []string, interval string, callback func(response Data)) error
	Unsubscribe(topics ...string) error
	Receive() (int, []byte, error)
	Close()
	GetMessagesChan() <-chan []byte
	Stop()
}

// Assuming you want to store callbacks for each topic
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

func (k *klineImpl) Subscribe(symbols []string, interval string, callback func(response Data)) error {

	// Initialize if your implementation doesn't already have this
	if k.topicCallbacks == nil {
		k.topicCallbacks = make(map[string]topicCallback)
	}

	topics := make([]string, len(symbols))
	for i, symbol := range symbols {
		topic := fmt.Sprintf("kline.%s.%s", interval, symbol)
		topics[i] = topic
		// Store the callback for each topic
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
	// Using client's Send method for sending the unsubscription message
	unsubscription := map[string]interface{}{
		"op":   "unsubscribe",
		"args": topics,
	}
	msg, err := json.Marshal(unsubscription)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	// Use the client's Send method to abstract away connection details
	if err := k.client.Send(msg); err != nil {
		return fmt.Errorf("failed to unsubscribe from kline channel: %v", err)
	}

	return nil
}

func (k *klineImpl) Receive() (int, []byte, error) {

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

func (k *klineImpl) SetClient(c *client.Client) error {

	k.client = c
	return nil
}

// New creates a new instance of KlineImpl
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

	go func() {
		for {
			select {
			case <-k.StopChan:
				return
			default:
				_, msg, err := k.client.Conn.ReadMessage()
				if err != nil {
					return
				}
				k.Messages <- msg
			}
		}
	}()
	return &k, nil
}
func (k *klineImpl) listenForMessages() {
	for {
		msg, err := k.client.Receive() // Make sure this method exists and is correctly implemented
		if err != nil {
			// Handle error, possibly breaking the loop or attempting to reconnect
			continue
		}

		var resp Response
		if err := json.Unmarshal(msg, &resp); err != nil {
			// Handle unmarshal error
			continue
		}

		// Find the callback for the topic and execute it with the data
		if tc, exists := k.topicCallbacks[resp.Topic]; exists {
			for _, data := range resp.Data {
				tc.callback(data) // Execute the callback for each item in the data array
			}
		}
	}
}

package kline

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type Interval string

// Kline represents the interface for the kline functionality.
type Kline interface {
	SetClient(client *client.WSClient) (Kline, error)
	Subscribe(symbols []string, interval Interval) (Kline, error)
	Unsubscribe(topics ...string) (Kline, error)
	Receive() (int, []byte, error)
	Close()
	GetMessagesChan() <-chan []byte
	Stop()
}

type klineImpl struct {
	client   *client.WSClient
	Messages chan []byte
	StopChan chan struct{}
	isTest   bool
	mu       sync.Mutex // Mutex to make operations thread-safe
}

func (w *klineImpl) Subscribe(symbols []string, interval Interval) (Kline, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	topics := make([]string, len(symbols))
	for i, symbol := range symbols {
		topics[i] = fmt.Sprintf("kline.%s.%s", interval, symbol)
	}

	topicJSON, err := json.Marshal(topics)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal topics: %w", err)
	}

	message := fmt.Sprintf(`{"op":"subscribe","args":%s}`, string(topicJSON))
	err = w.client.Conn.WriteMessage(client.WSMessageText, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to kline channel: %w", err)
	}

	return w, nil
}

func (w *klineImpl) Unsubscribe(topics ...string) (Kline, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	topicJSON, err := json.Marshal(topics)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal topics: %w", err)
	}

	message := fmt.Sprintf(`{"op":"unsubscribe","args":%s}`, string(topicJSON))
	err = w.client.Conn.WriteMessage(client.WSMessageText, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("failed to unsubscribe to kline channel: %w", err)
	}

	return w, nil
}

func (w *klineImpl) Receive() (int, []byte, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.client.Conn.ReadMessage()
}

func (w *klineImpl) Close() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.client.Close()
}

func (w *klineImpl) GetMessagesChan() <-chan []byte {
	return w.Messages
}

func (w *klineImpl) Stop() {
	w.StopChan <- struct{}{}
}

func (w *klineImpl) SetClient(client_ *client.WSClient) (Kline, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.client = client_
	return w, nil
}

// New creates a new instance of KlineImpl
func New(client_ *client.WSClient, isTestNet bool) Kline {
	client_.Path = "linear"
	var k klineImpl
	k.client = client_
	k.Messages = make(chan []byte, 100)
	k.StopChan = make(chan struct{}, 1)
	k.isTest = isTestNet
	err := k.client.Connect()
	if err != nil {
		fmt.Println("Error connecting:", err)
		return nil
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
					fmt.Println(err)
					return
				}
				k.Messages <- msg
			}
		}
	}()
	return &k
}

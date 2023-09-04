package kline

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type Interval string

const (
	// Minute intervals
	OneMinute      Interval = "1"
	ThreeMinutes   Interval = "3"
	FiveMinutes    Interval = "5"
	FifteenMinutes Interval = "15"
	ThirtyMinutes  Interval = "30"
	SixtyMinutes   Interval = "60"
	TwoHours       Interval = "120"
	FourHours      Interval = "240"
	SixHours       Interval = "360"
	TwelveHours    Interval = "720"

	// Day, week, month intervals
	OneDay   Interval = "D"
	OneWeek  Interval = "W"
	OneMonth Interval = "M"
)

// Kline represents the interface for the kline functionality
type Kline interface {
	SetClient(client *client.WSClient) (Kline, error)
	Subscribe(symbol string, intervals ...Interval) (Kline, error)
	Unsubscribe(topics ...string) (Kline, error)
	Receive() (int, []byte, error)
	Close()
	GetMessagesChan() <-chan []byte
	Stop()
}

// KlineImpl provides the implementation for the Kline interface
type klineImpl struct {
	client        *client.WSClient
	Messages      chan []byte
	StopChan      chan struct{}
	isTest        bool
	interval      Interval
	initialSymbol string
}

func (w *klineImpl) Subscribe(symbol string, intervals ...Interval) (Kline, error) {
	// Creating topic strings for all provided intervals
	topics := make([]string, len(intervals))
	for i, intvl := range intervals {
		topics[i] = fmt.Sprintf("kline.%s.%s", intvl, symbol)
	}

	// Marshalling topics into a JSON string
	topicJSON, err := json.Marshal(topics)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal topics: %w", err)
	}

	// Forming the message with the marshalled topics
	message := fmt.Sprintf(`{"op":"subscribe","args":%s}`, string(topicJSON))

	// Sending the subscription message
	err = w.client.Conn.WriteMessage(client.WSMessageText, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to kline channel: %w", err)
	}

	return w, nil
}

func (w *klineImpl) Unsubscribe(topics ...string) (Kline, error) {
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
	return w.client.Conn.ReadMessage()
}

func (w *klineImpl) Close() {
	w.client.Close()
}

func (w *klineImpl) GetMessagesChan() <-chan []byte {
	return w.Messages
}

func (w *klineImpl) Stop() {
	w.StopChan <- struct{}{}
}

func (w *klineImpl) SetClient(client *client.WSClient) (Kline, error) {
	w.client = client
	return w, nil
}

// New creates a new instance of KlineImpl
func New(client *client.WSClient, symbol string, interval Interval, isTestNet bool) Kline {
	return &klineImpl{
		client:        client,
		Messages:      make(chan []byte, 100),
		StopChan:      make(chan struct{}, 1),
		isTest:        isTestNet,
		interval:      interval,
		initialSymbol: symbol,
	}
}

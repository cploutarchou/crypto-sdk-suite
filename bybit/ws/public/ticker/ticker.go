package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

type response struct {
	Topic string     `json:"topic"`
	Type  string     `json:"type"`
	Data  TickerData `json:"data"`
	Cs    int64      `json:"cs"`
	Ts    int64      `json:"ts"`
}
type TickerData struct {
	Symbol            string `json:"symbol"`
	TickDirection     string `json:"tickDirection"`
	Price24HPcnt      string `json:"price24hPcnt"`
	LastPrice         string `json:"lastPrice"`
	PrevPrice24H      string `json:"prevPrice24h"`
	HighPrice24H      string `json:"highPrice24h"`
	LowPrice24H       string `json:"lowPrice24h"`
	PrevPrice1H       string `json:"prevPrice1h"`
	MarkPrice         string `json:"markPrice"`
	IndexPrice        string `json:"indexPrice"`
	OpenInterest      string `json:"openInterest"`
	OpenInterestValue string `json:"openInterestValue"`
	Turnover24H       string `json:"turnover24h"`
	Volume24H         string `json:"volume24h"`
	NextFundingTime   string `json:"nextFundingTime"`
	FundingRate       string `json:"fundingRate"`
	Bid1Price         string `json:"bid1Price"`
	Bid1Size          string `json:"bid1Size"`
	Ask1Price         string `json:"ask1Price"`
	Ask1Size          string `json:"ask1Size"`
}

// Ticker manages ticker subscriptions and updates.
type Ticker struct {
	client      *client.WSClient
	subscribers map[string]func(TickerData)
	mu          sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
}

// New initializes a new Ticker instance with context for graceful shutdown.
func New(client *client.WSClient) *Ticker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Ticker{
		client:      client,
		subscribers: make(map[string]func(TickerData)),
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Subscribe to the ticker updates for a given symbol.
func (t *Ticker) Subscribe(symbol string, callback func(TickerData)) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	topic := fmt.Sprintf("tickers.%s", symbol)
	t.subscribers[topic] = callback

	// Construct the subscription message
	subscriptionMessage := map[string]interface{}{
		"op":    "subscribe",
		"topic": topic,
	}
	msg, err := json.Marshal(subscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription message: %v", err)
	}

	// Send the subscription message
	return t.client.Send(msg)
}

// Listen method modified to support graceful shutdown.
func (t *Ticker) Listen() {
	for {
		select {
		case <-t.ctx.Done(): // Check if shutdown has been initiated.
			return
		default:
			message, err := t.client.Receive()
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				continue
			}

			var response response
			if err := json.Unmarshal(message, &response); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			t.mu.Lock()
			callback, exists := t.subscribers[response.Topic]
			t.mu.Unlock()

			if exists && (response.Type == "snapshot" || response.Type == "delta") {
				// Execute callback in its own goroutine to not block the Listen loop
				// and to handle callbacks concurrently.
				go callback(response.Data)
			}
		}
	}
}

// Unsubscribe from the ticker updates for a given symbol.
func (t *Ticker) Unsubscribe(symbol string) error {
	topic := fmt.Sprintf("tickers.%s", symbol)
	t.mu.Lock()
	delete(t.subscribers, topic)
	t.mu.Unlock()

	// Construct the unsubscription message
	unsubscriptionMessage := map[string]interface{}{
		"op":    "unsubscribe",
		"topic": topic,
	}
	msg, err := json.Marshal(unsubscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	// Send the unsubscription message
	return t.client.Send(msg)
}

// Shutdown method to cleanly terminate the Listen loop.
func (t *Ticker) Shutdown() {
	t.cancel() // Trigger context cancellation.
}

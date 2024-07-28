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
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  Data   `json:"data"`
	Cs    int64  `json:"cs"`
	Ts    int64  `json:"ts"`
}

type Data struct {
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
	client      *client.Client
	subscribers map[string]func(Data)
	ctx         context.Context
	cancel      context.CancelFunc
	mu          sync.RWMutex
	sendCh      chan []byte
}

// New initializes a new Ticker instance with context for graceful shutdown.
func New(client *client.Client) Ticker {
	ctx, cancel := context.WithCancel(context.Background())
	t := Ticker{
		client:      client,
		subscribers: make(map[string]func(Data)),
		ctx:         ctx,
		cancel:      cancel,
		sendCh:      make(chan []byte),
	}

	go t.writer()

	return t
}

// writer is a goroutine that handles all outgoing messages to the WebSocket connection.
func (t *Ticker) writer() {
	for msg := range t.sendCh {
		err := t.client.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

// Subscribe to the ticker updates for a given symbol.
func (t *Ticker) Subscribe(symbol string, callback func(Data)) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	topic := fmt.Sprintf("tickers.%s", symbol)
	t.subscribers[topic] = callback

	// Correctly construct the subscription message with "args"
	subscriptionMessage := map[string]interface{}{
		"op":   "subscribe",
		"args": []string{topic}, // Use an array for topics as per API requirements
	}
	msg, err := json.Marshal(subscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription message: %v", err)
	}

	// Send the subscription message
	t.sendCh <- msg
	return nil
}

// Listen method modified to support graceful shutdown.
func (t *Ticker) Listen() {
	for {
		select {
		case <-t.ctx.Done(): // Check if shutdown has been initiated.
			close(t.sendCh)
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

			t.mu.RLock()
			callback, exists := t.subscribers[response.Topic]
			t.mu.RUnlock()

			if exists && (response.Type == "snapshot" || response.Type == "delta") {
				go callback(response.Data)
			}
		}
	}
}

// Unsubscribe from the ticker updates for a given symbol.
func (t *Ticker) Unsubscribe(symbol string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	topic := fmt.Sprintf("tickers.%s", symbol)

	delete(t.subscribers, topic)

	// Construct the unsubscription message
	unsubscriptionMessage := map[string]interface{}{
		"op":   "unsubscribe",
		"args": []string{topic},
	}
	msg, err := json.Marshal(unsubscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	// Send the unsubscription message
	t.sendCh <- msg
	return nil
}

// Shutdown method to cleanly terminate the Listen loop.
func (t *Ticker) Shutdown() {
	t.cancel() // Trigger context cancellation.
}

//func main() {
//	// Initialize client with increased buffer sizes
//	client, err := client.NewClient("wss://stream.bybit.com/v5/public/linear")
//	if err != nil {
//		log.Fatalf("Failed to connect to WebSocket: %v", err)
//	}
//
//	ticker := New(client)
//
//	// Example usage
//	err = ticker.Subscribe("1000TURBOUSDT", func(data Data) {
//		fmt.Printf("Received data: %+v\n", data)
//	})
//	if err != nil {
//		log.Fatalf("Failed to subscribe: %v", err)
//	}
//
//	go ticker.Listen()
//
//	// Use the ticker...
//}

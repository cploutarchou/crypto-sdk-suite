package liquidation

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

var oneHundred = 100

// Liquidation represents the interface for the liquidation functionality.
type Liquidation interface {
	// SetClient sets the client for the liquidation functionality.
	SetClient(client *client.Client) error

	// Subscribe subscribes to liquidation data for the specified symbols.
	// It also stores the callback for each topic.
	Subscribe(symbols []string, callback func(response Data)) error

	// Unsubscribe unsubscribes from the specified topics.
	Unsubscribe(topics ...string) error

	// Listen reads the next message from the liquidation channel.
	Listen() (int, []byte, error)

	// Close closes the connection to the liquidation channel.
	Close()

	// GetMessagesChan returns a channel that receives messages from the liquidation channel.
	GetMessagesChan() <-chan []byte

	// Stop stops the liquidation functionality.
	Stop()
}

// Response struct represents the liquidation response from the server.
type Response struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  Data   `json:"data"`
	TS    int64  `json:"ts"`
}

// Data struct represents individual liquidation data points.
type Data struct {
	UpdatedTime int64  `json:"updatedTime"`
	Symbol      string `json:"symbol"`
	Size        string `json:"size"`
	Price       string `json:"price"`
	Side        string `json:"side"`
}

// New creates a new instance of LiquidationImpl.
func New(cli *client.Client) Liquidation {
	var l liquidationImpl
	l.client = cli
	l.Messages = make(chan []byte, oneHundred)
	l.StopChan = make(chan struct{}, 1)
	l.isTest = cli.IsTestNet
	err := l.client.Connect()
	if err != nil {
		fmt.Printf("Failed to connect: %v", err)
	}

	<-l.client.Connected
	fmt.Println("Connected to WS")

	go l.listenForMessages()

	return &l
}

type topicCallback struct {
	callback func(data Data)
}

type liquidationImpl struct {
	client         *client.Client
	Messages       chan []byte
	StopChan       chan struct{}
	isTest         bool
	topicCallbacks map[string]topicCallback
}

func (l *liquidationImpl) SetClient(c *client.Client) error {
	l.client = c
	return nil
}

func (l *liquidationImpl) Subscribe(symbols []string, callback func(response Data)) error {
	if l.topicCallbacks == nil {
		l.topicCallbacks = make(map[string]topicCallback)
	}

	topics := make([]string, len(symbols))
	for i, symbol := range symbols {
		topic := fmt.Sprintf("liquidation.%s", symbol)
		topics[i] = topic
		l.topicCallbacks[topic] = topicCallback{callback: callback}
	}

	subscription := map[string]any{
		"op":   "subscribe",
		"args": topics,
	}

	msg, err := json.Marshal(subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription message: %v", err)
	}

	if err := l.client.Send(msg); err != nil {
		return fmt.Errorf("failed to subscribe to liquidation channel: %v", err)
	}

	return nil
}

func (l *liquidationImpl) Unsubscribe(topics ...string) error {
	unsubscription := map[string]any{
		"op":   "unsubscribe",
		"args": topics,
	}
	msg, err := json.Marshal(unsubscription)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscription message: %v", err)
	}

	if err := l.client.Send(msg); err != nil {
		return fmt.Errorf("failed to unsubscribe from liquidation channel: %v", err)
	}

	return nil
}

func (l *liquidationImpl) Listen() (int, []byte, error) {
	return l.client.Conn.ReadMessage()
}

func (l *liquidationImpl) Close() {
	l.client.Close()
}

func (l *liquidationImpl) GetMessagesChan() <-chan []byte {
	return l.Messages
}

func (l *liquidationImpl) Stop() {
	l.StopChan <- struct{}{}
}

func (l *liquidationImpl) listenForMessages() {
	for {
		select {
		case <-l.StopChan:
			return
		default:
			_, msg, err := l.client.Conn.ReadMessage()
			if err != nil {
				continue
			}
			l.Messages <- msg

			var resp Response
			if err := json.Unmarshal(msg, &resp); err != nil {
				continue
			}

			if tc, exists := l.topicCallbacks[resp.Topic]; exists {
				tc.callback(resp.Data)
			}
		}
	}
}

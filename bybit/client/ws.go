package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DefaultReqID                     = "100001"
	PingOperation                    = "ping"
	SubscribeOperation               = "subscribe"
	UnsubscribeOperation             = "unsubscribe"
	AuthOperation                    = "auth"
	MainNetEnvironment               = "localhost:8080"
	DefaultScheme                    = "wss"
	LocalhostScheme                  = "ws"
	MainNet              Environment = "stream.bybit.com"
	TestNet              Environment = "stream-testnet.bybit.com"

	Public              ChannelType = "public"
	Private             ChannelType = "private"
	ApiV5               ChannelType = "v5"
	Spot                SubChannel  = "spot"
	Linear              SubChannel  = "linear"
	Inverse             SubChannel  = "inverse"
	Option              SubChannel  = "option"
	PingInterval                    = 20 * time.Second
	ReconnectionRetries             = 3
	ReconnectionDelay               = 2 * time.Second
)

type WSClient struct {
	conn      *websocket.Conn
	mu        sync.Mutex
	readMu    sync.Mutex
	closeOnce sync.Once
	isClosed  bool
	logger    *log.Logger

	env        Environment
	channel    ChannelType
	subChannel string
}

func NewWSClient(env Environment, channel ChannelType, subChannel string) (*WSClient, error) {
	client := &WSClient{
		logger:     log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		env:        env,
		channel:    channel,
		subChannel: subChannel,
	}
	err := client.connect()
	if err != nil {
		return nil, err
	}

	go client.keepAlive()
	return client, nil
}

func (c *WSClient) connect() error {
	var scheme string
	switch c.env {
	case MainNetEnvironment:
		scheme = LocalhostScheme
	default:
		scheme = DefaultScheme
	}
	url := fmt.Sprintf("%s://%s/%s/%s/%s", scheme, c.env, ApiV5, c.channel, c.subChannel)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to dial %s: %v", url, err)
	}

	c.conn = conn
	c.logger.Printf("Connected to %s", url)
	return nil
}
func (c *WSClient) Subscribe(topics ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	subRequest := map[string]interface{}{
		"op":   SubscribeOperation,
		"args": topics,
	}
	jsonData, err := json.Marshal(subRequest)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, jsonData)
}

func (c *WSClient) keepAlive() {
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		if c.isClosed {
			c.mu.Unlock()
			return
		}

		pingMsg := WSPingMsg{
			ReqId: DefaultReqID,
			Op:    PingOperation,
		}
		jsonData, err := json.Marshal(pingMsg)
		if err != nil {
			c.logger.Printf("Error marshaling ping message: %v", err)
		}

		if err := c.conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			c.logger.Printf("Error sending ping: %v", err)
			c.conn.Close()

			// Reconnect logic
			for i := 0; i < ReconnectionRetries; i++ {
				if err := c.connect(); err == nil {
					break
				}
				time.Sleep(ReconnectionDelay)
			}
		}
		c.mu.Unlock()
	}
}

func (c *WSClient) Authenticate(apiKey, expires, signature string) error {
	c.logger.Printf("Authenticating with apiKey %s", apiKey)
	authRequest := map[string]interface{}{
		"op":   AuthOperation,
		"args": []interface{}{apiKey, expires, signature},
	}
	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteMessage(websocket.TextMessage, jsonData)
}

func (c *WSClient) ReadMessage() (*SuccessResponse, error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	var response SuccessResponse
	err = json.Unmarshal(message, &response)
	if err != nil {
		c.logger.Printf("Error reading message: %v", err)
		return nil, err
	}
	return &response, nil
}

func (c *WSClient) Close() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.isClosed = true
		c.logger.Println("Connection closed")
		c.conn.Close()
	})
}

func (c *WSClient) Unsubscribe(topics ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	unsubRequest := map[string]interface{}{
		"op":   UnsubscribeOperation,
		"args": topics,
	}
	jsonData, err := json.Marshal(unsubRequest)
	if err != nil {
		return err
	}

	return c.conn.WriteMessage(websocket.TextMessage, jsonData)
}

func (c *WSClient) Receive() ([]byte, error) {
	c.readMu.Lock()
	defer c.readMu.Unlock()

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

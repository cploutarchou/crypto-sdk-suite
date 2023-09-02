package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MainNet Environment = "stream.bybit.com"
	TestNet Environment = "stream-testnet.bybit.com"

	Public              ChannelType = "public"
	Private             ChannelType = "private"
	ApiV5               ChannelType = "v5"
	Spot                SubChannel  = "spot"
	Linear              SubChannel  = "linear"
	Inverse             SubChannel  = "inverse"
	Option              SubChannel  = "option"
	PingInterval                    = 20 * time.Second // Change ping interval to 20 seconds.
	ReconnectionRetries             = 3
	ReconnectionDelay               = 2 * time.Second
)

type Client struct {
	conn      *websocket.Conn
	mu        sync.Mutex
	readMu    sync.Mutex
	closeOnce sync.Once
	isClosed  bool
	logger    *log.Logger
}

func NewClient(env Environment, channel ChannelType, subChannel string) (*Client, error) {
	client := &Client{
		logger: log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
	}
	err := client.connect(env, channel, subChannel)
	if err != nil {
		return nil, err
	}

	go client.keepAlive()
	return client, nil
}

func (c *Client) connect(env Environment, channel ChannelType, subChannel string) error {
	var scheme string
	if env == "localhost:8080" {
		scheme = "ws"
	} else {
		scheme = "wss"
	}
	url := fmt.Sprintf("%s://%s/%s/%s/%s", scheme, env, ApiV5, channel, subChannel)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	c.conn = conn
	c.logger.Printf("Connected to %s", url)
	return nil
}

func (c *Client) Subscribe(topics ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	subRequest := map[string]interface{}{
		"op":   "subscribe",
		"args": topics,
	}
	jsonData, err := json.Marshal(subRequest)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, jsonData)
}
func (c *Client) keepAlive() {
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		if c.isClosed {
			c.mu.Unlock()
			return
		}

		// Send custom ping message instead of websocket's built-in ping
		pingMsg := PingMsg{
			ReqId: "100001",
			Op:    "ping",
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
				if err := c.connect(MainNet, Public, string(Spot)); err == nil {
					break
				}
				time.Sleep(ReconnectionDelay)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Client) Authenticate(apiKey, expires, signature string) error {
	c.logger.Printf("Authenticating with apiKey %s", apiKey)
	authRequest := map[string]interface{}{
		"op":   "auth",
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

func (c *Client) ReadMessage() (*SuccessResponse, error) {
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

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.isClosed = true
		c.logger.Println("Connection closed")
		c.conn.Close()
	})
}

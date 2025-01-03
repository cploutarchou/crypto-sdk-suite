package client

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	DefaultScheme       = "wss"
	PingInterval        = 20 * time.Second
	PingOperation       = "ping"
	AuthOperation       = "auth"
	ReconnectionRetries = 3
	ReconnectionDelay   = 10 * time.Second
	Public              = "public"
	Private             = "private"
)

var (
	DefaultReqID = randomString(eightNumber)
)

const eightNumber = 8

// PingMsg represents the WebSocket ping message format.
type PingMsg struct {
	Op    string `json:"op"`
	ReqID string `json:"req_id,omitempty"`
}

// ChannelType defines the types of channels (public/private) that the WebSocket client can connect to.
type ChannelType string

// Client is the main WebSocket client struct, managing the connection and its state.
type Client struct {
	closeOnce         sync.Once
	connOnce          sync.Once
	isClosed          bool
	logger            *log.Logger
	IsTestNet         bool
	APIKey            string
	APISecret         string
	Channel           ChannelType
	Path              string
	Connected         chan struct{}
	OnConnected       func()
	OnConnectionError func(err error)
	Category          string
	MaxActiveTime     string
	wsURL             string // WebSocket URL for dependency injection in tests

	Conn     *websocket.Conn
	connLock sync.Mutex
}

// NewPublicClient initializes a new public WSClient instance.
func NewPublicClient(isTestNet bool, category string) (*Client, error) {
	client := &Client{
		logger:    log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		IsTestNet: isTestNet,
		Channel:   Public,
		Connected: make(chan struct{}),
		Category:  category,
	}
	DefaultReqID = randomString(eightNumber)
	return client, nil
}

// NewPrivateClient initializes a new private WSClient instance.
func NewPrivateClient(apiKey, apiSecret string, isTestNet bool, maxActiveTime string, category string) (*Client, error) {
	client := &Client{
		logger:        log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		IsTestNet:     isTestNet,
		APIKey:        apiKey,
		APISecret:     apiSecret,
		Channel:       Private,
		Connected:     make(chan struct{}),
		MaxActiveTime: maxActiveTime,
		Category:      category,
	}
	DefaultReqID = randomString(eightNumber)
	return client, nil
}

// Connect establishes a WebSocket connection to the server based on the configuration.
func (c *Client) Connect() error {
	var err error
	c.connOnce.Do(func() {
		c.connLock.Lock()
		defer c.connLock.Unlock()

		if c.isClosed {
			err = errors.New("connection already closed")
			c.handleConnectionError(err)
			return
		}

		url := c.buildURL()
		c.Conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			c.handleConnectionError(fmt.Errorf("failed to dial %s: %v", url, err))
			c.Conn = nil
			return
		}

		c.logger.Printf("Connected to %s", url)
		if c.OnConnected != nil {
			c.OnConnected()
		}
		closeOnce(c.Connected)

		go c.keepAlive()
	})
	return err
}

// buildURL constructs the WebSocket URL based on client configuration.
func (c *Client) buildURL() string {
	if c.wsURL != "" {
		return c.wsURL
	}

	var baseURL string
	if c.IsTestNet {
		baseURL = "stream-testnet.bybit.com"
	} else {
		baseURL = "stream.bybit.com"
	}

	switch c.Channel {
	case Public:
		switch c.Category {
		case "spot":
			return fmt.Sprintf("%s://%s/v5/public/spot", DefaultScheme, baseURL)
		case "usdt_contract", "usdc_contract", "usdc_futures":
			return fmt.Sprintf("%s://%s/v5/public/linear", DefaultScheme, baseURL)
		case "inverse_contract":
			return fmt.Sprintf("%s://%s/v5/public/inverse", DefaultScheme, baseURL)
		case "usdc_option":
			return fmt.Sprintf("%s://%s/v5/public/option", DefaultScheme, baseURL)
		default:
			return fmt.Sprintf("%s://%s/v5/public/linear", DefaultScheme, baseURL) // default to linear (USDT/USDC)
		}
	case Private:
		return fmt.Sprintf("%s://%s/v5/private", DefaultScheme, baseURL)
	default:
		return fmt.Sprintf("%s://%s/v5/public/linear", DefaultScheme, baseURL) // default URL
	}
}

// authenticateIfRequired authenticates the WebSocket client if the channel is private.
func (c *Client) authenticateIfRequired() error {
	if c.Channel == Private {
		expires := fmt.Sprintf("%d", time.Now().UnixMilli()+1000)
		signatureData := fmt.Sprintf("GET/realtime%s", expires)
		signed := GenerateWsSignature(c.APISecret, signatureData)
		c.logger.Printf("Authenticating with apiKey %s, expires %s, signed %s", c.APIKey, expires, signed)
		return c.Authenticate(c.APIKey, expires, signed)
	}
	return nil
}

// GenerateWsSignature generates a signature for the WebSocket API.
func GenerateWsSignature(apiSecret, data string) string {
	if data == "" {
		return ""
	}
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// keepAlive sends a ping message to the WebSocket server every PingInterval and handles reconnection if the ping fails.
func (c *Client) keepAlive() {
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.sendPingAndHandleReconnection()
	}
}

// sendPingAndHandleReconnection sends a ping message to the WebSocket server and handles reconnection if the ping fails.
func (c *Client) sendPingAndHandleReconnection() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.isClosed || c.Conn == nil {
		return
	}

	pingMsg := PingMsg{
		ReqID: DefaultReqID,
		Op:    PingOperation,
	}
	jsonData, err := json.Marshal(pingMsg)
	if err != nil {
		c.logger.Printf("Error marshaling ping message: %v", err)
		return
	}

	if err = c.Conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		c.logger.Printf("Error sending ping: %v", err)
		go c.handleReconnection()
		return
	}
	c.logger.Println("Ping sent")
}

// Authenticate sends an authentication request to the WebSocket server.
func (c *Client) Authenticate(apiKey, expires, signature string) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.Channel != Private {
		return errors.New("cannot authenticate on a public channel")
	}
	c.logger.Printf("Authenticating with apiKey %s, expires %s, signed %s", apiKey, expires, signature)
	authRequest := map[string]any{
		"op":   AuthOperation,
		"args": []any{apiKey, expires, signature},
	}
	jsonData, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}
	if err := c.Conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		c.handleConnectionError(err)
		return err
	}
	return nil
}

// Close gracefully closes the WebSocket connection.
func (c *Client) Close() {
	c.closeOnce.Do(func() {
		c.connLock.Lock()
		defer c.connLock.Unlock()

		c.isClosed = true
		c.logger.Println("Connection closed")
		if c.Conn != nil {
			if err := c.Conn.Close(); err != nil && c.OnConnectionError != nil {
				c.OnConnectionError(err)
			}
			c.Conn = nil
		}
	})
}

// randomString generates a random string of specified length.
func randomString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// Send sends a message to the WebSocket server.
func (c *Client) Send(message []byte) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.isClosed {
		return errors.New("attempt to send message on closed connection")
	}

	if c.Conn == nil {
		log.Println("Connection is nil, attempting to reconnect...")
		if err := c.Connect(); err != nil {
			log.Printf("Reconnection failed: %v", err)
			return err
		}
	}

	if c.Conn == nil {
		return errors.New("connection is still nil after attempting to reconnect")
	}

	if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	return nil
}

// Receive listens for a message from the WebSocket server and returns it.
func (c *Client) Receive() ([]byte, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.Conn == nil {
		return nil, errors.New("attempt to receive message on nil connection")
	}

	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		log.Printf("Error receiving message: %v", err)
		go c.handleReconnection()
		return nil, err
	}

	return message, nil
}

// handleReconnection attempts to reconnect to the WebSocket server.
func (c *Client) handleReconnection() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.isClosed {
		return // No need to reconnect if the client is intentionally closed
	}

	c.logger.Println("Attempting to reconnect...")
	if c.Conn != nil {
		_ = c.Conn.Close()
		c.Conn = nil
	}

	for i := 0; i < ReconnectionRetries; i++ {
		time.Sleep(ReconnectionDelay)
		if err := c.Connect(); err == nil {
			c.logger.Printf("Reconnection attempt %d successful", i+1)
			return
		}
		c.logger.Printf("Reconnection attempt %d failed", i+1)
	}
}

func (c *Client) handleConnectionError(err error) {
	if c.OnConnectionError != nil {
		c.OnConnectionError(err)
	}
	c.logger.Printf("Connection error: %v", err)
}

// closeOnce ensures the channel is only closed once
func closeOnce(ch chan struct{}) {
	select {
	case <-ch:
		// Channel is already closed
	default:
		close(ch)
	}
}

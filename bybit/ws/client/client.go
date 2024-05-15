package client

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DefaultScheme = "wss"
	ApiV5         = "v5"
	PingInterval  = 20 * time.Second

	PingOperation                   = "ping"
	AuthOperation                   = "auth"
	ReconnectionRetries             = 3
	ReconnectionDelay               = 10 * time.Second
	Public              ChannelType = "public"
	Private             ChannelType = "private"
)

var DefaultReqID = randomString(8)

// PingMsg represents the WebSocket ping message format.
type PingMsg struct {
	Op    string `json:"op"`
	ReqId string `json:"req_id,omitempty"`
}

// ChannelType defines the types of channels (public/private) that the WebSocket client can connect to.
type ChannelType string

// Client is the main WebSocket client struct, managing the connection and its state.
type Client struct {
	Conn              *websocket.Conn
	closeOnce         sync.Once
	isClosed          bool
	logger            *log.Logger
	IsTestNet         bool
	IsPublic          bool
	ApiKey            string
	ApiSecret         string
	Channel           ChannelType
	Path              string
	Connected         chan struct{}
	OnConnected       func()
	OnConnectionError func(err error)
	Category          string
	MaxActiveTime     string
	mu                sync.Mutex
	wsURL             string // WebSocket URL for dependency injection in tests
}

// NewClient initializes a new WSClient instance.
func NewClient(apiKey, apiSecret string, isTestNet, isPublic bool, maxActiveTime string) (*Client, error) {
	client := &Client{
		logger:        log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		IsTestNet:     isTestNet,
		ApiKey:        apiKey,
		ApiSecret:     apiSecret,
		IsPublic:      isPublic,
		Connected:     make(chan struct{}),
		MaxActiveTime: maxActiveTime,
	}
	DefaultReqID = randomString(8)
	return client, nil
}

// Connect establishes a WebSocket connection to the server based on the configuration.
func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed {
		err := errors.New("connection already closed")
		c.handleConnectionError(err)
		return err
	}

	url := c.buildURL()
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		c.handleConnectionError(fmt.Errorf("failed to dial %s: %v", url, err))
		return err
	}

	c.Conn = conn
	close(c.Connected)
	c.logger.Printf("Connected to %s", url)
	if c.OnConnected != nil {
		c.OnConnected()
	}
	go c.keepAlive()

	// Authenticate if required
	if c.Channel == Private {
		if err := c.authenticateIfRequired(); err != nil {
			return err
		}
	}

	return nil
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

	channelType := "public"
	if c.Channel == Private {
		channelType = "private"
	}

	url := fmt.Sprintf("%s://%s/%s/%s", DefaultScheme, baseURL, ApiV5, channelType)
	if c.IsPublic {
		switch c.Category {
		case "spot":
			url += "/spot"
		case "linear":
			url += "/linear"
		case "inverse":
			url += "/inverse"
		case "option":
			url += "/option"
		default:
			url += "/spot" // default to spot
		}
	}
	if c.MaxActiveTime != "" && c.Channel == Private {
		url += fmt.Sprintf("?max_active_time=%s", c.MaxActiveTime)
	}
	return url
}

// authenticateIfRequired authenticates the WebSocket client if the channel is private.
func (c *Client) authenticateIfRequired() error {
	if c.Channel == Private {
		expires := fmt.Sprintf("%d", time.Now().UnixMilli()+1000)
		signatureData := fmt.Sprintf("GET/realtime%s", expires)
		signed := GenerateWsSignature(c.ApiSecret, signatureData)
		return c.Authenticate(c.ApiKey, expires, signed)
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
	if c.isClosed {
		return
	}

	pingMsg := PingMsg{
		ReqId: DefaultReqID,
		Op:    PingOperation,
	}
	jsonData, err := json.Marshal(pingMsg)
	if err != nil {
		c.logger.Printf("Error marshaling ping message: %v", err)
		return
	}

	if err = c.Conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		c.logger.Printf("Error sending ping: %v", err)
		c.handleReconnection()
	}
	c.logger.Println("Ping sent")
}

// Authenticate sends an authentication request to the WebSocket server.
func (c *Client) Authenticate(apiKey, expires, signature string) error {
	if c.Channel != Private {
		return errors.New("cannot authenticate on a public channel")
	}
	c.logger.Printf("Authenticating with apiKey %s", apiKey)
	authRequest := map[string]interface{}{
		"op":   AuthOperation,
		"args": []interface{}{apiKey, expires, signature},
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
		c.mu.Lock()
		defer c.mu.Unlock()

		if c.isClosed {
			return
		}
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
	c.mu.Lock()
	defer c.mu.Unlock()

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
	fmt.Println(string(message))

	if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	return nil
}

// Receive listens for a message from the WebSocket server and returns it.
func (c *Client) Receive() ([]byte, error) {
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		log.Printf("Error receiving message: %v", err)
		return nil, err
	}

	fmt.Println(string(message))
	return message, nil
}

// handleReconnection attempts to reconnect to the WebSocket server.
func (c *Client) handleReconnection() {
	c.mu.Lock()
	defer c.mu.Unlock()

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

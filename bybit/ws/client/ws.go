// Package client provides a WebSocket client to interact with a server.
// It manages authentication, pinging, and reconnection logic.
package client

import (
	"crypto/hmac"
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
	DefaultScheme                   = "wss"
	ApiV5                           = "v5"
	PingInterval                    = 30 * time.Second
	DefaultReqID                    = "some_default_request_id"
	PingOperation                   = "ping"
	AuthOperation                   = "auth"
	ReconnectionRetries             = 3
	ReconnectionDelay               = 5 * time.Second
	Public              ChannelType = "public"
	Private             ChannelType = "private"
)

// WSPingMsg represents the WebSocket ping message format.
type WSPingMsg struct {
	Op    string `json:"op"`
	ReqId string `json:"req_id"`
}

// WSPublicChannels represents the type for public channels.
type WSPublicChannels string

// WSPrivateChannels represents the type for private channels.
type WSPrivateChannels string

// WSClient is the main WebSocket client struct, managing the connection and its state.
type WSClient struct {
	Conn      *websocket.Conn
	mu        sync.Mutex
	closeOnce sync.Once
	isClosed  bool
	logger    *log.Logger
	isTestNet bool
	isPublic  bool
	apiKey    string
	apiSecret string
	IsTest    bool
	channel   ChannelType
	Path      string
	Connected chan struct{}
}

// ChannelType defines the types of channels (public/private) that the WebSocket client can connect to.
type ChannelType string

// New initializes a new WSClient instance.
// apiKey and apiSecret are required for authentication with private channels.
// isTestNet determines if the client connects to a test network.
// isPublic specifies if the client connects to a public channel.
func New(apiKey, apiSecret string, isTestNet, isPublic bool) (*WSClient, error) {
	client := &WSClient{
		logger:    log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		isTestNet: isTestNet,
		isPublic:  isPublic,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		Path:      "",
		Connected: make(chan struct{}),
	}

	return client, nil
}

// Connect establishes a WebSocket connection to the server based on the configuration.
func (c *WSClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed {
		return errors.New("connection already closed")
	}

	if c.channel == Private {
		if err := c.authenticateIfRequired(); err != nil {
			return err
		}
	}
	var url string
	if c.isTestNet {
		url = "stream-testnet.bybit.com"
	} else {
		url = "stream.bybit.com"
	}
	if c.isPublic {
		c.channel = Public
	} else {
		c.channel = Private
	}

	url = fmt.Sprintf("%s://%s/%s/%s", DefaultScheme, url, ApiV5, c.channel)

	if c.Path != "" {
		url = fmt.Sprintf("%s/%s", url, c.Path)
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to dial %s: %v", url, err)
	}

	c.Conn = conn
	if c.Connected != nil {
		close(c.Connected)
		c.Connected = nil
	}
	c.logger.Printf("Connected to %s", url)
	go c.keepAlive()
	return nil
}

// authenticateIfRequired authenticates the WebSocket client if the channel is private.
func (c *WSClient) authenticateIfRequired() error {
	if c.channel == Private {
		expires := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)+1)
		signatureData := fmt.Sprintf("GET/realtime%s", expires)
		signed := GenerateWsSignature(c.apiSecret, signatureData)
		return c.Authenticate(c.apiKey, expires, signed)
	}
	return nil
}

// GenerateWsSignature generates a signature for the WebSocket API.
func GenerateWsSignature(apiSecret, data string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// keepAlive sends a ping message to the WebSocket server every PingInterval and handles reconnection if the ping fails.
func (c *WSClient) keepAlive() {
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.sendPingAndHandleReconnection()
	}
}

// sendPingAndHandleReconnection sends a ping message to the WebSocket server
// and handles reconnection if the ping fails.
func (c *WSClient) sendPingAndHandleReconnection() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed {
		return
	}

	pingMsg := WSPingMsg{
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

// handleReconnection closes the current connection and attempts to reconnect to the WebSocket server.
func (c *WSClient) handleReconnection() {
	_ = c.Conn.Close()

	for i := 0; i < ReconnectionRetries; i++ {
		if err := c.Connect(); err == nil {
			break
		}
		time.Sleep(ReconnectionDelay)
	}
}

// Authenticate sends an authentication request to the WebSocket server.
func (c *WSClient) Authenticate(apiKey, expires, signature string) error {
	if c.channel != Private {
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

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Conn.WriteMessage(websocket.TextMessage, jsonData)
}

// Close gracefully closes the WebSocket connection.
func (c *WSClient) Close() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if c.isClosed {
			return
		}
		c.isClosed = true
		c.logger.Println("Connection closed")
		if c.Conn != nil {
			err := c.Conn.Close()
			if err != nil {
				return
			}
			c.Conn = nil
		}
	})
}

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

// WSPingMsg represents the structure for WebSocket ping messages
type WSPingMsg struct {
	Op    string `json:"op"`
	ReqId string `json:"req_id"`
}

// WSPublicChannels represent public channels for WebSocket connection
type WSPublicChannels string

// WSPrivateChannels represent private channels for WebSocket connection
type WSPrivateChannels string

type WSClient struct {
	Conn      *websocket.Conn
	mu        sync.Mutex
	readMu    sync.Mutex
	closeOnce sync.Once
	isClosed  bool
	logger    *log.Logger
	isTestNet bool
	isPublic  bool
	apiKey    string
	apiSecret string
	IsTest    bool
	channel   ChannelType
}

type ChannelType string

// New initializes a new WSClient
func New(apiKey, apiSecret string, isTestNet, isPublic bool) (*WSClient, error) {
	client := &WSClient{
		logger:    log.New(os.Stdout, "[WebSocketClient] ", log.LstdFlags),
		isTestNet: isTestNet,
		isPublic:  isPublic,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
	if err := client.connect(); err != nil {
		return nil, err
	}

	go client.keepAlive()
	return client, nil
}

// connect establishes a connection to the server
func (c *WSClient) connect() error {
	if err := c.authenticateIfRequired(); err != nil {
		return err
	}
	fmt.Println("isTestNet", c.isTestNet)
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
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to dial %s: %v", url, err)
	}

	c.Conn = conn
	c.logger.Printf("Connected to %s", url)
	return nil
}

func (c *WSClient) authenticateIfRequired() error {
	if c.channel == Private {
		expires := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)+1)
		signatureData := fmt.Sprintf("GET/realtime%s", expires)
		signed := GenerateWsSignature(c.apiSecret, signatureData)
		return c.Authenticate(c.apiKey, expires, signed)
	}
	return nil
}

// GenerateWsSignature generates the WebSocket signature
func GenerateWsSignature(apiSecret, data string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// keepAlive sends a periodic ping to maintain the connection
func (c *WSClient) keepAlive() {
	ticker := time.NewTicker(PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.sendPingAndHandleReconnection()
	}
}

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

	if err := c.Conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		c.logger.Printf("Error sending ping: %v", err)
		c.handleReconnection()
	}
}

func (c *WSClient) handleReconnection() {
	c.Conn.Close()
	for i := 0; i < ReconnectionRetries; i++ {
		if err := c.connect(); err == nil {
			break
		}
		time.Sleep(ReconnectionDelay)
	}
}

// Authenticate performs authentication
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

// Close gracefully closes the WebSocket connection
func (c *WSClient) Close() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.isClosed = true
		c.logger.Println("Connection closed")
		c.Conn.Close()
	})
}

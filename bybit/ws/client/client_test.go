// Package client provides functions to interact with the Bybit API.
package client

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// Constants
const (
	testnetBaseURL = "wss://stream-testnet.bybit.com/v5"
)

// Environment variables
var (
	testnetAPIKey    = os.Getenv("BYBIT_FUTURES_TESTNET_API_KEY")
	testnetAPISecret = os.Getenv("BYBIT_FUTURES_TESTNET_API_SECRET")
)

// TestNewPublicClient verifies the NewPublicClient function initializes a public client correctly.
// It tests if the client is initialized with the correct testnet flag and channel type.
func TestNewPublicClient(t *testing.T) {
	client, err := NewPublicClient(true, "usdt_contract")
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, true, client.IsTestNet)
	assert.Equal(t, ChannelType(Public), client.Channel)
}

// TestNewPrivateClient verifies the NewPrivateClient function initializes a private client correctly.
// It tests if the client is initialized with the correct API key, API secret, testnet flag,
// channel type, and max active time.
func TestNewPrivateClient(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret

	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, true, maxActiveTime, "usdt_contract")
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, apiKey, client.ApiKey)
	assert.Equal(t, apiSecret, client.ApiSecret)
	assert.Equal(t, true, client.IsTestNet)
	assert.Equal(t, ChannelType(Private), client.Channel)
	assert.Equal(t, maxActiveTime, client.MaxActiveTime)
}

// TestClient_Connect verifies the Connect function establishes a WebSocket connection correctly.
// It tests if the function connects to the correct WebSocket URL and sets the appropriate callbacks.
func TestClient_Connect(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, true, maxActiveTime, "usdt_contract")
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	client.OnConnected = func() {
		t.Log("Connected to the WebSocket server")
	}

	client.OnConnectionError = func(err error) {
		t.Errorf("Connection error: %v", err)
	}

	err = client.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, client.Conn)
}

// TestClient_Send verifies the Send function sends a message correctly.
// It tests if the function sends a message to the WebSocket server.
func TestClient_Send(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, true, maxActiveTime, "usdt_contract")
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	err = client.Connect()
	assert.NoError(t, err)

	message := []byte(`{"op":"ping","req_id":"test21"}`)
	err = client.Send(message)
	assert.NoError(t, err)
}

// TestClient_Receive verifies the Receiving function receives a message correctly.
// It tests if the function receives a message from the WebSocket server and contains the expected content.
func TestClient_Receive(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, true, maxActiveTime, "usdt_contract")
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	err = client.Connect()
	assert.NoError(t, err)

	// Wait to receive a successful authentication message
	_, authMsg, err := client.Conn.ReadMessage()
	assert.NoError(t, err)
	var authResponse map[string]interface{}
	err = json.Unmarshal(authMsg, &authResponse)
	assert.NoError(t, err)
	assert.False(t, authResponse["success"].(bool), "Authentication failed")

	// Generate a random request ID
	reqID := randomString(4)
	// Simulate sending a pong message from server after successful authentication
	go func() {
		time.Sleep(3 * time.Second)
		requestPayload := `{"op":"ping","req_id":"` + reqID + `"}`
		err = client.Conn.WriteMessage(websocket.TextMessage, []byte(requestPayload))
		if err != nil {
			t.Errorf("Error sending pong message: %v", err)
			return
		}
	}()

	// Now test receiving a pong message
	message, err := client.Receive()
	assert.NoError(t, err)
	assert.Contains(t, string(message), "pong")
}

// TestClient_Close verifies the Close function closes the connection correctly.
// It tests if the function closes the WebSocket connection and sets the appropriate flag.
func TestClient_Close(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, true, maxActiveTime, "usdt_contract")
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	err = client.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, client.Conn)

	client.Close()
	assert.True(t, client.isClosed)
}

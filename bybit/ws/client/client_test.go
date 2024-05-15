package client

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

const (
	testnetBaseURL = "wss://stream-testnet.bybit.com/v5"
)

var (
	testnetAPIKey    = os.Getenv("BYBIT_FUTURES_TESTNET_API_KEY")
	testnetAPISecret = os.Getenv("BYBIT_FUTURES_TESTNET_API_SECRET")
)

// TestNewPublicClient verifies the NewPublicClient function initializes a public client correctly.
func TestNewPublicClient(t *testing.T) {
	isTestNet := true

	client, err := NewPublicClient(isTestNet)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, isTestNet, client.IsTestNet)
	assert.Equal(t, ChannelType(Public), client.Channel)
}

// TestNewPrivateClient verifies the NewPrivateClient function initializes a private client correctly.
func TestNewPrivateClient(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	isTestNet := true
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, isTestNet, maxActiveTime)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, apiKey, client.ApiKey)
	assert.Equal(t, apiSecret, client.ApiSecret)
	assert.Equal(t, isTestNet, client.IsTestNet)
	assert.Equal(t, ChannelType(Private), client.Channel)
	assert.Equal(t, maxActiveTime, client.MaxActiveTime)
}

// TestClient_Connect verifies the Connect function establishes a WebSocket connection correctly.
func TestClient_Connect(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	isTestNet := true
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, isTestNet, maxActiveTime)
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
func TestClient_Send(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	isTestNet := true
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, isTestNet, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	err = client.Connect()
	assert.NoError(t, err)

	message := []byte(`{"op":"ping","req_id":"test_req_id"}`)
	err = client.Send(message)
	assert.NoError(t, err)
}

// TestClient_Receive verifies the Receive function receives a message correctly.
func TestClient_Receive(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	isTestNet := true
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, isTestNet, maxActiveTime)
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
	assert.True(t, authResponse["success"].(bool), "Authentication failed")

	// Generate a random request ID
	reqID := randomString(4)
	// Simulate sending a pong message from server after successful authentication
	go func() {
		time.Sleep(3 * time.Second)
		requestPayload := `{"op":"ping","req_id":"` + reqID + `"}`
		client.Conn.WriteMessage(websocket.TextMessage, []byte(requestPayload))
	}()

	// Now test receiving a pong message
	message, err := client.Receive()
	assert.NoError(t, err)
	expectedMessage := "{\"req_id\":\"" + reqID + "\",\"op\":\"pong\""

	assert.Contains(t, string(message), expectedMessage)
}

// TestClient_Close verifies the Close function closes the connection correctly.
func TestClient_Close(t *testing.T) {
	apiKey := testnetAPIKey
	apiSecret := testnetAPISecret
	isTestNet := true
	maxActiveTime := "1m"

	client, err := NewPrivateClient(apiKey, apiSecret, isTestNet, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = testnetBaseURL + "/private"

	err = client.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, client.Conn)

	client.Close()
	assert.True(t, client.isClosed)
}

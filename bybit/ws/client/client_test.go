package client

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewClient verifies the NewClient function initializes a client correctly.
func TestNewClient(t *testing.T) {
	apiKey := "test_api_key"
	apiSecret := "test_api_secret"
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, true, true, maxActiveTime)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, apiKey, client.ApiKey)
	assert.Equal(t, apiSecret, client.ApiSecret)
	assert.Equal(t, true, client.IsTestNet)
	assert.Equal(t, true, client.IsPublic)
	assert.Equal(t, maxActiveTime, client.MaxActiveTime)
}

// TestClient_Connect verifies the Connect function establishes a WebSocket connection correctly.
func TestClient_Connect(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		assert.NoError(t, err)
		defer conn.Close()

		// Simulate server sending a ping response
		for {
			_, msg, err := conn.ReadMessage()
			assert.NoError(t, err)
			if string(msg) == `{"op":"ping","req_id":"test_req_id"}` {
				conn.WriteMessage(websocket.TextMessage, []byte(`{"op":"pong","req_id":"test_req_id"}`))
			}
		}
	}))
	defer server.Close()

	apiKey := "test_api_key"
	apiSecret := "test_api_secret"
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, true, true, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = "ws" + server.URL[len("http"):]

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
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		assert.NoError(t, err)
		defer conn.Close()

		// Read the message from the client
		for {
			_, msg, err := conn.ReadMessage()
			assert.NoError(t, err)
			if string(msg) == `{"op":"ping","req_id":"test_req_id"}` {
				conn.WriteMessage(websocket.TextMessage, []byte(`{"op":"pong","req_id":"test_req_id"}`))
			}
		}
	}))
	defer server.Close()

	apiKey := "test_api_key"
	apiSecret := "test_api_secret"
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, true, true, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = "ws" + server.URL[len("http"):]

	err = client.Connect()
	assert.NoError(t, err)

	message := []byte(`{"op":"ping","req_id":"test_req_id"}`)
	err = client.Send(message)
	assert.NoError(t, err)
}

// TestClient_Receive verifies the Receive function receives a message correctly.
func TestClient_Receive(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		assert.NoError(t, err)
		defer conn.Close()

		// Send a message to the client
		conn.WriteMessage(websocket.TextMessage, []byte(`{"op":"pong","req_id":"test_req_id"}`))
	}))
	defer server.Close()

	apiKey := "test_api_key"
	apiSecret := "test_api_secret"
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, true, true, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = "ws" + server.URL[len("http"):]

	err = client.Connect()
	assert.NoError(t, err)

	message, err := client.Receive()
	assert.NoError(t, err)
	assert.Equal(t, `{"op":"pong","req_id":"test_req_id"}`, string(message))
}

// TestClient_Close verifies the Close function closes the connection correctly.
func TestClient_Close(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		assert.NoError(t, err)
		defer conn.Close()
	}))
	defer server.Close()

	apiKey := "test_api_key"
	apiSecret := "test_api_secret"
	isTestNet := true
	isPublic := true
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, isTestNet, isPublic, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = "ws" + server.URL[len("http"):]

	err = client.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, client.Conn)

	client.Close()
	assert.True(t, client.isClosed)
}

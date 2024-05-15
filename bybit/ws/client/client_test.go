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
	maxActiveTime := "1m"

	client, err := NewClient(apiKey, apiSecret, true, true, maxActiveTime)
	assert.NoError(t, err)

	client.wsURL = "ws" + server.URL[len("http"):]

	err = client.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, client.Conn)

	client.Close()
	assert.True(t, client.isClosed)
}

func TestGenerateWsSignature_NilData(t *testing.T) {
	apiSecret := "test_api_secret"
	expectedSignature := ""

	signature := GenerateWsSignature(apiSecret, "")

	if signature != expectedSignature {
		t.Errorf("Expected signature to be empty, but got: %s", signature)
	}
}

func TestGenerateWsSignature_EmptyData(t *testing.T) {
	apiSecret := "test_api_secret"
	expectedSignature := ""

	signature := GenerateWsSignature(apiSecret, "")

	if signature != expectedSignature {
		t.Errorf("Expected signature to be empty, but got: %s", signature)
	}
}

func TestGenerateWsSignature_NonEmptyData(t *testing.T) {
	apiSecret := "test_api_secret"
	expectedSignature := "581ff2e8bfe2e6bed95c14bb8ed7e1921f49fa95d71404b6ea26251c0a3d8648"

	signature := GenerateWsSignature(apiSecret, "test_data")

	if signature != expectedSignature {
		t.Errorf("Expected signature to be: %s, but got: %s", expectedSignature, signature)
	}
}

func TestGenerateWsSignature_DifferentData(t *testing.T) {
	apiSecret := "test_api_secret"
	expectedSignature1 := "9386c982519cf138734ee1a5df3622dce5a0eb0cacdb5bfac723ec192148e26c"
	expectedSignature2 := "ea5d9db1068b4ae6208a7496975143f044aa4253860c1a2e50b5bc805c74b7f6"

	signature1 := GenerateWsSignature(apiSecret, "test_data1")
	signature2 := GenerateWsSignature(apiSecret, "test_data2")

	if signature1 != expectedSignature1 {
		t.Errorf("Expected signature for data1 to be: %s, but got: %s", expectedSignature1, signature1)
	}
	if signature2 != expectedSignature2 {
		t.Errorf("Expected signature for data2 to be: %s, but got: %s", expectedSignature2, signature2)
	}
}

func TestGenerateWsSignature_SameData(t *testing.T) {
	apiSecret := "test_api_secret"
	signature1 := GenerateWsSignature(apiSecret, "test_data")
	signature2 := GenerateWsSignature(apiSecret, "test_data")

	if signature1 != signature2 {
		t.Errorf("Expected signatures for the same data to be the same, but got: %s and %s", signature1, signature2)
	}
}

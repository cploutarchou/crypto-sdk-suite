package client

import (
	"encoding/json"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func mockWSServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			dummyResponse := ws.SuccessResponse{
				Success: true,
				RetMsg:  "Test message",
				Op:      "test_op",
				ConnId:  "test_id",
			}
			response, _ := json.Marshal(dummyResponse)

			if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
				log.Println(err)
				return
			}
		}
	})

	go http.ListenAndServe("localhost:8080", nil)
}

func TestClient(t *testing.T) {
	// Start the mock WebSocket server.
	mockWSServer()

	// Give some time for the server to start
	time.Sleep(1 * time.Second)

	t.Run("successful connection", func(t *testing.T) {
		client, err := NewClient("localhost:8080", Public, string(Spot))
		if err != nil {
			t.Fatalf("Failed to connect: %v", err)
		}
		client.Close()
	})

	t.Run("successful authentication", func(t *testing.T) {
		client, _ := NewClient("localhost:8080", Public, string(Spot))
		err := client.Authenticate("testAPI", "testExpires", "testSignature")
		if err != nil {
			t.Fatalf("Authentication failed: %v", err)
		}
		client.Close()
	})

	t.Run("failed connection", func(t *testing.T) {
		_, err := NewClient("stream-testnet-fake.bybit.com", Public, string(Spot))
		if err == nil {
			t.Fatal("Expected connection to fail")
		}
	})

	t.Run("read message", func(t *testing.T) {
		client, _ := NewClient("localhost:8080", Public, string(Spot))
		msg, err := client.ReadMessage()
		if err != nil {
			t.Fatalf("Failed to read message: %v", err)
		}
		if msg.Op != "test_op" {
			t.Fatalf("Unexpected message op: %v", msg.Op)
		}
		client.Close()
	})

	t.Run("connection close", func(t *testing.T) {
		client, _ := NewClient("localhost:8080", Public, string(Spot))
		client.Close()
		err := client.Authenticate("testAPI", "testExpires", "testSignature")
		if err == nil {
			t.Fatal("Expected error due to closed connection")
		}
	})
}

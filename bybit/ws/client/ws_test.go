package client

import (
	"encoding/json"

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
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error during connection upgrade:", err)
			return
		}
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			dummyResponse := SuccessResponse{
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

	go http.ListenAndServe("localhost:8080", nil) //nolint:errcheck
}
func TestClient(t *testing.T) {
	// Start the mock WebSocket server.
	mockWSServer()

	// Give some time for the server to start
	time.Sleep(1 * time.Second)

	t.Run("successful connection", func(t *testing.T) {
		client, err := NewWSClient("", "", true, true)
		if err != nil {
			t.Fatalf("Failed to connect: %v", err)
		}
		client.Close()
	})

	t.Run("successful authentication", func(t *testing.T) {
		client, _ := NewWSClient("", "", true, true)
		client.Close()
	})

	t.Run("failed connection", func(t *testing.T) {
		_, err := NewWSClient("", "", true, true)
		if err == nil {
			t.Fatal("Expected connection to fail")
		}
	})

	t.Run("connection close", func(t *testing.T) {
		client, _ := NewWSClient("", "", true, true)
		client.Close()
		err := client.Authenticate("testAPI", "testExpires", "testSignature")
		if err == nil {
			t.Fatal("Expected error due to closed connection")
		}
	})
}

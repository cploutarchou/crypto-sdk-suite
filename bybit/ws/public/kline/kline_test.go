package kline

import (
	"testing"
	"time"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

// TestSubscribe tests the Subscribe method of the klineImpl struct.
func TestSubscribe(t *testing.T) {
	cli, err := client.NewPublicClient(true)
	if err != nil {
		t.Fatalf("Failed to create public client: %v", err)
	}

	kl, err := New(cli)
	if err != nil {
		t.Fatalf("Failed to initialize kline service: %v", err)
	}

	err = kl.Subscribe([]string{"BTCUSDT", "SOLUSDT"}, "1", func(data Data) {
		t.Logf("Received kline update: %+v\n", data)
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to kline updates: %v", err)
	}

	go func() {
		for msg := range kl.GetMessagesChan() {
			t.Logf("Received raw message: %s", string(msg))
		}
	}()

	// Let the test run for a short while to capture some messages
	select {
	case <-time.After(60 * time.Second):
		t.Fatalf("Timed out waiting for kline updates")
	case msg := <-kl.GetMessagesChan():
		t.Log("Received kline update")
		t.Log(string(msg))
	}

	kl.Stop()
	kl.Close()
}

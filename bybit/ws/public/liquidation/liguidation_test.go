package liquidation

import (
	"testing"
	"time"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
)

// TestSubscribe tests the Subscribe method of the klineImpl struct.
func TestSubscribe(t *testing.T) {
	cli, err := client.NewPublicClient(true, "usdt_contract")
	if err != nil {
		t.Fatalf("Failed to create public client: %v", err)
	}

	kl := New(cli)

	err = kl.Subscribe([]string{"GALAUSDT"}, func(data Data) {
		t.Logf("Received liquidation update: %+v\n", data)
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to liquidation updates: %v", err)
	}

	go func() {
		for msg := range kl.GetMessagesChan() {
			t.Logf("Received raw message: %s", string(msg))
		}
	}()

	// Let the test run for a short while to capture some messages
	select {
	case <-time.After(120 * time.Second):
		t.Fatalf("Timed out waiting for liquidation updates")
	case msg := <-kl.GetMessagesChan():
		t.Log("Received liquidation update")
		t.Log(string(msg))
	}

	kl.Stop()
	kl.Close()
}

package lt_kline

import (
	"testing"
	"time"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	"github.com/stretchr/testify/assert"
)

// TestSubscribe tests the Subscribe method of the klineImpl struct.
func TestSubscribe(t *testing.T) {
	cli, err := client.NewPublicClient(true, "")
	if err != nil {
		t.Fatalf("Failed to create public client: %v", err)
	}

	ltKline := New(cli)

	err = ltKline.Subscribe("30", "BTC3SUSDT", func(response LTKlineResponse) {
		assert.Equal(t, "kline_lt.30.BTC3SUSDT", response.Topic)
		assert.Equal(t, "snapshot", response.Type)
		assert.Len(t, response.Data, 1)
		assert.Equal(t, "30", response.Data[0].Interval)
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to liquidation updates: %v", err)
	}

	go func() {
		for msg := range ltKline.GetMessagesChan() {
			t.Logf("Received raw message: %s", string(msg))
		}
	}()

	// Let the test run for a short while to capture some messages
	select {
	case <-time.After(120 * time.Second):
		t.Fatalf("Timed out waiting for liquidation updates")
	case msg := <-ltKline.GetMessagesChan():
		t.Log("Received liquidation update")
		t.Log(string(msg))
	}

	ltKline.Stop()
	ltKline.Close()
}

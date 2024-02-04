package main

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"
	"os"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures"
)

var b futures.Futures

var dsd = binance.New(
	os.Getenv("BINANCE_FUTURES_TESTNET_API_KEY"),
	os.Getenv("BINANCE_FUTURES_TESTNET_API_SECRET"),
	true)

func init() {
	b = dsd.Futures()

}

// testPing tests the ping endpoint.
func testPing() {
	data, err := b.Market().Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
func testGetExchangeInfo() {
	data, err := b.Market().GetExchangeInfo()
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(pretty))
}

func testGetServerTime() {
	data, err := b.Market().CheckServerTime()
	if err != nil {
		fmt.Println(err)
	}
	formatted := models.ServerTimeResponse{ServerTime: data}
	// marshal the data
	formattedTime := formatted.Format("2024-01-02 15:04:05")
	pretty, err := json.MarshalIndent(formattedTime, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

func MarketOrderBook() {
	data, err := b.Market().OrderBook("BTCUSDT", 100)
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

func RecentTradesList() {
	data, err := b.Market().RecentTradesList("BTCUSDT", 100)
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

func OldTradesLookup() {
	data, err := b.Market().OldTradesLookup("BTCUSDT", -1, -1)
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

func CompressedAggregateTradesList() {
	data, err := b.Market().CompressedAggregateTradesList("BTCUSDT", -1, -1, -1, -1)
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))

}
func KlineCandlestickData() {
	data, err := b.Market().KlineCandlestickData("BTCUSDT", "1m", -1, -1, -1)
	if err != nil {
		fmt.Println(err)
	}
	// marshal the data
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

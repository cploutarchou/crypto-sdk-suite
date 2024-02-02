package main

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures"
)

var b futures.Futures

var dsd = binance.New(
	"",
	"",
	true)

func init() {
	b = dsd.Futures()

}

// testPing tests the ping endpoint.
func testPing() {
	data, err := b.Generic().Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
func testGetExchangeInfo() {
	data, err := b.Generic().GetExchangeInfo()
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
	data, err := b.Generic().CheckServerTime()
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

package main

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/binance"
	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/models"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures"
)

var b futures.Futures

var dsd = binance.New("",
	"",
	true)

func init() {
	b = dsd.Futures()

}

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
	formated := models.ServerTimeResponse{ServerTime: data}
	// marshal the data
	formatedTime := formated.Format("2024-01-02 15:04:05")
	pretty, err := json.MarshalIndent(formatedTime, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pretty))
}

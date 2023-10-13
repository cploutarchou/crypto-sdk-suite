package main

import (
	"fmt"
	coinmarketcap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency"
	"os"
)

var client_ *coinmarketcap.Client
var cr cryptocurrency.Cryptocurrency

func init() {
	testnetKey := os.Getenv("COINMARKERCAP_TESTNET_API_KEY")
	client_ = coinmarketcap.NewClient(testnetKey, true)
	cr = cryptocurrency.New(client_)
}

func GetIDMap() {

	resp, err := cr.GetMapID(nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, coin := range resp {
		fmt.Printf("ID: %d, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}
}

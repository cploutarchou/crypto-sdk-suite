package main

import (
	coinmarketcap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal(err)
		return
	}

	for _, coin := range resp {
		logrus.Infof("ID: %d, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}
}

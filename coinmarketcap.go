package main

import (
	"os"

	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency/info"

	coinmarketcap "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/cryptocurrency"
	"github.com/sirupsen/logrus"
)

var client_ *coinmarketcap.Client
var cr cryptocurrency.Cryptocurrency

func init() {
	testnetKey := os.Getenv("COINMARKERCAP_TESTNET_API_KEY")
	client_ = coinmarketcap.NewClient(testnetKey, true)
	cr = cryptocurrency.New(client_)
}

func GetIDMap() {
	resp, err := cr.Map().GetID(nil)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	for _, coin := range resp {
		logrus.Infof("%+v\n", coin)
	}
}

func GetGainersAndLosers() {
	resp, err := cr.GainersAndLosers().FetchGainersLosers(nil)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	for _, coin := range resp {
		logrus.Infof("%+v\n", coin)
	}
}
func GetMetadata() {
	parameters := info.Params{Symbols: []string{"BTC", "ETH"}}
	resp, err := cr.Info().GetMetadata(&parameters)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	for _, coin := range resp {
		logrus.Infof("%+v\n", coin)
	}
}

package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/account"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

var bybitCli *client.Client
var acc account.Account

func init() {
	key := os.Getenv("BYBIT_TESTNET_KEY")
	secret := os.Getenv("BYBIT_TESTNET_SECRET")
	bybitCli = client.NewClient(key, secret, true)
	acc = account.New(bybitCli)
}

func printJSON(v interface{}) {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error marshaling to JSON:", err)
		return
	}
	log.Println(string(d))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleErrorWithPrint(data interface{}, err error) {
	if err != nil {
		handleError(err)
		return
	}
	printJSON(data)
}

func initBybit() {
	key := os.Getenv("BYBIT_TESTNET_KEY")
	secret := os.Getenv("BYBIT_TESTNET_SECRET")
	bybitCli = client.NewClient(key, secret, true)
	acc = account.New(bybitCli)

}

func getWalletBalance() (interface{}, error) {
	wallet := acc.Wallet()
	return wallet.GetContractWalletBalance("BTC")
}

func upgradeToUnified() (interface{}, error) {
	upgrade := acc.UpgradeToUnified()
	return upgrade.Upgrade()
}

func getBorrowHistory() (interface{}, error) {
	borrow := acc.Borrow()
	return borrow.GetHistory("BTC", 0, 0, 0, "")
}

func getCollateralCoin() (interface{}, error) {
	collateral := acc.Collateral()
	return collateral.GetInfo("BTC")
}

func setCollateralCoin() (interface{}, error) {
	collateral := acc.Collateral()
	return collateral.Set("BTC", "ON")
}

func getCoinGreeks() (interface{}, error) {
	coinGreeks := acc.CoinGreek()
	return coinGreeks.Get("BTC")
}
func getFeeRates() (interface{}, error) {
	feeRates := acc.FeeRates()
	return feeRates.GetFeeRate("taker", "BTCUSDT", "USDT")
}

func getInfo() (interface{}, error) {
	info := acc.Info()
	return info.Get()
}

func getTransactionLog() (interface{}, error) {
	params := map[string]string{
		"accountType": "UNIFIED",
		"category":    "linear",
		"currency":    "USDT",
	}
	transactionLog := acc.TransactionLog()
	return transactionLog.Get(params)
}

func setMargin() (interface{}, error) {
	margin := acc.Margin()
	return margin.SetMarginMode("ISOLATED")
}

func runAccountExamples() {
	handleErrorWithPrint(getWalletBalance())
	handleErrorWithPrint(upgradeToUnified())
	handleErrorWithPrint(getBorrowHistory())
	handleErrorWithPrint(getCollateralCoin())
	handleErrorWithPrint(setCollateralCoin())
	handleErrorWithPrint(getCoinGreeks())
	handleErrorWithPrint(getFeeRates())
	handleErrorWithPrint(getInfo())
	handleErrorWithPrint(getTransactionLog())
	handleErrorWithPrint(setMargin())
}
func bybitExamples() {
	initBybit()
	runAccountExamples()

}

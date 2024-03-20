package main

import (
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/account"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws"
	wsClient "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	kline2 "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/kline"
	ticker2 "github.com/cploutarchou/crypto-sdk-suite/bybit/ws/public/ticker"
	"log"
	"os"
	"strconv"
)

var bybitCli *client.Client
var acc account.Account
var websocket ws.WebSocket
var key string
var secret string

func init() {
	key = os.Getenv("BYBIT_FUTURES_TESTNET_API_KEY")
	secret = os.Getenv("k2MbmSgKbnVnc4xxY7FzmADq2GjoYWZTJZcM")
	bybitCli = client.NewClient(key, secret, true)
	acc = account.New(bybitCli)
}

func getWalletBalance() (interface{}, error) {
	wallet := acc.Wallet()
	fmt.Println("getWalletBalance")
	return wallet.GetContractWalletBalance("BTC")
}

func upgradeToUnified() (interface{}, error) {
	fmt.Println("upgradeToUnified")
	upgrade := acc.UpgradeToUnified()
	return upgrade.Upgrade()
}

func getBorrowHistory() (interface{}, error) {
	fmt.Println("getBorrowHistory")
	borrow := acc.Borrow()
	return borrow.GetHistory("BTC", 0, 0, 0, "")
}

func getCollateralCoin() (interface{}, error) {
	fmt.Println("getCollateralCoin")
	collateral := acc.Collateral()
	return collateral.GetInfo("BTCUSDT")
}

func setCollateralCoin() (interface{}, error) {
	fmt.Println("setCollateralCoin")
	collateral := acc.Collateral()
	return collateral.Set("BTC", "ON")
}

func getCoinGreeks() (interface{}, error) {
	fmt.Println("getCoinGreeks")
	coinGreeks := acc.CoinGreek()
	return coinGreeks.Get("BTC")
}
func getFeeRates() (interface{}, error) {
	fmt.Println("getFeeRates")
	feeRates := acc.FeeRates()
	return feeRates.GetFeeRate("taker", "BTCUSDT", "USDT")
}

func getInfo() (interface{}, error) {
	fmt.Println("getInfo")
	info := acc.Info()
	return info.Get()
}

func getTransactionLog() (interface{}, error) {
	params := map[string]string{
		"accountType": "UNIFIED",
		"category":    "linear",
		"currency":    "USDT",
	}
	fmt.Println("getTransactionLog")
	transactionLog := acc.TransactionLog()
	return transactionLog.Get(params)
}

func setMargin() (interface{}, error) {
	margin := acc.Margin()
	fmt.Println("setMargin")
	return margin.SetMarginMode("ISOLATED")
}

func setMMP() (interface{}, error) {
	margin := acc.Margin()
	params := &account.MMPParams{
		BaseCoin:     "BTC",
		Window:       200,
		FrozenPeriod: 10,
		QtyLimit:     100,
		DeltaLimit:   100,
	}
	fmt.Println("setMMP")
	return margin.SetMMP(params)
}

func resetMMP() (interface{}, error) {
	margin := acc.Margin()
	fmt.Println("resetMMP")
	return margin.ResetMMP("BTC")
}

func getMMPState() (interface{}, error) {
	margin := acc.Margin()
	fmt.Println("getMMPState")
	return margin.GetMMPState("BTC")
}

func wsConnectTicker() {
	// Correctly initialized a buffered channel of float64 values
	b := make(chan float64, 1)
	fmt.Println("wsConnect")
	client_, err := wsClient.NewClient(key, secret, true)
	if err != nil {
		log.Printf("ERROR: Failed to create WebSocket client: %v", err)
		return
	}

	websocket = ws.New(client_, "linear")
	publicWS, err := websocket.Public()
	if err != nil {
		log.Printf("ERROR: Failed to access public WebSocket endpoint: %v", err)
		return
	}

	ticker := publicWS.Ticker()

	// Subscribe to ticker updates
	err = ticker.Subscribe("BTCUSDT", func(data ticker2.Data) {
		if data.LastPrice != "" {
			lastPrice, parseErr := strconv.ParseFloat(data.LastPrice, 64)
			if parseErr != nil {
				log.Printf("ERROR: Parsing last price: %v", parseErr)
				return
			}
			// Use the local channel 'b' directly for sending lastPrice
			select {
			case b <- lastPrice: // Corrected to use 'b' directly
			default:
				log.Println("WARNING: Channel is blocked or full, skipping update.")
			}
		}
	})
	if err != nil {
		log.Printf("ERROR: Failed to subscribe to ticker updates: %v", err)
		return
	}

	go ticker.Listen()

	log.Println("INFO: Successfully subscribed to live price updates for BTCUSDT")

	for price := range b {
		log.Printf("Received price update: %f", price)
	}
}
func wsConnectKline() {
	fmt.Println("wsConnectKline")

	client_, err := wsClient.NewClient(key, secret, true)
	if err != nil {
		log.Printf("ERROR: Failed to create WebSocket client: %v", err)
		return
	}

	klineService, err := kline2.New(client_) // Adjust based on your New function
	if err != nil {
		log.Printf("ERROR: Failed to initialize kline service: %v", err)
		return
	}

	err = klineService.Subscribe([]string{"BTCUSDT", "SOLUSDT"}, "1", func(data kline2.Data) {})

	if err != nil {
		log.Printf("ERROR: Failed to subscribe to kline updates: %v", err)
		return
	}

	log.Println("INFO: Successfully subscribed to kline updates for BTCUSDT and THETAUSDT")

	for kline := range klineService.GetMessagesChan() {
		log.Printf("Received kline update: %+v\n", string(kline))
	}
}

func runAccountExamples() {
	//handleErrorWithPrint(getWalletBalance())
	//handleErrorWithPrint(upgradeToUnified())
	//handleErrorWithPrint(getBorrowHistory())
	////handleErrorWithPrint(getCollateralCoin())
	//handleErrorWithPrint(setCollateralCoin())
	//handleErrorWithPrint(getCoinGreeks())
	//handleErrorWithPrint(getFeeRates())
	//handleErrorWithPrint(getInfo())
	//handleErrorWithPrint(getTransactionLog())
	//handleErrorWithPrint(setMargin())
	//handleErrorWithPrint(setMMP())
	//handleErrorWithPrint(resetMMP())
	//handleErrorWithPrint(getMMPState())
	//wsConnectTicker()
	wsConnectKline()
}
func bybitExamples() {
	runAccountExamples()

}

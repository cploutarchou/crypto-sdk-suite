package constants

// Base URLs
const (
	ProductionBaseURL = "https://fapi.binance.com"
	TestnetBaseURL    = "https://testnet.binancefuture.com"
	ProductionWSURL   = "wss://fstream.binancefuture.com"
	TestnetWSURL      = "wss://fstream.testnet.binancefuture.com"
)

// API Endpoints
const (
	PingEndpoint         = "/fapi/v1/ping"
	ServerTimeEndpoint   = "/fapi/v1/time"
	ExchangeInfoEndpoint = "/fapi/v1/exchangeInfo"
	OrderBookEndpoint    = "/fapi/v1/depth"
	RecentTradesEndpoint = "/fapi/v1/trades"
)

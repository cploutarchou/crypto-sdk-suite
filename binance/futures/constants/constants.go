// Package constants defines various constants used for Binance Futures API.
package constants

// Base URLs for Binance Futures API.
const (
	// ProductionBaseURL is the base URL for the Binance Futures production environment.
	ProductionBaseURL = "https://fapi.binance.com"

	// TestnetBaseURL is the base URL for the Binance Futures testnet environment.
	TestnetBaseURL = "https://testnet.binancefuture.com"
)

// WebSocket URLs for Binance Futures API.
const (
	// ProductionWSURL is the WebSocket URL for the Binance Futures production environment.
	ProductionWSURL = "wss://fstream.binancefuture.com"

	// TestnetWSURL is the WebSocket URL for the Binance Futures testnet environment.
	TestnetWSURL = "wss://fstream.testnet.binancefuture.com"
)

// API Endpoints for Binance Futures.
const (
	// PingEndpoint is the endpoint for server ping.
	PingEndpoint = "/fapi/v1/ping"

	// ServerTimeEndpoint is the endpoint to get the server time.
	ServerTimeEndpoint = "/fapi/v1/time"

	// ExchangeInfoEndpoint is the endpoint to get exchange information.
	ExchangeInfoEndpoint = "/fapi/v1/exchangeInfo"

	// OrderBookEndpoint is the endpoint to get order book data.
	OrderBookEndpoint = "/fapi/v1/depth"

	// RecentTradesEndpoint is the endpoint to get recent trade information.
	RecentTradesEndpoint = "/fapi/v1/trades"
)

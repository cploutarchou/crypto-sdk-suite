package client

import "time"

const (
	// BaseURL is the base URL for the Bybit API
	BaseURL = "https://api.bybit.com"

	// TestnetBaseURL is the base URL for the Bybit Testnet API
	TestnetBaseURL = "https://api-testnet.bybit.com"

	// WsBaseURL is the base URL for the Bybit Websocket API
	WsBaseURL = "wss://stream.bybit.com/realtime"

	// TestnetWsBaseURL is the base URL for the Bybit Testnet Websocket API
	TestnetWsBaseURL = "wss://stream-testnet.bybit.com/realtime"

	// ApiVersion is the version of the Bybit API
	ApiVersion = "v5"

	// TestnetApiVersion is the version of the Bybit Testnet API
	TestnetApiVersion = "v5"

	// UserAgent is the default user agent for the Bybit API
	UserAgent = "go-bybit"

	// TestnetUserAgent is the default user agent for the Bybit Testnet API
	TestnetUserAgent = "go-bybit-testnet"

	// DefaultHTTPTimeout is the default HTTP timeout for the Bybit API
	DefaultHTTPTimeout = 30 * time.Second
)

package client

const (
	BaseURLSpotMargin = "https://api.kucoin.com/api/"

	BaseURLFuture = "https://api-futures.kucoin.com/api/"

	// ApiVersion is the version of the Bybit API
	ApiVersion = "v2"
)

type MarketType string

const (
	SpotMargin MarketType = "spot-margin"
	Future     MarketType = "future"
)

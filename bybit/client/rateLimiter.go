package client

import (
	"golang.org/x/time/rate"
)

const (
	tenPerMinute        = 10
	onePerMinute        = 1
	twentyPerMinute     = 20
	fivePerMinute       = 5
	oneHoundedPerMinute = 100
)

var endpointLimits = map[string]rate.Limit{
	// Orders
	"POST /v5/order/create":     rate.Limit(tenPerMinute),
	"POST /v5/order/amend":      rate.Limit(tenPerMinute),
	"POST /v5/order/cancel":     rate.Limit(tenPerMinute),
	"POST /v5/order/cancel-all": rate.Limit(tenPerMinute),
	"GET /v5/order/realtime":    rate.Limit(tenPerMinute),
	"GET /v5/order/history":     rate.Limit(tenPerMinute),
	"GET /v5/execution/list":    rate.Limit(tenPerMinute),

	// Position
	"GET /v5/position/list":          rate.Limit(tenPerMinute),
	"GET /v5/position/closed-pnl":    rate.Limit(tenPerMinute),
	"POST /v5/position/set-leverage": rate.Limit(tenPerMinute),

	// Account
	"GET /v5/account/wallet-balance": rate.Limit(twentyPerMinute),
	"GET /v5/account/fee-rate":       rate.Limit(tenPerMinute),

	// Asset
	"GET /v5/asset/transfer/query-asset-info":              rate.Limit(onePerMinute), // Corrected for 60 req/min
	"GET /v5/asset/transfer/query-transfer-coin-list":      rate.Limit(onePerMinute), // Corrected for 60 req/min
	"GET /v5/asset/transfer/query-inter-transfer-list":     rate.Limit(onePerMinute), // Corrected for 60 req/min
	"GET /v5/asset/transfer/query-sub-member-list":         rate.Limit(onePerMinute), // Corrected for 60 req/min
	"GET /v5/asset/transfer/query-universal-transfer-list": rate.Limit(fivePerMinute),
	"GET /v5/asset/transfer/query-account-coins-balance":   rate.Limit(fivePerMinute),
	"GET /v5/asset/deposit/query-record":                   rate.Limit(oneHoundedPerMinute),
	"GET /v5/asset/deposit/query-sub-member-record":        rate.Limit(300),
	"GET /v5/asset/deposit/query-address":                  rate.Limit(300),
	"GET /v5/asset/deposit/query-sub-member-address":       rate.Limit(300),
	"GET /v5/asset/withdraw/query-record":                  rate.Limit(300),
	"GET /v5/asset/coin/query-info":                        rate.Limit(2),
	"GET /v5/asset/exchange/order-record":                  rate.Limit(600),
	"POST /v5/asset/transfer/inter-transfer":               rate.Limit(3), // Corrected for 20 req/min
	"POST /v5/asset/transfer/save-transfer-sub-member":     rate.Limit(twentyPerMinute),
	"POST /v5/asset/transfer/universal-transfer":           rate.Limit(fivePerMinute),
	"POST /v5/asset/withdraw/create":                       rate.Limit(onePerMinute),
	"POST /v5/asset/withdraw/cancel":                       rate.Limit(onePerMinute), // Corrected for 60 req/min

	// User
	"POST /v5/user/create-sub-member": rate.Limit(fivePerMinute),
	"POST /v5/user/create-sub-api":    rate.Limit(fivePerMinute),
	"POST /v5/user/frozen-sub-member": rate.Limit(fivePerMinute),
	"POST /v5/user/update-api":        rate.Limit(fivePerMinute),
	"POST /v5/user/update-sub-api":    rate.Limit(fivePerMinute),
	"POST /v5/user/delete-api":        rate.Limit(fivePerMinute),
	"POST /v5/user/delete-sub-api":    rate.Limit(fivePerMinute),
	"GET /v5/user/query-sub-members":  rate.Limit(tenPerMinute),
	"GET /v5/user/query-api":          rate.Limit(tenPerMinute),

	// Spot Leverage Token
	"GET /v5/spot-lever-token/order-record": rate.Limit(50),
	"POST /v5/spot-lever-token/purchase":    rate.Limit(twentyPerMinute),
	"POST /v5/spot-lever-token/redeem":      rate.Limit(twentyPerMinute),

	// Spot Margin Trade (Classic)
	"GET /v5/spot-cross-margin-trade/loan-info":     rate.Limit(50),
	"GET /v5/spot-cross-margin-trade/account":       rate.Limit(50),
	"GET /v5/spot-cross-margin-trade/orders":        rate.Limit(50),
	"GET /v5/spot-cross-margin-trade/repay-history": rate.Limit(50),
	"POST /v5/spot-cross-margin-trade/loan":         rate.Limit(twentyPerMinute),
	"POST /v5/spot-cross-margin-trade/repay":        rate.Limit(twentyPerMinute),
	"POST /v5/spot-cross-margin-trade/switch":       rate.Limit(twentyPerMinute),
}

type EndpointRateLimiter struct {
	limiters map[string]*rate.Limiter
}

func NewEndpointRateLimiter() *EndpointRateLimiter {
	return &EndpointRateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

// SetLimiter updates or creates a rate limiter for a specific endpoint
func (e *EndpointRateLimiter) SetLimiter(endpointKey string, limiter *rate.Limiter) {
	e.limiters[endpointKey] = limiter
}

// GetLimiter retrieves an existing rate limiter for an endpoint, returning nil if not found
func (e *EndpointRateLimiter) GetLimiter(endpointKey string) *rate.Limiter {
	if limiter, ok := e.limiters[endpointKey]; ok {
		return limiter
	}

	// Set default rate limiter to 30 requests per minute
	defaultRate := rate.Limit(30.0 / 60.0) // 30 requests per minute (30/60 = 0.5 requests per second)
	return rate.NewLimiter(defaultRate, 1) // Burst size of 1
}

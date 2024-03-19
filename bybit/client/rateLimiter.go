package client

import (
	"golang.org/x/time/rate"
	"time"
)

var endpointLimits = map[string]rate.Limit{
	// Orders
	"POST /v5/order/create":     rate.Limit(10 / time.Second),
	"POST /v5/order/amend":      rate.Limit(10 / time.Second),
	"POST /v5/order/cancel":     rate.Limit(10 / time.Second),
	"POST /v5/order/cancel-all": rate.Limit(10 / time.Second),
	"GET /v5/order/realtime":    rate.Limit(10 / time.Second),
	"GET /v5/order/history":     rate.Limit(10 / time.Second),
	"GET /v5/execution/list":    rate.Limit(10 / time.Second),

	// Position
	"GET /v5/position/list":          rate.Limit(10 / time.Second),
	"GET /v5/position/closed-pnl":    rate.Limit(10 / time.Second),
	"POST /v5/position/set-leverage": rate.Limit(10 / time.Second),

	// Account
	"GET /v5/account/wallet-balance": rate.Limit(20 / time.Second), // Assuming SPOT type for simplicity
	"GET /v5/account/fee-rate":       rate.Limit(10 / time.Second), // Assuming linear category for simplicity

	// Asset
	"GET /v5/asset/transfer/query-asset-info":              rate.Limit(1), // 60 req/min
	"GET /v5/asset/transfer/query-transfer-coin-list":      rate.Limit(1), // 60 req/min
	"GET /v5/asset/transfer/query-inter-transfer-list":     rate.Limit(1), // 60 req/min
	"GET /v5/asset/transfer/query-sub-member-list":         rate.Limit(1), // 60 req/min
	"GET /v5/asset/transfer/query-universal-transfer-list": rate.Limit(5 / time.Second),
	"GET /v5/asset/transfer/query-account-coins-balance":   rate.Limit(5 / time.Second),
	"GET /v5/asset/deposit/query-record":                   rate.Limit(100 / time.Second),
	"GET /v5/asset/deposit/query-sub-member-record":        rate.Limit(300 / time.Second),
	"GET /v5/asset/deposit/query-address":                  rate.Limit(300 / time.Second),
	"GET /v5/asset/deposit/query-sub-member-address":       rate.Limit(300 / time.Second),
	"GET /v5/asset/withdraw/query-record":                  rate.Limit(300 / time.Second),
	"GET /v5/asset/coin/query-info":                        rate.Limit(2 / time.Second),
	"GET /v5/asset/exchange/order-record":                  rate.Limit(600 / time.Second),
	"POST /v5/asset/transfer/inter-transfer":               rate.Limit(1 / 3 * time.Second), // 20 req/min
	"POST /v5/asset/transfer/save-transfer-sub-member":     rate.Limit(20 / time.Second),
	"POST /v5/asset/transfer/universal-transfer":           rate.Limit(5 / time.Second),
	"POST /v5/asset/withdraw/create":                       rate.Limit(1 / time.Second),
	"POST /v5/asset/withdraw/cancel":                       rate.Limit(1), // 60 req/min

	// User
	"POST /v5/user/create-sub-member": rate.Limit(5 / time.Second),
	"POST /v5/user/create-sub-api":    rate.Limit(5 / time.Second),
	"POST /v5/user/frozen-sub-member": rate.Limit(5 / time.Second),
	"POST /v5/user/update-api":        rate.Limit(5 / time.Second),
	"POST /v5/user/update-sub-api":    rate.Limit(5 / time.Second),
	"POST /v5/user/delete-api":        rate.Limit(5 / time.Second),
	"POST /v5/user/delete-sub-api":    rate.Limit(5 / time.Second),
	"GET /v5/user/query-sub-members":  rate.Limit(10 / time.Second),
	"GET /v5/user/query-api":          rate.Limit(10 / time.Second),

	// Spot Leverage Token
	"GET /v5/spot-lever-token/order-record": rate.Limit(50 / time.Second),
	"POST /v5/spot-lever-token/purchase":    rate.Limit(20 / time.Second),
	"POST /v5/spot-lever-token/redeem":      rate.Limit(20 / time.Second),

	// Spot Margin Trade (Classic)
	"GET /v5/spot-cross-margin-trade/loan-info":     rate.Limit(50 / time.Second),
	"GET /v5/spot-cross-margin-trade/account":       rate.Limit(50 / time.Second),
	"GET /v5/spot-cross-margin-trade/orders":        rate.Limit(50 / time.Second),
	"GET /v5/spot-cross-margin-trade/repay-history": rate.Limit(50 / time.Second),
	"POST /v5/spot-cross-margin-trade/loan":         rate.Limit(20 / time.Second),
	"POST /v5/spot-cross-margin-trade/repay":        rate.Limit(20 / time.Second),
	"POST /v5/spot-cross-margin-trade/switch":       rate.Limit(20 / time.Second),
}

type EndpointRateLimiter struct {
	limiters map[string]*rate.Limiter
}

func NewEndpointRateLimiter() *EndpointRateLimiter {
	return &EndpointRateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

// GetLimiter fetches or creates a new limiter for the given endpointKey.
// GetLimiter retrieves an existing limiter for the endpoint. Returns nil if not found.
func (e *EndpointRateLimiter) GetLimiter(endpointKey string) *rate.Limiter {
	return e.limiters[endpointKey]
}

// SetLimiter sets a new limiter for the endpoint.
func (e *EndpointRateLimiter) SetLimiter(endpointKey string, limiter *rate.Limiter) {
	e.limiters[endpointKey] = limiter
}

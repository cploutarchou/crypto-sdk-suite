package account

type AccountCategory int
type AccountType string
type TimeInterval int

const (
	Unified  AccountType = "UNIFIED"
	Contract AccountType = "CONTRACT"
	Spot     AccountType = "SPOT"
)

type CollateralSwitch string

// AccountCategory constants using iota
const (
	ON  CollateralSwitch = "ON"
	OFF CollateralSwitch = "OFF"
)

type EndpointsStruct struct {
	Borrow           string
	CoinGreek        string
	Collateral       string
	UpgradeToUnified string
	Wallet           string
}

var Endpoints = EndpointsStruct{
	Borrow:           "/v5/account/borrow-history",
	CoinGreek:        "/v5/asset/coin-greeks",
	Collateral:       "/v5/account/set-collateral-switch",
	UpgradeToUnified: "/v5/account/upgrade-to-uta",
	Wallet:           "/v5/account/wallet-balance",
}

func (a AccountCategory) String() string {
	names := [...]string{
		"UnifiedSpot",
		"UnifiedLinearUSDT",
		"UnifiedLinearUSDC",
		"UnifiedLinearUSDCPerp",
		"UnifiedLinearUSDCFutures",
		"UnifiedInverse",
		"UnifiedInversePerp",
		"UnifiedInverseFutures",
		"UnifiedOption",
		"NormalLinearUSDTPerp",
		"NormalInverse",
		"NormalInversePerp",
		"NormalInverseFutures",
		"NormalSpot",
	}
	return names[a]
}

func (ti TimeInterval) String() string {
	intervals := [...]string{
		"1 Minute",
		"3 Minutes",
		"5 Minutes",
		"15 Minutes",
		"30 Minutes",
		"60 Minutes",
		"120 Minutes",
		"240 Minutes",
		"360 Minutes",
		"720 Minutes",
		"Day",
		"Week",
		"Month",
	}
	return intervals[ti-1] // subtract 1 because the iota starts from 1 in this case
}

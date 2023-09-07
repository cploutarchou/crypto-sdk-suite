package account

type AccountCategory int

type TimeInterval int

const (
	Minute1 TimeInterval = iota + 1
	Minute3
	Minute5
	Minute15
	Minute30
	Minute60
	Minute120
	Minute240
	Minute360
	Minute720
	Day
	Week
	Month

	// Unified Account
	UnifiedSpot AccountCategory = iota
	UnifiedLinearUSDT
	UnifiedLinearUSDC
	UnifiedLinearUSDCPerp
	UnifiedLinearUSDCFutures
	UnifiedInverse
	UnifiedInversePerp
	UnifiedInverseFutures
	UnifiedOption

	// Normal Account
	NormalLinearUSDTPerp
	NormalInverse
	NormalInversePerp
	NormalInverseFutures
	NormalSpot
)

func (a AccountCategory) String() string {
	return [...]string{
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
	}[a]
}

func (ti TimeInterval) String() string {
	return [...]string{
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
	}[ti-1] // subtract 1 because the iota starts from 1 in this case
}

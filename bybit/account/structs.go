package account

// BaseResponse holds fields that are common for many API responses.
type BaseResponse struct {
	RetCode    int    `json:"retCode"`
	RetMsg     string `json:"retMsg"`
	Time       int64  `json:"time"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
}

// CoinDetails represents individual coin details within an account.
type CoinDetails struct {
	AvailableToBorrow   string `json:"availableToBorrow"`
	Bonus               string `json:"bonus"`
	AccruedInterest     string `json:"accruedInterest"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	TotalOrderIM        string `json:"totalOrderIM"`
	Equity              string `json:"equity"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UsdValue            string `json:"usdValue"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	BorrowAmount        string `json:"borrowAmount"`
	TotalPositionIM     string `json:"totalPositionIM"`
	WalletBalance       string `json:"walletBalance"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Locked              string `json:"locked"`
	MarginCollateral    bool   `json:"marginCollateral"`
	Coin                string `json:"coin"`
}

// AccountDetails holds details related to an account's wallet balance.
type AccountDetails struct {
	TotalEquity            string        `json:"totalEquity"`
	AccountIMRate          string        `json:"accountIMRate"`
	TotalMarginBalance     string        `json:"totalMarginBalance"`
	TotalInitialMargin     string        `json:"totalInitialMargin"`
	AccountType            string        `json:"accountType"`
	TotalAvailableBalance  string        `json:"totalAvailableBalance"`
	AccountMMRate          string        `json:"accountMMRate"`
	TotalPerpUPL           string        `json:"totalPerpUPL"`
	TotalWalletBalance     string        `json:"totalWalletBalance"`
	AccountLTV             string        `json:"accountLTV"`
	TotalMaintenanceMargin string        `json:"totalMaintenanceMargin"`
	Coin                   []CoinDetails `json:"coin"`
}

// WalletBalance represents the wallet balance response.
type WalletBalance struct {
	BaseResponse
	Result struct {
		List []AccountDetails `json:"list"`
	} `json:"result"`
}

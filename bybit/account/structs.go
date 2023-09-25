package account

// BaseResponse is a generic struct used to parse the common response received from Bybit API
type BaseResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Time       int64                  `json:"time"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
}

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

type AccDetails struct {
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

// AccInfo represents the response from the /v5/account/info endpoint.
type AccInfo struct {
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	MarginMode          string `json:"marginMode"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	UpdatedTime         string `json:"updatedTime"`
}
type WalletBalance struct {
	BaseResponse
	Result struct {
		List []AccDetails `json:"list"`
	}
}

type BorrowItem struct {
	CreatedTime               int64  `json:"createdTime"`
	CostExemption             string `json:"costExemption"`
	InterestBearingBorrowSize string `json:"interestBearingBorrowSize"`
	Currency                  string `json:"currency"`
	HourlyBorrowRate          string `json:"hourlyBorrowRate"`
	BorrowCost                string `json:"borrowCost"`
}

type BorrowRes struct {
	BaseResponse
	Result struct {
		NextPageCursor string       `json:"nextPageCursor"`
		List           []BorrowItem `json:"list"`
	}
}

type CoinGreekItem struct {
	BaseCoin   string `json:"baseCoin"`
	TotalDelta string `json:"totalDelta"`
	TotalGamma string `json:"totalGamma"`
	TotalVega  string `json:"totalVega"`
	TotalTheta string `json:"totalTheta"`
}

type CoinGreekRes struct {
	BaseResponse
	Result struct {
		List []CoinGreekItem `json:"list"`
	}
}

type UnifiedUpdateMsg struct {
	Msg []string `json:"msg"`
}

type UpgradeToUnifiedResponse struct {
	BaseResponse
	Result struct {
		UnifiedUpdateStatus string           `json:"unifiedUpdateStatus"`
		UnifiedUpdateMsg    UnifiedUpdateMsg `json:"unifiedUpdateMsg"`
	}
}

type CollateralData struct {
	CollateralSwitch    bool   `json:"collateralSwitch"`
	BorrowAmount        string `json:"borrowAmount"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	FreeBorrowingAmount string `json:"freeBorrowingAmount"`
	Borrowable          bool   `json:"borrowable"`
	Currency            string `json:"currency"`
	MaxBorrowingAmount  string `json:"maxBorrowingAmount"`
	HourlyBorrowRate    string `json:"hourlyBorrowRate"`
	BorrowUsageRate     string `json:"borrowUsageRate"`
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralRatio     string `json:"collateralRatio"`
}

type CollateralInfoResponse struct {
	BaseResponse
	Result CollateralResult
}

type CollateralResult struct {
	List []CollateralData `json:"list"`
}

type FeeRate struct {
	Symbol       string `json:"symbol"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

type FeeRatesResponse struct {
	BaseResponse
	Result struct {
		List []FeeRate
	}
}

type SetMarginModeResponse struct {
	BaseResponse
	Result struct {
		Reasons []struct {
			ReasonCode string `json:"reasonCode"`
			ReasonMsg  string `json:"reasonMsg"`
		} `json:"reasons"`
	} `json:"result"`
}

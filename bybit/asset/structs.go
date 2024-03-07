package asset

// GetCoinExchangeRecordsRequest represents the query parameters for fetching coin exchange records.
type GetCoinExchangeRecordsRequest struct {
	FromCoin *string `json:"fromCoin,omitempty"` // Optional: The currency to convert from
	ToCoin   *string `json:"toCoin,omitempty"`   // Optional: The currency to convert to
	Limit    *int    `json:"limit,omitempty"`    // Optional: Limit for data size per page
	Cursor   *string `json:"cursor,omitempty"`   // Optional: Cursor for pagination
}

// CoinExchangeRecord represents a single record of a coin exchange.
type CoinExchangeRecord struct {
	FromCoin     string `json:"fromCoin"`
	FromAmount   string `json:"fromAmount"`
	ToCoin       string `json:"toCoin"`
	ToAmount     string `json:"toAmount"`
	ExchangeRate string `json:"exchangeRate"`
	CreatedTime  string `json:"createdTime"`
	ExchangeTxId string `json:"exchangeTxId"`
}

// GetCoinExchangeRecordsResponse represents the response from fetching coin exchange records.
type GetCoinExchangeRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderBody      []CoinExchangeRecord `json:"orderBody"`
		NextPageCursor string               `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetDeliveryRecordRequest represents the query parameters for fetching delivery records.
type GetDeliveryRecordRequest struct {
	Category  string  `json:"category"`            // Required: Product type. option, linear
	Symbol    *string `json:"symbol,omitempty"`    // Optional: Symbol name
	StartTime *int64  `json:"startTime,omitempty"` // Optional: Start timestamp (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: End time. timestamp (ms)
	ExpDate   *string `json:"expDate,omitempty"`   // Optional: Expiry date. 25MAR22
	Limit     *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Cursor for pagination
}

// DeliveryRecordEntry represents a single entry in the delivery record list.
type DeliveryRecordEntry struct {
	DeliveryTime  int64  `json:"deliveryTime"`  // Delivery time (ms)
	Symbol        string `json:"symbol"`        // Symbol name
	Side          string `json:"side"`          // Buy, Sell
	Position      string `json:"position"`      // Executed size
	DeliveryPrice string `json:"deliveryPrice"` // Delivery price
	Strike        string `json:"strike"`        // Exercise price
	Fee           string `json:"fee"`           // Trading fee
	DeliveryRpl   string `json:"deliveryRpl"`   // Realized PnL of the delivery
}

// GetDeliveryRecordResponse represents the response from fetching delivery records.
type GetDeliveryRecordResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		NextPageCursor string                `json:"nextPageCursor"`
		Category       string                `json:"category"`
		List           []DeliveryRecordEntry `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetSessionSettlementRecordRequest represents the query parameters for fetching session settlement records.
type GetSessionSettlementRecordRequest struct {
	Category  string  `json:"category"`            // Required: Product type, e.g., "linear"
	Symbol    *string `json:"symbol,omitempty"`    // Optional: Symbol name
	StartTime *int64  `json:"startTime,omitempty"` // Optional: Start timestamp (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: End time (ms)
	Limit     *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Cursor for pagination
}

// SessionSettlementRecord represents a single entry in the session settlement record list.
type SessionSettlementRecord struct {
	Symbol          string `json:"symbol"`          // Symbol name
	Side            string `json:"side"`            // Buy or Sell
	Size            string `json:"size"`            // Position size
	SessionAvgPrice string `json:"sessionAvgPrice"` // Settlement price
	MarkPrice       string `json:"markPrice"`       // Mark price
	RealisedPnl     string `json:"realisedPnl"`     // Realised PnL
	CreatedTime     string `json:"createdTime"`     // Created time (ms)
}

// GetSessionSettlementRecordResponse represents the response from fetching session settlement records.
type GetSessionSettlementRecordResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		NextPageCursor string                    `json:"nextPageCursor"`
		Category       string                    `json:"category"`
		List           []SessionSettlementRecord `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetAssetInfoRequest represents the query parameters for fetching asset information.
type GetAssetInfoRequest struct {
	AccountType string  `json:"accountType"`    // Required: Account type, e.g., "SPOT"
	Coin        *string `json:"coin,omitempty"` // Optional: Coin name
}

// AssetInfoEntry represents a single asset entry within the asset information list.
type AssetInfoEntry struct {
	Coin     string `json:"coin"`     // Coin
	Frozen   string `json:"frozen"`   // Freeze amount
	Free     string `json:"free"`     // Free balance
	Withdraw string `json:"withdraw"` // Amount in withdrawing
}

// GetAssetInfoResponse represents the response from fetching asset information.
type GetAssetInfoResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Spot struct {
			Status string           `json:"status"` // account status
			Assets []AssetInfoEntry `json:"assets"` // Assets array
		} `json:"spot"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetAllCoinsBalanceRequest represents the query parameters for fetching all coins' balances.
type GetAllCoinsBalanceRequest struct {
	MemberID    *string `json:"memberId,omitempty"`  // Optional: User Id, required for checking sub account coin balance with master API key
	AccountType string  `json:"accountType"`         // Required: Account type
	Coin        *string `json:"coin,omitempty"`      // Optional: Coin name(s), multiple coins separated by comma
	WithBonus   *int    `json:"withBonus,omitempty"` // Optional: 0(default): not query bonus. 1: query bonus
}

// CoinBalanceEntry represents a single coin's balance information.
type CoinBalanceEntry struct {
	Coin            string `json:"coin"`            // Currency type
	WalletBalance   string `json:"walletBalance"`   // Wallet balance
	TransferBalance string `json:"transferBalance"` // Transferable balance
	Bonus           string `json:"bonus,omitempty"` // The bonus (if queried)
}

// GetAllCoinsBalanceResponse represents the response from fetching all coins' balances.
type GetAllCoinsBalanceResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		MemberID    string             `json:"memberId"`    // UserID
		AccountType string             `json:"accountType"` // Account type
		Balance     []CoinBalanceEntry `json:"balance"`     // Array of balance entries
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetSingleCoinBalanceRequest represents the query parameters for fetching a single coin balance.
type GetSingleCoinBalanceRequest struct {
	MemberID                  *string `json:"memberId,omitempty"`                  // Optional: UID, required for querying sub UID balance with master API key
	ToMemberID                *string `json:"toMemberId,omitempty"`                // Optional: UID, required for querying transferable balance between UIDs
	AccountType               string  `json:"accountType"`                         // Required: Account type
	ToAccountType             *string `json:"toAccountType,omitempty"`             // Optional: To account type, required for transferable balance queries
	Coin                      string  `json:"coin"`                                // Required: Coin
	WithBonus                 *int    `json:"withBonus,omitempty"`                 // Optional: Query bonus
	WithTransferSafeAmount    *int    `json:"withTransferSafeAmount,omitempty"`    // Optional: Query delay withdraw/transfer safe amount
	WithLtvTransferSafeAmount *int    `json:"withLtvTransferSafeAmount,omitempty"` // Optional: Query transferable amount for ins loan account
}

// SingleCoinBalanceEntry represents the balance information for a single coin.
type SingleCoinBalanceEntry struct {
	Coin                  string `json:"coin"`
	WalletBalance         string `json:"walletBalance"`
	TransferBalance       string `json:"transferBalance"`
	Bonus                 string `json:"bonus,omitempty"`
	TransferSafeAmount    string `json:"transferSafeAmount,omitempty"`
	LtvTransferSafeAmount string `json:"ltvTransferSafeAmount,omitempty"`
}

// GetSingleCoinBalanceResponse represents the response from fetching a single coin balance.
type GetSingleCoinBalanceResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		AccountType string                 `json:"accountType"`
		BizType     int                    `json:"bizType"`
		AccountID   string                 `json:"accountId"`
		MemberID    string                 `json:"memberId"`
		Balance     SingleCoinBalanceEntry `json:"balance"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetTransferableCoinRequest represents the query parameters for fetching the transferable coin list.
type GetTransferableCoinRequest struct {
	FromAccountType string `json:"fromAccountType"` // Required: From account type
	ToAccountType   string `json:"toAccountType"`   // Required: To account type
}

// GetTransferableCoinResponse represents the response from fetching the transferable coin list.
type GetTransferableCoinResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []string `json:"list"` // A list of coins
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// CreateInternalTransferRequest represents the payload for creating an internal transfer.
type CreateInternalTransferRequest struct {
	TransferID      string `json:"transferId"`      // Required: UUID, manually generate a UUID
	Coin            string `json:"coin"`            // Required: Coin
	Amount          string `json:"amount"`          // Required: Amount
	FromAccountType string `json:"fromAccountType"` // Required: From account type
	ToAccountType   string `json:"toAccountType"`   // Required: To account type
}

// CreateInternalTransferResponse represents the response from creating an internal transfer.
type CreateInternalTransferResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		TransferID string `json:"transferId"` // UUID
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetUniversalTransferRecordsRequest represents the query parameters for fetching universal transfer records.
type GetUniversalTransferRecordsRequest struct {
	TransferID *string `json:"transferId,omitempty"` // Optional: UUID used in createTransfer
	Coin       *string `json:"coin,omitempty"`       // Optional: Coin
	Status     *string `json:"status,omitempty"`     // Optional: Transfer status (SUCCESS, FAILED, PENDING)
	StartTime  *int64  `json:"startTime,omitempty"`  // Optional: Start timestamp (ms)
	EndTime    *int64  `json:"endTime,omitempty"`    // Optional: End timestamp (ms)
	Limit      *int    `json:"limit,omitempty"`      // Optional: Data size limit per page
	Cursor     *string `json:"cursor,omitempty"`     // Optional: Pagination cursor
}

// UniversalTransferRecordEntry represents a single entry in the universal transfer record list.
type UniversalTransferRecordEntry struct {
	TransferID      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromMemberID    string `json:"fromMemberId"`
	ToMemberID      string `json:"toMemberId"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

// GetUniversalTransferRecordsResponse represents the response from fetching universal transfer records.
type GetUniversalTransferRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []UniversalTransferRecordEntry `json:"list"`
		NextPageCursor string                         `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetInternalTransferRecordsRequest represents the query parameters for fetching internal transfer records.
type GetInternalTransferRecordsRequest struct {
	TransferID *string `json:"transferId,omitempty"` // Optional: UUID used in createTransfer
	Coin       *string `json:"coin,omitempty"`       // Optional: Coin
	Status     *string `json:"status,omitempty"`     // Optional: Transfer status
	StartTime  *int64  `json:"startTime,omitempty"`  // Optional: Start timestamp (ms)
	EndTime    *int64  `json:"endTime,omitempty"`    // Optional: End timestamp (ms)
	Limit      *int    `json:"limit,omitempty"`      // Optional: Data size limit per page
	Cursor     *string `json:"cursor,omitempty"`     // Optional: Pagination cursor
}

// InternalTransferRecordEntry represents a single entry in the internal transfer record list.
type InternalTransferRecordEntry struct {
	TransferID      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

// GetInternalTransferRecordsResponse represents the response from fetching internal transfer records.
type GetInternalTransferRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List           []InternalTransferRecordEntry `json:"list"`
		NextPageCursor string                        `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// GetSubUIDsResponse represents the response from fetching sub UIDs.
type GetSubUIDsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		SubMemberIds             []string `json:"subMemberIds"`             // All sub UIDs under the main UID
		TransferableSubMemberIds []string `json:"transferableSubMemberIds"` // All sub UIDs that have universal transfer enabled
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// CreateUniversalTransferRequest represents the payload for creating a universal transfer.
type CreateUniversalTransferRequest struct {
	TransferID      string `json:"transferId"`      // Required: UUID, manually generate a UUID
	Coin            string `json:"coin"`            // Required: Coin
	Amount          string `json:"amount"`          // Required: Amount
	FromMemberID    int    `json:"fromMemberId"`    // Required: From UID
	ToMemberID      int    `json:"toMemberId"`      // Required: To UID
	FromAccountType string `json:"fromAccountType"` // Required: From account type
	ToAccountType   string `json:"toAccountType"`   // Required: To account type
}

// CreateUniversalTransferResponse represents the response from creating a universal transfer.
type CreateUniversalTransferResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		TransferID string `json:"transferId"` // UUID
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetAllowedDepositCoinInfoRequest struct {
	Coin   *string `json:"coin,omitempty"`   // Optional: Coin. coin and chain must be paired if passed
	Chain  *string `json:"chain,omitempty"`  // Optional: Chain. coin and chain must be paired if passed
	Limit  *int    `json:"limit,omitempty"`  // Optional: Limit for data size per page
	Cursor *string `json:"cursor,omitempty"` // Optional: Pagination cursor
}

type DepositConfig struct {
	Coin               string `json:"coin"`
	Chain              string `json:"chain"`
	CoinShowName       string `json:"coinShowName"`
	ChainType          string `json:"chainType"`
	BlockConfirmNumber int    `json:"blockConfirmNumber"`
	MinDepositAmount   string `json:"minDepositAmount"`
}

type GetAllowedDepositCoinInfoResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		ConfigList     []DepositConfig `json:"configList"`
		NextPageCursor string          `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type SetDepositAccountRequest struct {
	AccountType string `json:"accountType"` // Required: Account type (UNIFIED, SPOT, OPTION, CONTRACT, FUND)
}

type SetDepositAccountResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Status int `json:"status"` // Request result: 1 for SUCCESS, 0 for FAIL
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetDepositRecordsRequest struct {
	Coin      *string `json:"coin,omitempty"`      // Optional: Coin
	StartTime *int64  `json:"startTime,omitempty"` // Optional: The start timestamp (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: The end timestamp (ms)
	Limit     *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Pagination cursor
}

type DepositRecordEntry struct {
	Coin              string `json:"coin"`
	Chain             string `json:"chain"`
	Amount            string `json:"amount"`
	TxID              string `json:"txID"`
	Status            int    `json:"status"`
	ToAddress         string `json:"toAddress"`
	Tag               string `json:"tag"`
	DepositFee        string `json:"depositFee"`
	SuccessAt         string `json:"successAt"`
	Confirmations     string `json:"confirmations"`
	TxIndex           string `json:"txIndex"`
	BlockHash         string `json:"blockHash"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
	DepositType       int    `json:"depositType"`
}

type GetDepositRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Rows           []DepositRecordEntry `json:"rows"`
		NextPageCursor string               `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetSubDepositRecordsRequest struct {
	SubMemberId string  `json:"subMemberId"`         // Required: Sub UID
	Coin        *string `json:"coin,omitempty"`      // Optional: Coin
	StartTime   *int64  `json:"startTime,omitempty"` // Optional: The start timestamp (ms)
	EndTime     *int64  `json:"endTime,omitempty"`   // Optional: The end timestamp (ms)
	Limit       *int    `json:"limit,omitempty"`     // Optional: Limit for data size per page
	Cursor      *string `json:"cursor,omitempty"`    // Optional: Pagination cursor
}

type GetSubDepositRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Rows           []DepositRecordEntry `json:"rows"`
		NextPageCursor string               `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetInternalDepositRecordsRequest struct {
	TxID      *string `json:"txID,omitempty"`      // Optional: Internal transfer transaction ID
	StartTime *int64  `json:"startTime,omitempty"` // Optional: Start time (ms)
	EndTime   *int64  `json:"endTime,omitempty"`   // Optional: End time (ms)
	Coin      *string `json:"coin,omitempty"`      // Optional: Coin name
	Cursor    *string `json:"cursor,omitempty"`    // Optional: Pagination cursor
	Limit     *int    `json:"limit,omitempty"`     // Optional: Number of items per page
}

type InternalDepositRecordEntry struct {
	ID          string `json:"id"`
	Type        int    `json:"type"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	Status      int    `json:"status"`
	Address     string `json:"address"`
	CreatedTime string `json:"createdTime"`
	TxID        string `json:"txID"`
}

type GetInternalDepositRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Rows           []InternalDepositRecordEntry `json:"rows"`
		NextPageCursor string                       `json:"nextPageCursor"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetMasterDepositAddressRequest struct {
	Coin      string  `json:"coin"`                // Required: Coin
	ChainType *string `json:"chainType,omitempty"` // Optional: Chain type from coin-info endpoint
}

type DepositChainInfo struct {
	ChainType         string `json:"chainType"`
	AddressDeposit    string `json:"addressDeposit"`
	TagDeposit        string `json:"tagDeposit"`
	Chain             string `json:"chain"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
}

type GetMasterDepositAddressResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Coin   string             `json:"coin"`
		Chains []DepositChainInfo `json:"chains"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type GetSubDepositAddressRequest struct {
	Coin        string `json:"coin"`        // Required: Coin
	ChainType   string `json:"chainType"`   // Required: Chain type from the coin-info endpoint
	SubMemberId string `json:"subMemberId"` // Required: Sub user ID
}

type SubDepositChainInfo struct {
	ChainType         string `json:"chainType"`
	AddressDeposit    string `json:"addressDeposit"`
	TagDeposit        string `json:"tagDeposit,omitempty"`
	Chain             string `json:"chain"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
}

type GetSubDepositAddressResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Coin   string                `json:"coin"`
		Chains []SubDepositChainInfo `json:"chains"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
type CoinChainInfo struct {
	Chain                 string `json:"chain"`
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	MinAccuracy           string `json:"minAccuracy"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
}

type CoinInfoEntry struct {
	Name         string          `json:"name"`
	Coin         string          `json:"coin"`
	RemainAmount string          `json:"remainAmount"`
	Chains       []CoinChainInfo `json:"chains"`
}

type GetCoinInfoResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Rows []CoinInfoEntry `json:"rows"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

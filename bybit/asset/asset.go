package asset

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Asset interface {
	// GetCoinExchangeRecords queries the coin exchange records.
	GetCoinExchangeRecords(req *GetCoinExchangeRecordsRequest) (*GetCoinExchangeRecordsResponse, error)
	// GetDeliveryRecords queries the delivery records of USDC futures and Options.
	GetDeliveryRecords(req *GetDeliveryRecordRequest) (*GetDeliveryRecordResponse, error)
	// GetSessionSettlementRecords queries the session settlement records of USDC perpetual and futures.
	GetSessionSettlementRecords(req *GetSessionSettlementRecordRequest) (*GetSessionSettlementRecordResponse, error)
	// GetAssetInfo queries the asset information for SPOT accounts.
	GetAssetInfo(req *GetAssetInfoRequest) (*GetAssetInfoResponse, error)
	// GetAllCoinsBalance retrieves all coin balances for specified account types.
	GetAllCoinsBalance(req *GetAllCoinsBalanceRequest) (*GetAllCoinsBalanceResponse, error)
	// GetSingleCoinBalance queries the balance of a specific coin in a specific account type.
	GetSingleCoinBalance(req *GetSingleCoinBalanceRequest) (*GetSingleCoinBalanceResponse, error)
	// GetTransferableCoin queries the list of transferable coins between account types.
	GetTransferableCoin(req *GetTransferableCoinRequest) (*GetTransferableCoinResponse, error)
	CreateInternalTransfer(req *CreateInternalTransferRequest) (*CreateInternalTransferResponse, error)
	GetInternalTransferRecords(req *GetInternalTransferRecordsRequest) (*GetInternalTransferRecordsResponse, error)
	GetSubUIDs() (*GetSubUIDsResponse, error)
	CreateUniversalTransfer(req *CreateUniversalTransferRequest) (*CreateUniversalTransferResponse, error)
	GetUniversalTransferRecords(req *GetUniversalTransferRecordsRequest) (*GetUniversalTransferRecordsResponse, error)
}

type impl struct {
	client *client.Client
}

func New(client *client.Client) Asset {
	return &impl{
		client: client,
	}
}
func (i *impl) GetCoinExchangeRecords(req *GetCoinExchangeRecordsRequest) (*GetCoinExchangeRecordsResponse, error) {
	var allRecords []CoinExchangeRecord
	var finalResponse GetCoinExchangeRecordsResponse

	for {
		// Construct query parameters for each iteration
		queryParams := make(client.Params)
		if req.FromCoin != nil {
			queryParams["fromCoin"] = *req.FromCoin
		}
		if req.ToCoin != nil {
			queryParams["toCoin"] = *req.ToCoin
		}
		if req.Limit != nil {
			queryParams["limit"] = strconv.Itoa(*req.Limit)
		}
		if req.Cursor != nil {
			queryParams["cursor"] = *req.Cursor
		}

		// Perform the GET request
		response, err := i.client.Get("/v5/asset/exchange/order-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching coin exchange records: %w", err)
		}
		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}

		// Parse the JSON response for each iteration
		var exchangeRecordsResponse GetCoinExchangeRecordsResponse
		if err := json.Unmarshal(data, &exchangeRecordsResponse); err != nil {
			return nil, fmt.Errorf("error parsing coin exchange records response: %w", err)
		}

		// Accumulate records from the current page
		allRecords = append(allRecords, exchangeRecordsResponse.Result.OrderBody...)

		// Prepare for the next iteration or break the loop
		if exchangeRecordsResponse.Result.NextPageCursor == "" {
			break // No more pages
		}
		req.Cursor = &exchangeRecordsResponse.Result.NextPageCursor // Set cursor for next page
	}
	finalResponse.RetCode = 0
	finalResponse.RetMsg = "OK"
	finalResponse.Result.OrderBody = allRecords
	finalResponse.Result.NextPageCursor = ""
	return &finalResponse, nil
}
func (i *impl) GetDeliveryRecords(req *GetDeliveryRecordRequest) (*GetDeliveryRecordResponse, error) {
	var allRecords []DeliveryRecordEntry
	var finalResponse GetDeliveryRecordResponse

	for {
		// Prepare query parameters for each request
		queryParams := make(client.Params)
		queryParams["category"] = req.Category
		if req.Symbol != nil {
			queryParams["symbol"] = *req.Symbol
		}
		if req.StartTime != nil {
			queryParams["startTime"] = strconv.FormatInt(*req.StartTime, 10)
		}
		if req.EndTime != nil {
			queryParams["endTime"] = strconv.FormatInt(*req.EndTime, 10)
		}
		if req.ExpDate != nil {
			queryParams["expDate"] = *req.ExpDate
		}
		if req.Limit != nil {
			queryParams["limit"] = strconv.Itoa(*req.Limit)
		}
		if req.Cursor != nil {
			queryParams["cursor"] = *req.Cursor
		}

		// Perform the GET request
		response, err := i.client.Get("/v5/asset/delivery-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching delivery records: %w", err)
		}
		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		var currentPageResponse GetDeliveryRecordResponse
		if err := json.Unmarshal(data, &currentPageResponse); err != nil {
			return nil, fmt.Errorf("error parsing delivery records response: %w", err)
		}

		// Accumulate records from the current page
		allRecords = append(allRecords, currentPageResponse.Result.List...)

		// Check if there's a next page
		if currentPageResponse.Result.NextPageCursor == "" {
			break // Exit loop if there's no next page cursor
		} else {
			// Update the cursor for the next request
			req.Cursor = &currentPageResponse.Result.NextPageCursor
		}
	}

	finalResponse.RetCode = 0
	finalResponse.RetMsg = "OK"
	finalResponse.Result.List = allRecords
	finalResponse.Result.NextPageCursor = ""
	return &finalResponse, nil
}
func (i *impl) GetSessionSettlementRecords(req *GetSessionSettlementRecordRequest) (*GetSessionSettlementRecordResponse, error) {
	queryParams := make(client.Params)
	queryParams["category"] = req.Category
	if req.Symbol != nil {
		queryParams["symbol"] = *req.Symbol
	}
	if req.StartTime != nil {
		queryParams["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		queryParams["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.Limit != nil {
		queryParams["limit"] = strconv.Itoa(*req.Limit)
	}
	if req.Cursor != nil {
		queryParams["cursor"] = *req.Cursor
	}

	// Perform the GET request with pagination logic to fetch all records
	var allRecords []SessionSettlementRecord
	var finalResponse GetSessionSettlementRecordResponse

	for {
		response, err := i.client.Get("/v5/asset/settlement-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching session settlement records: %w", err)
		}
		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		var pageResponse GetSessionSettlementRecordResponse
		if err := json.Unmarshal(data, &pageResponse); err != nil {
			return nil, fmt.Errorf("error parsing session settlement records response: %w", err)
		}

		// Accumulate records from the current page
		allRecords = append(allRecords, pageResponse.Result.List...)

		// Check if there's a next page
		if pageResponse.Result.NextPageCursor == "" {
			break // Exit the loop if there's no next page cursor
		} else {
			// Update the cursor for the next request
			queryParams["cursor"] = pageResponse.Result.NextPageCursor
		}
	}

	finalResponse.RetCode = 0
	finalResponse.RetMsg = "OK"
	finalResponse.Result.List = allRecords
	finalResponse.Result.NextPageCursor = ""

	return &finalResponse, nil
}

func (i *impl) GetAssetInfo(req *GetAssetInfoRequest) (*GetAssetInfoResponse, error) {
	queryParams := make(client.Params)
	queryParams["accountType"] = req.AccountType
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-asset-info", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching asset information: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var assetInfoResponse GetAssetInfoResponse
	if err := json.Unmarshal(data, &assetInfoResponse); err != nil {
		return nil, fmt.Errorf("error parsing asset information response: %w", err)
	}

	return &assetInfoResponse, nil
}

func (i *impl) GetSingleCoinBalance(req *GetSingleCoinBalanceRequest) (*GetSingleCoinBalanceResponse, error) {
	queryParams := make(client.Params)
	if req.MemberID != nil {
		queryParams["memberId"] = *req.MemberID
	}
	if req.ToMemberID != nil {
		queryParams["toMemberId"] = *req.ToMemberID
	}
	queryParams["accountType"] = req.AccountType
	if req.ToAccountType != nil {
		queryParams["toAccountType"] = *req.ToAccountType
	}
	queryParams["coin"] = req.Coin
	if req.WithBonus != nil {
		queryParams["withBonus"] = strconv.Itoa(*req.WithBonus)
	}
	if req.WithTransferSafeAmount != nil {
		queryParams["withTransferSafeAmount"] = strconv.Itoa(*req.WithTransferSafeAmount)
	}
	if req.WithLtvTransferSafeAmount != nil {
		queryParams["withLtvTransferSafeAmount"] = strconv.Itoa(*req.WithLtvTransferSafeAmount)
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-account-coin-balance", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching single coin balance: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var coinBalanceResponse GetSingleCoinBalanceResponse
	if err := json.Unmarshal(data, &coinBalanceResponse); err != nil {
		return nil, fmt.Errorf("error parsing single coin balance response: %w", err)
	}

	return &coinBalanceResponse, nil
}
func (i *impl) GetTransferableCoin(req *GetTransferableCoinRequest) (*GetTransferableCoinResponse, error) {
	// Prepare query parameters
	queryParams := make(client.Params)
	queryParams["fromAccountType"] = req.FromAccountType
	queryParams["toAccountType"] = req.ToAccountType

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-transfer-coin-list", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching transferable coin list: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var transferableCoinResponse GetTransferableCoinResponse
	if err := json.Unmarshal(data, &transferableCoinResponse); err != nil {
		return nil, fmt.Errorf("error parsing transferable coin list response: %w", err)
	}

	return &transferableCoinResponse, nil
}

func (i *impl) GetAllCoinsBalance(req *GetAllCoinsBalanceRequest) (*GetAllCoinsBalanceResponse, error) {
	queryParams := make(client.Params)
	if req.MemberID != nil {
		queryParams["memberId"] = *req.MemberID
	}
	queryParams["accountType"] = req.AccountType
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}
	if req.WithBonus != nil {
		queryParams["withBonus"] = strconv.Itoa(*req.WithBonus)
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-account-coins-balance", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching all coins balance: %w", err)
	}

	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	var coinsBalanceResponse GetAllCoinsBalanceResponse
	if err := json.Unmarshal(data, &coinsBalanceResponse); err != nil {
		return nil, fmt.Errorf("error parsing all coins balance response: %w", err)
	}

	return &coinsBalanceResponse, nil
}
func (i *impl) CreateInternalTransfer(req *CreateInternalTransferRequest) (*CreateInternalTransferResponse, error) {
	// Initialize Params and populate with request data
	params := client.Params{
		"transferId":      req.TransferID,
		"coin":            req.Coin,
		"amount":          req.Amount,
		"fromAccountType": req.FromAccountType,
		"toAccountType":   req.ToAccountType,
	}

	// Ensure all required fields are provided
	if req.TransferID == "" || req.Coin == "" || req.Amount == "" || req.FromAccountType == "" || req.ToAccountType == "" {
		return nil, errors.New("missing required fields in request")
	}

	// Perform the POST request
	response, err := i.client.Post("/v5/asset/transfer/inter-transfer", params)
	if err != nil {
		return nil, fmt.Errorf("error creating internal transfer: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the CreateInternalTransferResponse struct
	var transferResponse CreateInternalTransferResponse
	if err := json.Unmarshal(data, &transferResponse); err != nil {
		return nil, fmt.Errorf("error parsing internal transfer response: %w", err)
	}

	return &transferResponse, nil
}

func (i *impl) GetUniversalTransferRecords(req *GetUniversalTransferRecordsRequest) (*GetUniversalTransferRecordsResponse, error) {
	queryParams := client.Params{}
	if req.TransferID != nil {
		queryParams["transferId"] = *req.TransferID
	}
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}
	if req.Status != nil {
		queryParams["status"] = *req.Status
	}
	if req.StartTime != nil {
		queryParams["startTime"] = *req.StartTime
	}
	if req.EndTime != nil {
		queryParams["endTime"] = *req.EndTime
	}
	if req.Limit != nil {
		queryParams["limit"] = *req.Limit
	}
	if req.Cursor != nil {
		queryParams["cursor"] = *req.Cursor
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-universal-transfer-list", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching universal transfer records: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var transferRecordsResponse GetUniversalTransferRecordsResponse
	if err := json.Unmarshal(data, &transferRecordsResponse); err != nil {
		return nil, fmt.Errorf("error parsing universal transfer records response: %w", err)
	}

	return &transferRecordsResponse, nil
}
func (i *impl) GetInternalTransferRecords(req *GetInternalTransferRecordsRequest) (*GetInternalTransferRecordsResponse, error) {
	queryParams := make(client.Params)
	if req.TransferID != nil {
		queryParams["transferId"] = *req.TransferID
	}
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}
	if req.Status != nil {
		queryParams["status"] = *req.Status
	}
	if req.StartTime != nil {
		queryParams["startTime"] = *req.StartTime
	}
	if req.EndTime != nil {
		queryParams["endTime"] = *req.EndTime
	}
	if req.Limit != nil {
		queryParams["limit"] = *req.Limit
	}
	if req.Cursor != nil {
		queryParams["cursor"] = *req.Cursor
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-inter-transfer-list", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching internal transfer records: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var transferRecordsResponse GetInternalTransferRecordsResponse
	err = json.Unmarshal(data, &transferRecordsResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing internal transfer records response: %w", err)
	}

	return &transferRecordsResponse, nil
}
func (i *impl) GetSubUIDs() (*GetSubUIDsResponse, error) {
	// Perform the GET request
	response, err := i.client.Get("/v5/asset/transfer/query-sub-member-list", nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching sub UIDs: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var subUIDsResponse GetSubUIDsResponse
	err = json.Unmarshal(data, &subUIDsResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing sub UIDs response: %w", err)
	}

	return &subUIDsResponse, nil
}
func (i *impl) CreateUniversalTransfer(req *CreateUniversalTransferRequest) (*CreateUniversalTransferResponse, error) {
	queryParams := make(client.Params)
	queryParams["transferId"] = req.TransferID
	queryParams["coin"] = req.Coin
	queryParams["amount"] = req.Amount
	queryParams["fromMemberId"] = req.FromMemberID
	queryParams["toMemberId"] = req.ToMemberID
	queryParams["fromAccountType"] = req.FromAccountType
	queryParams["toAccountType"] = req.ToAccountType

	// Ensure all required fields are populated
	if req.TransferID == "" || req.Coin == "" || req.Amount == "" || req.FromMemberID == 0 || req.ToMemberID == 0 || req.FromAccountType == "" || req.ToAccountType == "" {
		return nil, errors.New("missing required fields in request")
	}

	// Perform the POST request
	response, err := i.client.Post("/v5/asset/transfer/universal-transfer", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error creating universal transfer: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var transferResponse CreateUniversalTransferResponse
	err = json.Unmarshal(data, &transferResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing universal transfer response: %w", err)
	}

	return &transferResponse, nil
}
func (i *impl) GetAllowedDepositCoinInfo(req *GetAllowedDepositCoinInfoRequest) (*GetAllowedDepositCoinInfoResponse, error) {
	queryParams := make(client.Params)
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}
	if req.Chain != nil {
		queryParams["chain"] = *req.Chain
	}
	if req.Limit != nil {
		queryParams["limit"] = *req.Limit
	}
	if req.Cursor != nil {
		queryParams["cursor"] = *req.Cursor
	}

	// Perform the GET request
	response, err := i.client.Get("/v5/asset/deposit/query-allowed-list", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching allowed deposit coin information: %w", err)
	}

	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var allowedDepositCoinInfoResponse GetAllowedDepositCoinInfoResponse
	err = json.Unmarshal(data, &allowedDepositCoinInfoResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing allowed deposit coin information response: %w", err)
	}

	return &allowedDepositCoinInfoResponse, nil
}

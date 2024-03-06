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
	GetAllowedDepositCoinInfo(req *GetAllowedDepositCoinInfoRequest) (*GetAllowedDepositCoinInfoResponse, error)
	GetDepositRecords(req *GetDepositRecordsRequest) (*GetDepositRecordsResponse, error)
	GetSubDepositRecords(req *GetSubDepositRecordsRequest) (*GetSubDepositRecordsResponse, error)
	GetInternalDepositRecords(req *GetInternalDepositRecordsRequest) (*GetInternalDepositRecordsResponse, error)
	GetMasterDepositAddress(req *GetMasterDepositAddressRequest) (*GetMasterDepositAddressResponse, error)
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
func (i *impl) SetDepositAccount(req *SetDepositAccountRequest) (*SetDepositAccountResponse, error) {
	// Initialize Params and populate with request data
	params := client.Params{
		"accountType": req.AccountType, // Direct assignment since AccountType is required and assumed to be always provided
	}

	responseBytes, err := i.client.Post("/v5/asset/deposit/deposit-to-account", params)
	if err != nil {
		return nil, fmt.Errorf("error during POST request for setting deposit account: %w", err)
	}
	data, err := json.Marshal(responseBytes)
	if err != nil {
		return nil, err
	}

	var response SetDepositAccountResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response from setting deposit account: %w", err)
	}

	return &response, nil
}
func (i *impl) GetDepositRecords(req *GetDepositRecordsRequest) (*GetDepositRecordsResponse, error) {
	allDepositRecords := []DepositRecordEntry{}
	var finalResponse GetDepositRecordsResponse

	// Initial queryParams setup
	queryParams := make(client.Params)
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
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

	for {
		// Perform the GET request
		response, err := i.client.Get("/v5/asset/deposit/query-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching deposit records: %w", err)
		}

		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}

		// Deserialize the current page of response
		var currentPageResponse GetDepositRecordsResponse
		err = json.Unmarshal(data, &currentPageResponse)
		if err != nil {
			return nil, fmt.Errorf("error parsing deposit records response: %w", err)
		}

		// Accumulate records from the current page
		allDepositRecords = append(allDepositRecords, currentPageResponse.Result.Rows...)

		// Check if there's a next page. If not, break out of the loop
		if currentPageResponse.Result.NextPageCursor == "" {
			finalResponse = currentPageResponse // Use the last page's meta info for final response
			break
		}

		// Set cursor for the next page
		queryParams["cursor"] = currentPageResponse.Result.NextPageCursor
	}

	// Populate finalResponse with all accumulated records
	finalResponse.Result.Rows = allDepositRecords
	finalResponse.Result.NextPageCursor = ""
	finalResponse.RetCode = 0

	return &finalResponse, nil
}
func (i *impl) GetSubDepositRecords(req *GetSubDepositRecordsRequest) (*GetSubDepositRecordsResponse, error) {
	var allRows []DepositRecordEntry
	var finalResponse GetSubDepositRecordsResponse

	queryParams := make(client.Params)
	queryParams["subMemberId"] = req.SubMemberId
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
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
	queryParams["cursor"] = req.Cursor // Start with nil or provided cursor

	for {
		response, err := i.client.Get("/v5/asset/deposit/query-sub-member-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching sub deposit records: %w", err)
		}

		// Assuming the response is already unmarshaled into the appropriate struct
		var currentPageResponse GetSubDepositRecordsResponse
		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &currentPageResponse)
		if err != nil {
			return nil, fmt.Errorf("error parsing sub deposit records response: %w", err)
		}
		allRows = append(allRows, currentPageResponse.Result.Rows...)
		if currentPageResponse.Result.NextPageCursor == "" {
			break
		}
		queryParams["cursor"] = currentPageResponse.Result.NextPageCursor // Prepare for the next iteration
	}

	// Assign collected rows and last page's meta to finalResponse
	finalResponse.Result.Rows = allRows
	finalResponse.Result.NextPageCursor = ""
	finalResponse.RetCode = 0
	return &finalResponse, nil
}
func (i *impl) GetInternalDepositRecords(req *GetInternalDepositRecordsRequest) (*GetInternalDepositRecordsResponse, error) {
	var allRows []InternalDepositRecordEntry
	var finalResponse GetInternalDepositRecordsResponse

	queryParams := make(client.Params)
	if req.TxID != nil {
		queryParams["txID"] = *req.TxID
	}
	if req.StartTime != nil {
		queryParams["startTime"] = *req.StartTime
	}
	if req.EndTime != nil {
		queryParams["endTime"] = *req.EndTime
	}
	if req.Coin != nil {
		queryParams["coin"] = *req.Coin
	}
	if req.Cursor != nil {
		queryParams["cursor"] = *req.Cursor
	}
	if req.Limit != nil {
		queryParams["limit"] = *req.Limit
	}
	var currentPageResponse GetInternalDepositRecordsResponse
	// Loop through pages to collect all records
	for {
		response, err := i.client.Get("/v5/asset/deposit/query-internal-record", queryParams)
		if err != nil {
			return nil, fmt.Errorf("error fetching internal deposit records: %w", err)
		}

		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		// Assuming response is a JSON body byte slice

		err = json.Unmarshal(data, &currentPageResponse)
		if err != nil {
			return nil, fmt.Errorf("error parsing internal deposit records response: %w", err)
		}

		allRows = append(allRows, currentPageResponse.Result.Rows...)
		if currentPageResponse.Result.NextPageCursor == "" {
			finalResponse = currentPageResponse
			break
		}

		queryParams["cursor"] = currentPageResponse.Result.NextPageCursor // Prepare for the next iteration
	}

	finalResponse.Result.Rows = allRows
	finalResponse.Result.NextPageCursor = currentPageResponse.Result.NextPageCursor
	finalResponse.RetCode = currentPageResponse.RetCode
	finalResponse.RetExtInfo = currentPageResponse.RetExtInfo
	finalResponse.RetMsg = currentPageResponse.RetMsg
	return &finalResponse, nil
}

func (i *impl) GetMasterDepositAddress(req *GetMasterDepositAddressRequest) (*GetMasterDepositAddressResponse, error) {
	queryParams := make(client.Params)
	queryParams["coin"] = req.Coin
	if req.ChainType != nil {
		queryParams["chainType"] = *req.ChainType
	}

	// Perform the GET request
	responseBytes, err := i.client.Get("/v5/asset/deposit/query-address", queryParams)
	if err != nil {
		return nil, fmt.Errorf("error querying master deposit address: %w", err)
	}

	data, err := json.Marshal(responseBytes)
	if err != nil {
		return nil, err
	}
	// Deserialize the response into the response struct
	var response GetMasterDepositAddressResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing master deposit address response: %w", err)
	}

	return &response, nil
}

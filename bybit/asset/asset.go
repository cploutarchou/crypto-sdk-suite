package asset

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"strconv"
)

type Asset interface {
	// GetCoinExchangeRecords queries the coin exchange records.
	GetCoinExchangeRecords(req *GetCoinExchangeRecordsRequest) (*GetCoinExchangeRecordsResponse, error)
	// GetDeliveryRecords queries the delivery records of USDC futures and Options.
	GetDeliveryRecords(req *GetDeliveryRecordRequest) (*GetDeliveryRecordResponse, error)
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
	finalResponse.RetCode = 0 // Assume success or set based on your logic
	finalResponse.RetMsg = "OK"
	finalResponse.Result.OrderBody = allRecords
	finalResponse.Result.NextPageCursor = ""
	return &finalResponse, nil
}
func (i *impl) GetDeliveryRecords(req *GetDeliveryRecordRequest) (*GetDeliveryRecordResponse, error) {
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
	var deliveryRecordsResponse GetDeliveryRecordResponse
	if err := json.Unmarshal(data, &deliveryRecordsResponse); err != nil {
		return nil, fmt.Errorf("error parsing delivery records response: %w", err)
	}

	return &deliveryRecordsResponse, nil
}

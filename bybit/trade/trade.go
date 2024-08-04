package trade

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"strings"
)

type Trade interface {
	PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error)
	AmendOrder(req *AmendOrderRequest) (*AmendOrderResponse, error)
	CancelOrder(req *CancelOrderRequest) (*CancelOrderResponse, error)
	GetOpenOrders(req *GetOpenOrdersRequest) (*GetOpenOrdersResponse, error)
	CancelAllOrders(req *CancelAllOrdersRequest) (*CancelAllOrdersResponse, error)
	GetOrderHistory(req *GetOrderHistoryRequest) (*GetOrderHistoryResponse, error)
	GetTradeHistory(req *GetTradeHistoryRequest) (*GetTradeHistoryResponse, error)
	BatchPlaceOrder(req *BatchPlaceOrderRequest) (*BatchPlaceOrderResponse, error)
	GetBorrowQuotaSpot(symbol, side string) (*BorrowQuotaResponse, error)
}

// Helper function to generate cURL command from request parameters
func generateCurlCommand(endpoint string, params map[string]string) string {
	var paramList []string
	for key, value := range params {
		paramList = append(paramList, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}
	queryString := strings.Join(paramList, "&")
	return fmt.Sprintf("curl -X GET '%s?%s'", endpoint, queryString)
}

type tradeImpl struct {
	client *client.Client
}

func New(c *client.Client) Trade {
	return &tradeImpl{client: c}
}

func (t *tradeImpl) PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	params := ConvertPlaceOrderRequestToParams(req)
	res, err := t.client.Post("/v5/order/create", params)
	if err != nil {
		return nil, err
	}

	var placeOrderResponse PlaceOrderResponse
	err = res.Unmarshal(&placeOrderResponse)
	if err != nil {
		return nil, err
	}
	if placeOrderResponse.RetCode != 0 {
		return &placeOrderResponse, fmt.Errorf("API returned error: %s", placeOrderResponse.RetMsg)
	}
	return &placeOrderResponse, nil
}

func ConvertPlaceOrderRequestToParams(req *PlaceOrderRequest) client.Params {
	params := client.Params{
		"category":    req.Category,
		"symbol":      req.Symbol,
		"side":        req.Side,
		"orderType":   req.OrderType,
		"qty":         req.Qty,
		"orderLinkId": req.OrderLinkId,
	}

	if req.Price != "" {
		params["price"] = req.Price
	}

	if req.IsLeverage != 0 {
		params["isLeverage"] = strconv.Itoa(req.IsLeverage)
	}
	if req.TriggerPrice != nil {
		params["triggerPrice"] = *req.TriggerPrice
	}
	if req.TriggerDirection != nil {
		params["triggerDirection"] = strconv.Itoa(*req.TriggerDirection)
	}
	if req.TriggerBy != nil {
		params["triggerBy"] = *req.TriggerBy
	}
	if req.OrderFilter != nil {
		params["orderFilter"] = *req.OrderFilter
	}
	if req.OrderIv != nil {
		params["orderIv"] = *req.OrderIv
	}
	if req.TimeInForce != "" {
		params["timeInForce"] = req.TimeInForce
	}
	if req.PositionIdx != nil {
		params["positionIdx"] = strconv.Itoa(*req.PositionIdx)
	}
	if req.TakeProfit != nil {
		params["takeProfit"] = *req.TakeProfit
	}
	if req.StopLoss != nil {
		params["stopLoss"] = *req.StopLoss
	}
	if req.TpTriggerBy != nil {
		params["tpTriggerBy"] = *req.TpTriggerBy
	}
	if req.SlTriggerBy != nil {
		params["slTriggerBy"] = *req.SlTriggerBy
	}
	if req.ReduceOnly != nil {
		params["reduceOnly"] = strconv.FormatBool(*req.ReduceOnly)
	}
	if req.CloseOnTrigger != nil {
		params["closeOnTrigger"] = strconv.FormatBool(*req.CloseOnTrigger)
	}
	if req.SmpType != nil {
		params["smpType"] = *req.SmpType
	}
	if req.Mmp != nil {
		params["mmp"] = strconv.FormatBool(*req.Mmp)
	}
	if req.TpslMode != nil {
		params["tpslMode"] = *req.TpslMode
	}
	if req.TpLimitPrice != nil {
		params["tpLimitPrice"] = *req.TpLimitPrice
	}
	if req.SlLimitPrice != nil {
		params["slLimitPrice"] = *req.SlLimitPrice
	}
	if req.TpOrderType != nil {
		params["tpOrderType"] = *req.TpOrderType
	}
	if req.SlOrderType != nil {
		params["slOrderType"] = *req.SlOrderType
	}

	return params
}
func (t *tradeImpl) AmendOrder(req *AmendOrderRequest) (*AmendOrderResponse, error) {
	params := ConvertAmendOrderRequestToParams(req)
	res, err := t.client.Post("/v5/order/amend", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	var response AmendOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) CancelOrder(req *CancelOrderRequest) (*CancelOrderResponse, error) {
	params := ConvertCancelOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response CancelOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetOpenOrders(req *GetOpenOrdersRequest) (*GetOpenOrdersResponse, error) {
	queryParams := ConvertGetOpenOrdersRequestToParams(req)

	// Assuming the client.Get method constructs the query string from the provided params and sends a GET request.
	resBytes, err := t.client.Get("/v5/order/realtime", queryParams)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response GetOpenOrdersResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) CancelAllOrders(req *CancelAllOrdersRequest) (*CancelAllOrdersResponse, error) {
	params := ConvertCancelAllOrdersRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel-all", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response CancelAllOrdersResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetOrderHistory(req *GetOrderHistoryRequest) (*GetOrderHistoryResponse, error) {
	queryParams := ConvertGetOrderHistoryRequestToParams(req)

	// Constructing the endpoint URL
	endpoint := "https://api.bybit.com/v5/order/history"

	// Converting queryParams to map[string]string for cURL command generation
	paramMap := make(map[string]string)
	for k, v := range queryParams {
		if v != nil {
			paramMap[k] = fmt.Sprintf("%v", v)
		}
	}

	// Printing the parameters being sent
	log.Infof("Sending request to %s with parameters: %+v\n", endpoint, queryParams)

	resBytes, err := t.client.Get(endpoint, queryParams)
	if err != nil {
		// Print the cURL command if there's an error
		curlCommand := generateCurlCommand(endpoint, paramMap)
		log.Errorf("Error in GET request: %v\ncURL Command: %s\n", err, curlCommand)
		return nil, err
	}

	data, err := json.Marshal(resBytes)
	if err != nil {
		// Print the cURL command if there's an error
		curlCommand := generateCurlCommand(endpoint, paramMap)
		log.Errorf("Error marshaling response: %v\ncURL Command: %s\n", err, curlCommand)
		return nil, err
	}

	// Printing the raw response
	log.Info("Received raw response: %s\n", string(data))

	var response GetOrderHistoryResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		// Print the cURL command if there's an error
		curlCommand := generateCurlCommand(endpoint, paramMap)
		log.Errorf("Error unmarshalling response: %v\ncURL Command: %s\n", err, curlCommand)
		return nil, err
	}

	if response.RetCode != 0 {
		// Print the cURL command if there's an API error
		curlCommand := generateCurlCommand(endpoint, paramMap)
		log.Errorf("API returned error: %s\ncURL Command: %s\n", response.RetMsg, curlCommand)
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	// Printing the successful response
	log.Infof("Successfully retrieved order history: %+v\n", response)
	return &response, nil
}
func (t *tradeImpl) GetTradeHistory(req *GetTradeHistoryRequest) (*GetTradeHistoryResponse, error) {
	queryParams := ConvertGetTradeHistoryRequestToParams(req)

	// Assuming the client.Get method constructs the query string from the provided params and sends a GET request.
	resBytes, err := t.client.Get("/v5/execution/list", queryParams)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response GetTradeHistoryResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) BatchPlaceOrder(req *BatchPlaceOrderRequest) (*BatchPlaceOrderResponse, error) {
	params := ConvertBatchPlaceOrderRequestToParams(req)
	resBytes, err := t.client.Post("/v5/order/create-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchPlaceOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}

func (t *tradeImpl) BatchAmendOrder(req *BatchAmendOrderRequest) (*BatchAmendOrderResponse, error) {
	params := ConvertBatchAmendOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/amend-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchAmendOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) BatchCancelOrder(req *BatchCancelOrderRequest) (*BatchCancelOrderResponse, error) {
	params := ConvertBatchCancelOrderRequestToParams(req)

	resBytes, err := t.client.Post("/v5/order/cancel-batch", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}
	var response BatchCancelOrderResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) GetBorrowQuotaSpot(symbol, side string) (*BorrowQuotaResponse, error) {
	params := client.Params{
		"category": "spot",
		"symbol":   symbol,
		"side":     side,
	}
	resBytes, err := t.client.Get("/v5/order/spot-borrow-check", params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resBytes)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var response BorrowQuotaResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Check for API error
	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}
func (t *tradeImpl) SetDisconnectCancelAll(req *SetDisconnectCancelAllRequest) (*APIResponse, error) {
	dcpRequest := NewDCPParams(req.TimeWindow)

	// Send POST request to the Bybit API
	responseBody, err := t.client.Post("/v5/order/disconnected-cancel-all", dcpRequest)
	if err != nil {
		return nil, fmt.Errorf("error sending request to API: %w", err)
	}
	data, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var response APIResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Check for API error
	if response.RetCode != 0 {
		return &response, fmt.Errorf("API returned error: %s", response.RetMsg)
	}

	return &response, nil
}

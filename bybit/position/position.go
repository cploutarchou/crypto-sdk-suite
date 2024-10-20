package position

import (
	"encoding/json"
	"fmt"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

// Position defines the interface for interacting with position-related operations in the Bybit API.
type Position interface {
	// GetPositionInfo retrieves position information based on the provided parameters.
	// params: RequestParams - the parameters for the position info request.
	// returns: *Response - the response containing position information.
	//          error - an error if the request fails.
	GetPositionInfo(params *RequestParams) (*Response, error)

	// SetLeverage sets the leverage for a position based on the provided request.
	// req: SetLeverageRequest - the request containing leverage settings.
	// returns: *Response - the response after setting the leverage.
	//          error - an error if the request fails.
	SetLeverage(req *SetLeverageRequest) (*Response, error)

	// SwitchMarginMode switches the margin mode (cross or isolated) for a position.
	// req: SwitchMarginModeRequest - the request containing margin mode settings.
	// returns: *Response - the response after switching the margin mode.
	//          error - an error if the request fails.
	SwitchMarginMode(req *SwitchMarginModeRequest) (*Response, error)

	// SetTPSLMode sets the Take Profit/Stop Loss mode for a given symbol.
	// req: SetTPSLModeRequest - the request containing TP/SL mode settings.
	// returns: *Response - the response after setting the TP/SL mode.
	//          error - an error if the request fails.
	SetTPSLMode(req *SetTPSLModeRequest) (*Response, error)

	// SwitchPositionMode switches the position mode for USDT perpetual and Inverse futures.
	// req: SwitchPositionModeRequest - the request containing position mode settings.
	// returns: *Response - the response after switching the position mode.
	//          error - an error if the request fails.
	SwitchPositionMode(req *SwitchPositionModeRequest) (*Response, error)

	// SetRiskLimit sets the risk limit for a specific symbol.
	// req: SetRiskLimitRequest - the request containing risk limit settings.
	// returns: *Response - the response after setting the risk limit.
	//          error - an error if the request fails.
	SetRiskLimit(req *SetRiskLimitRequest) (*Response, error)

	// SetTradingStop sets take profit, stop loss, or trailing stop for the position.
	// req: SetTradingStopRequest - the request containing trading stop settings.
	// returns: *Response - the response after setting the trading stop.
	//          error - an error if the request fails.
	SetTradingStop(req *SetTradingStopRequest) (*Response, error)

	// SetAutoAddMargin toggles auto-add-margin for an isolated margin position.
	// req: SetAutoAddMarginRequest - the request containing auto-add-margin settings.
	// returns: *Response - the response after setting auto-add-margin.
	//          error - an error if the request fails.
	SetAutoAddMargin(req *SetAutoAddMarginRequest) (*Response, error)

	// AddOrReduceMargin manually adds or reduces margin for an isolated margin position.
	// req: AddReduceMarginRequest - the request containing add/reduce margin settings.
	// returns: *Response - the response after adding or reducing margin.
	//          error - an error if the request fails.
	AddOrReduceMargin(req *AddReduceMarginRequest) (*Response, error)

	// MovePositions transfers positions between UIDs.
	// req: MovePositionRequest - the request containing move position settings.
	// returns: *MovePositionResponse - the response after moving positions.
	//          error - an error if the request fails.
	MovePositions(req *MovePositionRequest) (*MovePositionResponse, error)

	// GetMovePositionHistory queries the history of moved positions.
	// req: GetMovePositionHistoryRequest - the request containing query parameters for move position history.
	// returns: *GetMovePositionHistoryResponse - the response containing the move position history.
	//          error - an error if the request fails.
	GetMovePositionHistory(req *GetMovePositionHistoryRequest) (*GetMovePositionHistoryResponse, error)

	// ConfirmNewRiskLimit confirms the new risk limit for a position, removing the reduceOnly mark if successful.
	// req: ConfirmNewRiskLimitRequest - the request containing new risk limit settings.
	// returns: *Response - the response after confirming the new risk limit.
	//          error - an error if the request fails.
	ConfirmNewRiskLimit(req *ConfirmNewRiskLimitRequest) (*Response, error)
	GetClosedPnLup2Years(req *GetClosedPnLRequest) (*ClosedPnLResponse, error)
}
type impl struct {
	client *client.Client
}

// New creates a new instance of the Position interface, which can be used to interact with the Bybit API.
func New(c *client.Client) Position {
	return &impl{client: c}
}

// GetPositionInfo fetches position information from Bybit.
func (i *impl) GetPositionInfo(params *RequestParams) (*Response, error) {
	requestParams := ConvertPositionRequestParams(params)
	response, err := i.client.Get("/v5/position/list", requestParams)
	if err != nil {
		return nil, fmt.Errorf("error fetching position info: %w", err)
	}

	// Parse the JSON response
	var positionResponse Response
	err = response.Unmarshal(&positionResponse)
	if err != nil {
		return nil, err
	}

	return &positionResponse, nil
}

// SetLeverage sets the leverage for a given symbol and account type.
func (i *impl) SetLeverage(req *SetLeverageRequest) (*Response, error) {
	params := ConvertSetLeverageRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/set-leverage", params)
	if err != nil {
		return nil, fmt.Errorf("error setting leverage: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var apiResponse Response
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	if apiResponse.RetCode != 0 {
		return nil, fmt.Errorf("API returned error: %s", apiResponse.RetMsg)
	}

	return &apiResponse, nil
}

// SwitchMarginMode switches between cross-margin mode and isolated margin mode for a symbol.
func (i *impl) SwitchMarginMode(req *SwitchMarginModeRequest) (*Response, error) {
	// Convert payload to Params type expected by the client.Post method
	params := ConvertSwitchMarginModeRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/switch-isolated", params)
	if err != nil {
		return nil, fmt.Errorf("error switching margin mode: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Optionally, check the response.RetCode here and handle any errors
	var apiResponse Response
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	if apiResponse.RetCode != 0 {
		return nil, fmt.Errorf("API returned error: %s", apiResponse.RetMsg)
	}

	return &apiResponse, nil
}
func (i *impl) SetTPSLMode(req *SetTPSLModeRequest) (*Response, error) {
	params := ConvertSetTPSLModeRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/set-tpsl-mode", params)
	if err != nil {
		return nil, fmt.Errorf("error setting TP/SL mode: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing TP/SL mode response: %w", err)
	}

	return &positionResponse, nil
}
func (i *impl) SwitchPositionMode(req *SwitchPositionModeRequest) (*Response, error) {
	params := ConvertSwitchPositionModeRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/switch-mode", params)
	if err != nil {
		return nil, fmt.Errorf("error switching position mode: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing switch position mode response: %w", err)
	}
	return &positionResponse, nil
}

func (i *impl) SetRiskLimit(req *SetRiskLimitRequest) (*Response, error) {
	params := ConvertSetRiskLimitRequestToParams(req)

	// Perform the POST request
	response, err := i.client.Post("/v5/position/set-risk-limit", params)
	if err != nil {
		return nil, fmt.Errorf("error setting risk limit: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing set risk limit response: %w", err)
	}

	return &positionResponse, nil
}

func (i *impl) SetTradingStop(req *SetTradingStopRequest) (*Response, error) {
	params := ConvertSetTradingStopRequestToParams(req)

	response, err := i.client.Post("/v5/position/trading-stop", params)
	if err != nil {
		return nil, fmt.Errorf("error setting trading stop: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing set trading stop response: %w", err)
	}

	return &positionResponse, nil
}
func (i *impl) SetAutoAddMargin(req *SetAutoAddMarginRequest) (*Response, error) {
	params := ConvertSetAutoAddMarginRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/set-auto-add-margin", params)
	if err != nil {
		return nil, fmt.Errorf("error setting auto add margin: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing set auto add margin response: %w", err)
	}

	return &positionResponse, nil
}
func (i *impl) AddOrReduceMargin(req *AddReduceMarginRequest) (*Response, error) {
	params := ConvertAddReduceMarginRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/add-margin", params)
	if err != nil {
		return nil, fmt.Errorf("error adding or reducing margin: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing add or reduce margin response: %w", err)
	}

	return &positionResponse, nil
}
func (i *impl) GetClosedPnLup2Years(req *GetClosedPnLRequest) (*ClosedPnLResponse, error) {
	params := ConvertGetClosedPnLRequestToParams(req)
	var allRecords []interface{}
	finalResponse := &ClosedPnLResponse{}
	for {
		responseData, err := i.client.Get("/v5/position/closed-pnl", params)
		if err != nil {
			return nil, fmt.Errorf("error fetching closed PnL records: %w", err)
		}
		data, err := json.Marshal(responseData)
		if err != nil {
			return nil, err
		}
		var response ClosedPnLResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, fmt.Errorf("error parsing closed PnL response: %w", err)
		}

		allRecords = append(allRecords, response.Result.List...)

		if response.Result.NextPageCursor == "" {
			break
		}

		params["cursor"] = response.Result.NextPageCursor
		finalResponse = &ClosedPnLResponse{
			RetCode:    response.RetCode,
			RetMsg:     response.RetMsg,
			Result:     response.Result,
			RetExtInfo: response.RetExtInfo,
			Time:       response.Time,
		}
		finalResponse.Result.List = allRecords
	}

	return finalResponse, nil
}
func (i *impl) MovePositions(req *MovePositionRequest) (*MovePositionResponse, error) {
	params := ConvertMovePositionRequestToParams(req)
	// Perform the POST request
	response, err := i.client.Post("/v5/position/move-positions", params)
	if err != nil {
		return nil, fmt.Errorf("error moving positions: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var movePositionResponse MovePositionResponse
	if err := json.Unmarshal(data, &movePositionResponse); err != nil {
		return nil, fmt.Errorf("error parsing move position response: %w", err)
	}

	return &movePositionResponse, nil
}
func (i *impl) GetMovePositionHistory(req *GetMovePositionHistoryRequest) (*GetMovePositionHistoryResponse, error) {
	var allEntries []MovePositionHistoryEntry
	var finalResponse GetMovePositionHistoryResponse

	for {
		// Construct query parameters
		params := ConvertGetMovePositionHistoryRequestToParams(req)

		// Perform the GET request
		response, err := i.client.Get("/v5/position/move-history", params)
		if err != nil {
			return nil, fmt.Errorf("error fetching move position history: %w", err)
		}
		data, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		// Parse the JSON response
		var historyResponse GetMovePositionHistoryResponse
		if err := json.Unmarshal(data, &historyResponse); err != nil {
			return nil, fmt.Errorf("error parsing move position history response: %w", err)
		}

		// Accumulate entries from this page
		allEntries = append(allEntries, historyResponse.Result.List...)

		// Check if there's a next page
		if historyResponse.Result.NextPageCursor == "" {
			break // Exit loop if there's no next page
		}

		// Prepare for the next request
		req.Cursor = &historyResponse.Result.NextPageCursor
		// Prepare the final consolidated response
		finalResponse.RetCode = historyResponse.RetCode
		finalResponse.RetMsg = historyResponse.RetMsg
		finalResponse.Result.NextPageCursor = historyResponse.Result.NextPageCursor
		finalResponse.Result.List = allEntries
		finalResponse.RetExtInfo = historyResponse.RetExtInfo
		finalResponse.Time = historyResponse.Time

	}

	return &finalResponse, nil
}
func (i *impl) ConfirmNewRiskLimit(req *ConfirmNewRiskLimitRequest) (*Response, error) {
	params := ConvertConfirmNewRiskLimitRequestToParams(req)

	// Perform the POST request
	response, err := i.client.Post("/v5/position/confirm-pending-mmr", params)
	if err != nil {
		return nil, fmt.Errorf("error confirming new risk limit: %w", err)
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing confirm new risk limit response: %w", err)
	}

	return &positionResponse, nil
}

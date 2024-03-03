package position

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Position interface {
	GetPositionInfo(params *RequestParams) (*Response, error)
	SetLeverage(req *SetLeverageRequest) (*Response, error)
	SwitchMarginMode(req *SwitchMarginModeRequest) (*Response, error)
	// SetTPSLMode sets the TP/SL mode for a given symbol.
	SetTPSLMode(req *SetTPSLModeRequest) (*Response, error)
	// SwitchPositionMode switches the position mode for USDT perpetual and Inverse futures.
	SwitchPositionMode(req *SwitchPositionModeRequest) (*Response, error)
	// SetRiskLimit sets the risk limit for a specific symbol.
	SetRiskLimit(req *SetRiskLimitRequest) (*Response, error)
	// SetTradingStop sets take profit, stop loss, or trailing stop for the position.
	SetTradingStop(req *SetTradingStopRequest) (*Response, error)
	// SetAutoAddMargin toggles auto-add-margin for an isolated margin position.
	SetAutoAddMargin(req *SetAutoAddMarginRequest) (*Response, error)
	// AddOrReduceMargin manually adds or reduces margin for an isolated margin position.
	AddOrReduceMargin(req *AddReduceMarginRequest) (*Response, error)
}
type impl struct {
	client *client.Client
}

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
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	// Parse the JSON response
	var positionResponse Response
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing position info response: %w", err)
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

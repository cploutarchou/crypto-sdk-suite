package position

import (
	"encoding/json"
	"fmt"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Position interface {
	GetPositionInfo(params *PositionRequestParams) (*PositionResponse, error)
	SetLeverage(req *SetLeverageRequest) (*PositionResponse, error)
	SwitchMarginMode(req *SwitchMarginModeRequest) (*PositionResponse, error)
	// SetTPSLMode sets the TP/SL mode for a given symbol.
	SetTPSLMode(req *SetTPSLModeRequest) (*PositionResponse, error)
	// SwitchPositionMode switches the position mode for USDT perpetual and Inverse futures.
	SwitchPositionMode(req *SwitchPositionModeRequest) (*PositionResponse, error)
	// SetRiskLimit sets the risk limit for a specific symbol.
	SetRiskLimit(req *SetRiskLimitRequest) (*PositionResponse, error)
}
type impl struct {
	client *client.Client
}

func New(c *client.Client) Position {
	return &impl{client: c}
}

// GetPositionInfo fetches position information from Bybit.
func (i *impl) GetPositionInfo(params *PositionRequestParams) (*PositionResponse, error) {
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
	var positionResponse PositionResponse
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing position info response: %w", err)
	}

	return &positionResponse, nil
}

// SetLeverage sets the leverage for a given symbol and account type.
func (i *impl) SetLeverage(req *SetLeverageRequest) (*PositionResponse, error) {
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
	var apiResponse PositionResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	if apiResponse.RetCode != 0 {
		return nil, fmt.Errorf("API returned error: %s", apiResponse.RetMsg)
	}

	return &apiResponse, nil
}

// SwitchMarginMode switches between cross-margin mode and isolated margin mode for a symbol.
func (i *impl) SwitchMarginMode(req *SwitchMarginModeRequest) (*PositionResponse, error) {
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
	var apiResponse PositionResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	if apiResponse.RetCode != 0 {
		return nil, fmt.Errorf("API returned error: %s", apiResponse.RetMsg)
	}

	return &apiResponse, nil
}
func (i *impl) SetTPSLMode(req *SetTPSLModeRequest) (*PositionResponse, error) {
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
	var positionResponse PositionResponse
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing TP/SL mode response: %w", err)
	}

	return &positionResponse, nil
}
func (i *impl) SwitchPositionMode(req *SwitchPositionModeRequest) (*PositionResponse, error) {
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
	var positionResponse PositionResponse
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing switch position mode response: %w", err)
	}
	return &positionResponse, nil
}

func (i *impl) SetRiskLimit(req *SetRiskLimitRequest) (*PositionResponse, error) {
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
	var positionResponse PositionResponse
	if err := json.Unmarshal(data, &positionResponse); err != nil {
		return nil, fmt.Errorf("error parsing set risk limit response: %w", err)
	}

	return &positionResponse, nil
}

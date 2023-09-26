package account

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

const (
	setMarginModePath = "/v5/account/set-margin-mode"
	resetMMPPath      = "/v5/account/mmp-reset"
	setMMPPath        = "/v5/account/mmp-modify"
	getMMPStatePath   = "/v5/account/mmp-state"
)

type Margin struct {
	client *client.Client
}

func NewMargin(client *client.Client) *Margin {
	return &Margin{client: client}
}

func (m *Margin) SetMarginMode(mode string) (*SetMarginModeResponse, error) {
	params := client.Params{
		"setMarginMode": mode,
	}

	response, err := m.client.Post(setMarginModePath, params)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", response.StatusCode())
	}

	var setMarginModeResponse SetMarginModeResponse
	err = response.Unmarshal(&setMarginModeResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", response.StatusCode(), response.Status())
	}

	return &setMarginModeResponse, nil
}

// SetMMP sets the Market Maker Protection for the client.
func (m *Margin) SetMMP(params *MMPParams) (*MMPResponse, error) {
	response, err := m.client.Post(setMMPPath, client.Params{
		"baseCoin":     params.BaseCoin,
		"window":       strconv.Itoa(params.Window),
		"frozenPeriod": strconv.Itoa(params.FrozenPeriod),
		"qtyLimit":     strconv.Itoa(params.QtyLimit),
		"deltaLimit":   strconv.Itoa(params.DeltaLimit),
	})
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", response.StatusCode())
	}

	var mmpResponse MMPResponse
	err = response.Unmarshal(&mmpResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if mmpResponse.RetCode != 0 {
		return nil, fmt.Errorf("unexpected retCode: %d, retMsg: %s", mmpResponse.RetCode, mmpResponse.RetMsg)
	}

	return &mmpResponse, nil
}

func (m *Margin) ResetMMP(baseCoin string) (*MMPResponse, error) {
	params := client.Params{
		"baseCoin": baseCoin,
	}

	response, err := m.client.Post(resetMMPPath, params)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", response.StatusCode())
	}

	var mmpResponse MMPResponse
	err = response.Unmarshal(&mmpResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if mmpResponse.RetCode != 0 {
		return nil, fmt.Errorf("unexpected retCode: %d, retMsg: %s", mmpResponse.RetCode, mmpResponse.RetMsg)
	}

	return &mmpResponse, nil
}

func (m *Margin) GetMMPState(baseCoin string) (*MMPStateResponse, error) {
	params := client.Params{
		"baseCoin": baseCoin,
	}

	response, err := m.client.Get(getMMPStatePath, params)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", response.StatusCode())
	}

	var mmpStateResponse MMPStateResponse
	err = response.Unmarshal(&mmpStateResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if mmpStateResponse.RetCode != 0 {
		return nil, fmt.Errorf("unexpected retCode: %d, retMsg: %s", mmpStateResponse.RetCode, mmpStateResponse.RetMsg)
	}

	return &mmpStateResponse, nil
}

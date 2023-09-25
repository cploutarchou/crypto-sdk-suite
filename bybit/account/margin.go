package account

import (
	"fmt"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

const setMarginModePath = "/v5/account/set-margin-mode"

type Margin struct {
	client *client.Client
}

func NewMargin(client *client.Client) *Margin {
	return &Margin{client: client}
}

type SetMarginModeResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Reasons []struct {
			ReasonCode string `json:"reasonCode"`
			ReasonMsg  string `json:"reasonMsg"`
		} `json:"reasons"`
	} `json:"result"`
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

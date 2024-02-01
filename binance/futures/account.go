package futures

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/client"
)

// Endpoints used in the futures package.
const (
	changePositionModeEndpoint = "/fapi/v1/positionSide/dual"
)

// Account defines the interface for account operations.
type Account interface {
	ChangePositionMode(enable bool) error
}

// accountImpl implements Account interface using Binance futures API.
type accountImpl struct {
	client *client.Client
}

// NewAccount creates a new Account instance.
func NewAccount(client *client.Client) Account {
	return &accountImpl{
		client: client,
	}
}

// ChangePositionMode toggles the position mode for a futures account.
func (a *accountImpl) ChangePositionMode(enable bool) error {
	endpoint := changePositionModeEndpoint
	data := fmt.Sprintf("dualSidePosition=%v", enable)

	resp, err := a.client.MakeAuthenticatedRequest(http.MethodPost, endpoint, data)
	if err != nil {
		return fmt.Errorf("error making authenticated request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, responseBody)
	}

	return nil
}

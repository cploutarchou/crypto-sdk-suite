package futures

import (
	"fmt"
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
	*client.Client
}

// NewAccount creates a new Account instance.
func NewAccount(client *client.Client) Account {
	return &accountImpl{
		client,
	}
}

// ChangePositionMode toggles the position mode for a futures account.
func (a *accountImpl) ChangePositionMode(enable bool) error {
	endpoint := changePositionModeEndpoint
	data := fmt.Sprintf("dualSidePosition=%v", enable)
	var resp any
	err := a.MakeAuthenticatedRequest(http.MethodPost, endpoint, data, resp)
	if err != nil {
		return fmt.Errorf("failed to change position mode: %w", err)
	}
	return nil
}

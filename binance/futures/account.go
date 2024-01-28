package binance

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/binance/futures/requests"
)

const (
	changePositionModeEndpoint = "/fapi/v1/positionSide/dual"
)

type Account struct {
	*requests.Client
}

func NewAccount(client *requests.Client) *Account {
	return &Account{
		Client: client,
	}
}
func (b *Account) ChangePositionMode(enable bool) error {
	endpoint := changePositionModeEndpoint
	data := fmt.Sprintf("dualSidePosition=%v", enable)

	resp, err := b.MakeAuthenticatedRequest(http.MethodPost, endpoint, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d, Response: %s", resp.StatusCode, string(responseBody))
	}

	return nil
}

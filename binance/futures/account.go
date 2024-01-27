package binance

import (
	"fmt"
	"io"
	"net/http"
)

const (
	changePositionModeEndpoint = "/fapi/v1/positionSide/dual"
)

func (b *BinanceClient) ChangePositionMode(enable bool) error {
	endpoint := changePositionModeEndpoint
	data := fmt.Sprintf("dualSidePosition=%v", enable)

	resp, err := b.makeAuthenticatedRequest(http.MethodPost, endpoint, data)
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

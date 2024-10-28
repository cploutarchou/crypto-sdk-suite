package account

import (
	"errors"
	"net/http"

	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
)

type Info struct {
	client *client.Client
}

func NewInfo(client *client.Client) *Info {
	return &Info{client: client}
}

// Get queries the margin mode configuration of the account.
func (info *Info) Get() (*AccInfo, error) {
	path := "/v5/account/info"
	resp, err := info.client.Get(path, nil) // Assuming the Get method is as per your client package.

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to get account info: non-200 status code received")
	}

	var accountInfo AccInfo
	err = resp.Unmarshal(&accountInfo)
	if err != nil {
		return nil, err
	} // Assuming the Unmarshal method is as per your client package.
	return &accountInfo, nil
}

package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type UpgradeToUnified struct {
	client *client.Client
}

func NewUpgradeToUnifiedRequest(c *client.Client) *UpgradeToUnified {
	return &UpgradeToUnified{c}
}

func (r *UpgradeToUnified) Upgrade() (*UpgradeToUnifiedResponse, error) {
	var ret UpgradeToUnifiedResponse
	res, err := r.client.Post(Endpoints.UpgradeToUnified, client.Params{})
	if err != nil {
		return nil, err
	}
	err = res.Unmarshal(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

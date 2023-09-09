package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

const (
	upgradeToUnified string = "/v5/account/upgrade-to-uta"
)

type UpgradeToUnifiedResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		UnifiedUpdateStatus string `json:"unifiedUpdateStatus"`
		UnifiedUpdateMsg    struct {
			Msg []string `json:"msg"`
		} `json:"unifiedUpdateMsg"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type UpgradeToUnified struct {
	client *client.Client
}

func (r *UpgradeToUnified) Upgrade() (*UpgradeToUnifiedResponse, error) {
	var ret UpgradeToUnifiedResponse
	res, err := r.client.Post(upgradeToUnified, client.Params{})
	if err != nil {
		return nil, err
	}
	err = res.Unmarshal(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
func NewUpgradeToUnifiedRequest(c *client.Client) *UpgradeToUnified {
	return &UpgradeToUnified{c}
}

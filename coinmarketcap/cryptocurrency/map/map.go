package idmap

import (
	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	"strconv"
)

type Map struct {
	client *client.Client
}

func (m *Map) GetID(params *Params) ([]Data, error) {
	path := "/v1/cryptocurrency/map"
	var queryParams map[string]string
	if params != nil {
		queryParams = map[string]string{
			"listing_status": params.ListingStatus,
			"start":          strconv.Itoa(params.Start),
			"limit":          strconv.Itoa(params.Limit),
			"sort":           params.Sort,
			"symbol":         params.Symbol,
			"aux":            params.Aux,
		}
	}

	resp, err := m.client.Get(path, queryParams)
	if err != nil {
		return nil, err
	}

	var idMapResponse Response
	if resp.StatusCode() != 200 {
		return nil, resp.Error()
	}

	err = resp.Unmarshal(&idMapResponse)
	if err != nil {
		return nil, err
	}
	return idMapResponse.Data, nil
}

func New(c *client.Client) *Map {
	return &Map{
		client: c,
	}
}

package idmap

import (
	"strconv"

	"github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
)

type Params struct {
	ListingStatus string
	Start         int
	Limit         int
	Sort          string
	Symbol        string
	Aux           string
}

type Map struct {
	*client.Client
}

func (m *Map) GetID(params *Params) ([]Data, error) {
	path := "/v1/cryptocurrency/map"
	queryParams := make(map[string]string)

	if params != nil {
		if params.ListingStatus != "" {
			queryParams["listing_status"] = params.ListingStatus
		}
		if params.Start != 0 {
			queryParams["start"] = strconv.Itoa(params.Start)
		}
		if params.Limit != 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Sort != "" {
			queryParams["sort"] = params.Sort
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Aux != "" {
			queryParams["aux"] = params.Aux
		}
	}

	resp, err := m.Get(path, queryParams)
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
		c,
	}
}

package gainer_looser

import (
	"fmt"
	c "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
	"strconv"
)

// GainersAndLosers represents the gainers and losers data.
type GainersAndLosers struct {
	*c.Client
}

// Sort represents the fields by which data should be sorted.
type Sort string

// SortDir represents the direction in which data should be sorted.
type SortDir string

// TimePeriod represents the overall window of time for the gainers and losers.
type TimePeriod string

const (
	PercentChange1h       Sort       = "percent_change_1h"
	PercentChange24h      Sort       = "percent_change_24h"
	PercentChange7d       Sort       = "percent_change_7d"
	PercentChange30d      Sort       = "percent_change_30d"
	ASC                   SortDir    = "asc"
	DESC                  SortDir    = "desc"
	OneHour               TimePeriod = "1h"
	TwentyFourHour        TimePeriod = "24h"
	ThirtyDay             TimePeriod = "30d"
	SevenDay              TimePeriod = "7d"
	GainersLosersEndpoint            = "/v1/cryptocurrency/trending/gainers-losers"
)

func New(c *c.Client) *GainersAndLosers {
	return &GainersAndLosers{
		c,
	}
}

// IsValidSort checks if the given Sort is a valid enumeration value.
func (g *GainersAndLosers) isValidSort(s Sort) bool {
	switch s {
	case PercentChange1h, PercentChange24h, PercentChange7d, PercentChange30d:
		return true
	default:
		return false
	}
}

// IsValidSortDir checks if the given SortDir is a valid enumeration value.
func (g *GainersAndLosers) isValidSortDir(dir SortDir) bool {
	switch dir {
	case ASC, DESC:
		return true
	default:
		return false
	}
}

// IsValidTimePeriod checks if the given TimePeriod is a valid enumeration value.
func (g *GainersAndLosers) isValidTimePeriod(tp TimePeriod) bool {
	switch tp {
	case OneHour, TwentyFourHour, ThirtyDay, SevenDay:
		return true
	default:
		return false
	}
}

func (g *GainersAndLosers) FetchGainersLosers(params *Params) ([]CryptocurrencyData, error) {
	queryParams := make(map[string]string)
	if params == nil {
		params = &Params{}
	}
	if params.Limit != nil {
		queryParams["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.TimePeriod != nil && g.isValidTimePeriod(*params.TimePeriod) {
		queryParams["time_period"] = string(*params.TimePeriod)
	}
	if params.Convert != nil {
		queryParams["convert"] = *params.Convert
	}
	if params.ConvertID != nil {
		queryParams["convert_id"] = strconv.Itoa(*params.ConvertID)
	}
	if params.Sort != nil && g.isValidSort(*params.Sort) {
		queryParams["sort"] = string(*params.Sort)
	}
	if params.SortDir != nil && g.isValidSortDir(*params.SortDir) {
		queryParams["sort_dir"] = string(*params.SortDir)
	}

	resp, err := g.Get(GainersLosersEndpoint, queryParams)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(resp.Data()))
	var data Response
	if err := resp.Unmarshal(&data); err != nil {
		return nil, err
	}

	return data.Data.Data, nil
}

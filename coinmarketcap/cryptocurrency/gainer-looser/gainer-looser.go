package gainer_looser

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/client"
	"strconv"
)

// GainersAndLosers represents the gainers and losers data.
type GainersAndLosers struct {
	*client.Client
}

// Sort represents the fields by which data should be sorted.
type Sort string

// SortDir represents the direction in which data should be sorted.
type SortDir string

// TimePeriod represents the overall window of time for the gainers and losers.
type TimePeriod string

const (
	PercentChange1h  Sort       = "percent_change_1h"
	PercentChange24h Sort       = "percent_change_24h"
	PercentChange7d  Sort       = "percent_change_7d"
	PercentChange30d Sort       = "percent_change_30d"
	ASC              SortDir    = "asc"
	DESC             SortDir    = "desc"
	OneHour          TimePeriod = "1h"
	TwentyFourHour   TimePeriod = "24h"
	ThirtyDay        TimePeriod = "30d"
	SevenDay         TimePeriod = "7d"
)

func New(c *client.Client) *GainersAndLosers {
	return &GainersAndLosers{
		c,
	}
}

// IsValidSort checks if the given Sort is a valid enumeration value.
func IsValidSort(s Sort) bool {
	switch s {
	case PercentChange1h, PercentChange24h, PercentChange7d, PercentChange30d:
		return true
	default:
		return false
	}
}

// IsValidSortDir checks if the given SortDir is a valid enumeration value.
func IsValidSortDir(dir SortDir) bool {
	switch dir {
	case ASC, DESC:
		return true
	default:
		return false
	}
}

// IsValidTimePeriod checks if the given TimePeriod is a valid enumeration value.
func IsValidTimePeriod(tp TimePeriod) bool {
	switch tp {
	case OneHour, TwentyFourHour, ThirtyDay, SevenDay:
		return true
	default:
		return false
	}
}

const GainersLosersEndpoint = "/v1/cryptocurrency/trending/gainers-losers"

func (g *GainersAndLosers) FetchGainersLosers(cli client.Requester, params *Params) ([]CryptocurrencyData, error) {
	queryParams := make(map[string]string)

	if params.Limit != nil {
		queryParams["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.TimePeriod != nil && IsValidTimePeriod(*params.TimePeriod) {
		queryParams["time_period"] = string(*params.TimePeriod)
	}
	if params.Convert != nil {
		queryParams["convert"] = *params.Convert
	}
	if params.ConvertID != nil {
		queryParams["convert_id"] = strconv.Itoa(*params.ConvertID)
	}
	if params.Sort != nil && IsValidSort(*params.Sort) {
		queryParams["sort"] = string(*params.Sort)
	}
	if params.SortDir != nil && IsValidSortDir(*params.SortDir) {
		queryParams["sort_dir"] = string(*params.SortDir)
	}

	resp, err := cli.Get(GainersLosersEndpoint, queryParams)
	if err != nil {
		return nil, err
	}

	var data Response
	if err := resp.Unmarshal(&data); err != nil {
		return nil, err
	}

	return data.Data, nil
}

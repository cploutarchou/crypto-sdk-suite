package gainer_looser

// GainersAndLosers represents the gainers and losers data.
type GainersAndLosers struct {
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

// Params represents the query parameters for fetching data.
type Params struct {
	Start      *int        // Start represents the starting index.
	Limit      *int        // The Limit represents the maximum number of results to retrieve.
	TimePeriod *TimePeriod // TimePeriod represents the time span for the data.
	Convert    *string     // Convert represents the currency conversion symbols.
	ConvertID  *string     // ConvertID represents the CoinMarketCap ID for conversion.
	Sort       *Sort       // Sort specifies the field to sort data by.
	SortDir    *SortDir    // SortDir specifies the direction of the sort.
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

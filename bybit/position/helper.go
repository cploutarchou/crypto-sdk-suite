package position

import "strconv"

// PreparePositionRequestParams prepares the request parameters for fetching position info.
func PreparePositionRequestParams(category string, symbol string, limit int) Params {
	params := make(Params)
	params["category"] = category
	if symbol != "" {
		params["symbol"] = symbol
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit) // Convert int to string
	}
	// Add more parameters as needed
	return params
}

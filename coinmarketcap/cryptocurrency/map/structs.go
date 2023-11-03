package idmap

type Data struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	Platform            string `json:"platform"`
	FirstHistoricalData string `json:"first_historical_data"`
	LastHistoricalData  string `json:"last_historical_data"`
	IsActive            int    `json:"is_active"`
	Status              string `json:"status"`
}

type Response struct {
	Data   []Data `json:"data"`
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
}

type Params struct {
	ListingStatus string
	Start         int
	Limit         int
	Sort          string
	Symbol        string
	Aux           string
}

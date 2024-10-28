package lt_kline

// LTKlineData represents the data structure for LT Kline.
type LTKlineData struct {
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Confirm   bool   `json:"confirm"`
	Timestamp int64  `json:"timestamp"`
}

// LTKlineResponse represents the response structure for LT Kline.
type LTKlineResponse struct {
	Topic string        `json:"topic"`
	Type  string        `json:"type"`
	TS    int64         `json:"ts"`
	Data  []LTKlineData `json:"data"`
}

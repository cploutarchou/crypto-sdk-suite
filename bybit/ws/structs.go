package ws

type Kline struct {
	Topic string      `json:"topic"`
	Data  []KlineData `json:"data"`
	Ts    int64       `json:"ts"`
	Type  string      `json:"type"`
}

type KlineData struct {
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Interval  string `json:"interval"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	Turnover  string `json:"turnover"`
	Confirm   bool   `json:"confirm"`
	Timestamp int64  `json:"timestamp"`
}

package kline

type Response struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  []Data `json:"data"`
	Ts    int64  `json:"ts"`
}
type Data struct {
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

package ws

type Environment string // Environment is the environment for the Bybit API
type ChannelType string // ChannelType is the channel type for the Bybit API
type SubChannel string  // SubChannel is the sub channel for the Bybit API
// SuccessResponse is the response for a successful request
type SuccessResponse struct {
	Success bool   `json:"success"` // Success is the success status of the request
	RetMsg  string `json:"ret_msg"` // RetMsg is the return message of the request
	Op      string `json:"op"`      // Op is the operation of the request
	ConnId  string `json:"conn_id"` // ConnId is the connection ID of the request
}

// PingMsg is the message for a ping request
type PingMsg struct {
	ReqId string `json:"req_id"` // ReqId is the request ID of the ping request
	Op    string `json:"op"`     // Op is the operation of the ping request
}

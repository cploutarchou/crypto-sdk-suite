package client

// WSPongPublicResponse contains fields common to public pong responses
type WSPongPublicResponse struct {
	SuccessResponse
}

// WSPongPrivateResponse contains fields common to private pong responses
type WSPongPrivateResponse struct {
	CommonResponse
	ReqId string   `json:"req_id"`
	Args  []string `json:"args"`
}

// Now, use the above base types to create specific responses

type WSPongPublicSpotResponse WSPongPublicResponse
type WSPongPrivateSpotResponse WSPongPrivateResponse

type WSPongPublicLinearInverseResponse struct {
	WSPongPublicResponse
	ReqId string `json:"req_id"`
}

type WSPongPrivateLinearInverseResponse WSPongPrivateResponse

type SoppingOptionsResponse struct {
	Args []string `json:"args"`
	Op   string   `json:"op"`
}

const (
	MainNet Environment = "stream.bybit.com"
	TestNet Environment = "stream-testnet.bybit.com"

	Public  ChannelType = "public"
	Private ChannelType = "private"
)

type CommonResponse struct {
	Op     string `json:"op"`      // Op is the operation of the request
	ConnId string `json:"conn_id"` // ConnId is the connection ID of the request
}

// SuccessResponse is the response for a successful request
type SuccessResponse struct {
	CommonResponse
	Success bool   `json:"success"` // Success is the success status of the request
	RetMsg  string `json:"ret_msg"` // RetMsg is the return message of the request
	Op      string `json:"op"`      // Op is the operation of the request
	ConnId  string `json:"conn_id"` // ConnId is the connection ID of the request
}

type Environment string // Environment is the environment for the Bybit API
type ChannelType string // ChannelType is the channel type for the Bybit API
type SubChannel string  // SubChannel is the sub channel for the Bybit API

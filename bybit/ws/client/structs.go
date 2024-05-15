package client

// Package client provides data structures and constants to manage
// WebSocket responses from a server. It also includes types that
// handle various response formats for both public and private messages.

import (
	"github.com/gorilla/websocket"
)

// WSPongPublicResponse contains fields common to public pong responses.
type WSPongPublicResponse struct {
	SuccessResponse
}

// WSPongPrivateResponse contains fields common to private pong responses.
// It also provides a unique request ID and arguments associated with the response.
type WSPongPrivateResponse struct {
	CommonResponse
	ReqId string   `json:"req_id"`
	Args  []string `json:"args"`
}

// WSPongPublicSpotResponse represents a specific public pong response for the spot.
type WSPongPublicSpotResponse WSPongPublicResponse

// WSPongPrivateSpotResponse represents a specific private pong response for the spot.
type WSPongPrivateSpotResponse WSPongPrivateResponse

// WSPongPublicLinearInverseResponse represents a specific public pong response
// for linear inverse operations and also contains a unique request ID.
type WSPongPublicLinearInverseResponse struct {
	WSPongPublicResponse
	ReqId string `json:"req_id"`
}

// WSPongPrivateLinearInverseResponse represents a specific private pong response
// for linear inverse operations.
type WSPongPrivateLinearInverseResponse WSPongPrivateResponse

// SoppingOptionsResponse describes the sopping options with associated arguments and operations.
type SoppingOptionsResponse struct {
	Args []string `json:"args"`
	Op   string   `json:"op"`
}

const (
	// WSMessageText represents the message type for text in WebSocket communications.
	WSMessageText = websocket.TextMessage
)

// CommonResponse encapsulates fields that are common across various WebSocket responses.
type CommonResponse struct {
	Op     string `json:"op"`      // Op is the operation of the request
	ConnId string `json:"conn_id"` // ConnId is the connection ID of the request
}

// SuccessResponse represents the response structure for a successful WebSocket request.
// It provides details about the operation, connection ID, success status, and return message.
type SuccessResponse struct {
	CommonResponse
	Success bool   `json:"success"` // Success indicates the success status of the request
	RetMsg  string `json:"ret_msg"` // RetMsg provides details on the return message of the request
}

// Environment denotes the environment in which the Bybit API operates (e.g., "production", "development").
type Environment string

// SubChannel represents a sub-channel for Bybit API's WebSocket communications.
type SubChannel string

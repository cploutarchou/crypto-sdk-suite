package client

import (
	"encoding/json"
)

type Response interface {
	Unmarshal(v interface{}) error
	Data() []byte
}

type ResponseImpl struct {
	data []byte
	err  error
}

func NewResponse(data []byte) Response {
	return &ResponseImpl{
		data: data,
	}
}

func (r *ResponseImpl) Unmarshal(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal(r.Data(), v)
}

func (r *ResponseImpl) Data() []byte {
	return r.data
}

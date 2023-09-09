package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response interface {
	Unmarshal(v interface{}) error
	Data() []byte
	Status() string
	StatusCode() int
}

type ResponseImpl struct {
	data       []byte
	err        error
	statusCode int
	status     string
}

func NewResponse(response *http.Response) Response {
	var res ResponseImpl
	body, err := io.ReadAll(response.Body)
	if err != nil {
		res.err = err
	}
	res.statusCode = response.StatusCode
	res.data = body
	res.status = response.Status
	return &res
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

func (r *ResponseImpl) StatusCode() int {
	return r.statusCode
}

func (r *ResponseImpl) Status() string {
	return r.status
}

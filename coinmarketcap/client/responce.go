package client

import (
	"encoding/json"
	"io"
	"net/http"

	httperrors "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/errors"
)

type Response interface {
	Unmarshal(v any) error
	Data() []byte
	Status() string
	StatusCode() int
	Error() error
}

type ResponseImpl struct {
	data       []byte
	err        error
	statusCode int
	status     string
	httpError  httperrors.HTTPError
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

	// Check if the status code indicates an error
	switch res.statusCode {
	case http.StatusBadRequest:
		res.httpError = httperrors.BadRequest(string(body))
	case http.StatusUnauthorized:
		res.httpError = httperrors.Unauthorized(string(body))
	case http.StatusForbidden:
		res.httpError = httperrors.Forbidden(string(body))
	case http.StatusTooManyRequests:
		res.httpError = httperrors.TooManyRequests(string(body))
	case http.StatusInternalServerError:
		res.httpError = httperrors.InternalServerError(string(body))
	}

	return &res
}

func (r *ResponseImpl) Error() error {
	if r.httpError != nil {
		return r.httpError
	}
	return r.err
}

func (r *ResponseImpl) Unmarshal(v any) error {
	// Check if there's an HTTP error first
	if r.httpError != nil {
		return r.httpError
	}
	// If there's a read error, return that next
	if r.err != nil {
		return r.err
	}
	// Unmarshal the data into the provided interface
	return json.Unmarshal(r.data, v)
}

func (r *ResponseImpl) Data() []byte {
	// Return the response data
	return r.data
}

func (r *ResponseImpl) Status() string {
	// Return the HTTP status text
	return r.status
}

func (r *ResponseImpl) StatusCode() int {
	// Return the HTTP status code
	return r.statusCode
}

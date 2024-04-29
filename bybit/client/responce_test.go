package client

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	res := &ResponseImpl{data: []byte(`{"key": "value"}`)}
	var v map[string]string
	err := res.Unmarshal(&v)

	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}
	if v["key"] != "value" {
		t.Error("Unmarshal did not parse data correctly")
	}
}

func TestData(t *testing.T) {
	res := &ResponseImpl{data: []byte("test data")}
	if string(res.Data()) != "test data" {
		t.Error("Data did not return expected value")
	}
}

func TestStatusCode(t *testing.T) {
	res := &ResponseImpl{statusCode: 200}
	if res.StatusCode() != 200 {
		t.Error("StatusCode did not return expected value")
	}
}

func TestStatus(t *testing.T) {
	res := &ResponseImpl{status: "OK"}
	if res.Status() != "OK" {
		t.Error("Status did not return expected value")
	}
}

func TestError(t *testing.T) {
	res := &ResponseImpl{err: http.ErrShortBody}
	if !errors.Is(res.Error(), http.ErrShortBody) {
		t.Error("Error did not return expected value")
	}
}

func TestNewResponse(t *testing.T) {
	body := strings.NewReader("test response body")
	response := &http.Response{
		Body:   io.NopCloser(body),
		Header: make(http.Header),
	}
	response.StatusCode = 200
	response.Status = "OK"

	res := NewResponse(response).(*ResponseImpl)

	if string(res.Data()) != "test response body" {
		t.Error("NewResponse did not set data correctly")
	}
	if res.StatusCode() != 200 {
		t.Error("NewResponse did not set statusCode correctly")
	}
	if res.Status() != "OK" {
		t.Error("NewResponse did not set status correctly")
	}
}

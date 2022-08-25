package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

type MockClient struct{}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: sCode,
		Body:       r,
	}, resErr
}

func SetMock(data string, statusCode int, err error) {
	r = io.NopCloser(bytes.NewReader([]byte(data)))
	sCode = statusCode
	resErr = err
}

var (
	sCode  int
	r      io.ReadCloser
	resErr error
)

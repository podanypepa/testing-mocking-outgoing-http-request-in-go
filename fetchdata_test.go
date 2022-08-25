package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/podanypepa/testing-mocking-outgoing-http-request-in-go/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestFetchData(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		httpclient.Client = &httpclient.MockClient{}
		httpclient.SetMock("pepa", http.StatusOK, nil)

		res, err := fetchData("")

		assert.NoError(t, err)
		assert.Equal(t, "pepa", res)
	})

	t.Run("http error", func(t *testing.T) {
		httpclient.Client = &httpclient.MockClient{}
		httpclient.SetMock("pepa", http.StatusOK, fmt.Errorf("http error"))

		res, err := fetchData("")

		assert.Error(t, err)
		assert.Equal(t, "", res)
	})

	t.Run("wrong status code", func(t *testing.T) {
		httpclient.Client = &httpclient.MockClient{}
		httpclient.SetMock("pepa", 400, nil)

		res, err := fetchData("")

		assert.Error(t, err)
		assert.Equal(t, "", res)
	})

	t.Run("wrong res body", func(t *testing.T) {
		httpclient.Client = &httpclient.MockClient{}
		httpclient.SetMock("", 200, nil)

		res, err := fetchData("")

		assert.Error(t, err)
		assert.Equal(t, "", res)
	})
}

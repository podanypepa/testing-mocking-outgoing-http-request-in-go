package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/podanypepa/testing-mocking-outgoing-http-request-in-go/httpclient"
)

func fetchData(url string) (string, error) {
	client := httpclient.Client

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api err: " + res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	if string(body) == "" {
		return "", fmt.Errorf("body reading error")
	}

	return string(body), nil
}

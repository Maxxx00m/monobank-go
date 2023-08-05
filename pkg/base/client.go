package base

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

// Get is a function to make a GET request to the Monobank API.
func (c *Client) Get(path string, headers map[string]string) ([]byte, error) {
	return c.makeRequest(http.MethodGet, path, headers, nil)
}

// makeRequest is a helper function to build a request to the Monobank API.
func (c *Client) makeRequest(method, path string, headers map[string]string, data []byte) ([]byte, error) {
	var bodyBuffer io.Reader
	if data != nil {
		bodyBuffer = bytes.NewBuffer(data)
	}

	request, err := http.NewRequest(method, c.makeUrl(path), bodyBuffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	if headers == nil {
		for headerName, headerValue := range headers {
			request.Header.Set(headerName, headerValue)
		}
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("too many requests")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// makeUrl is a helper function to build a URL to the Monobank API.
func (c *Client) makeUrl(path string) string {
	return fmt.Sprintf("https://api.monobank.ua/%s", path)
}

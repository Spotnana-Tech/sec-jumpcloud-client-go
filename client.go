package jcclient

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const HostURL = "https://console.jumpcloud.com"

type Client struct {
	HostURL    *url.URL
	HTTPClient *http.Client
	Headers    http.Header
}

// NewClient factory returns a prepared client
func NewClient(token string) (*Client, error) {
	parsedUrl, err := url.Parse(HostURL)
	c := Client{
		HostURL:    parsedUrl,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Headers: http.Header{
			"Accept":       {"application/json"},
			"Content-Type": {"application/json"},
			"x-api-key":    {token},
		},
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// doRequest is a helper function to prepare http requests
// This is not implemented yet
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	// Prepare and send request
	req.Header = c.Headers
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

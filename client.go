package jcclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const HostURL = "https://console.jumpcloud.com"

type Client struct {
	HostURL    *url.URL // or string for simplicity
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) (*Client, error) {
	parsedUrl, err := url.Parse(HostURL)
	c := Client{
		HostURL:    parsedUrl,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      token,
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

	req.Header.Set("x-api-token", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

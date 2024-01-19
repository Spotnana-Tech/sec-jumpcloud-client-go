package jcclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const HostURL = "https://console.jumpcloud.com"

// TODO: Add pagination function, add HTTPrequest function

type Client struct {
	HostURL    *url.URL // or string for simplicity
	HTTPClient *http.Client
	Headers    http.Header
}

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

// This isn't working, and I'm not sure why. Handling requests in methods until I can fix this
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	//token := c.Token
	req.Header = c.Headers

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(res.Status)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	} else if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// jumpcloud is a package to interact with the JumpCloud API
//
// # Create a new client
// c, err := jumpcloud.NewClient(os.Getenv("JC_API_KEY"))
//
// # Use the Client
// groups, err := c.GetAllUserGroups()
//

package jumpcloud

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const HostURL = "https://console.jumpcloud.com"

// Client is a struct to hold the client data and http client
type Client struct {
	HostURL    *url.URL
	HTTPClient *http.Client
	Headers    http.Header
}

// NewClient factory returns a prepared client
// token is the JumpCloud API key
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
// This is not implemented yet. Some methods need access to the request object still.
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

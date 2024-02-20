package jcclient

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetUser returns the details of a user by ID
//
// Example:
// user, err := c.GetUser("5f5e3e3e")
func (c *Client) GetUser(userId string) (user SystemUser, err error) {
	// Prepare request
	c.HostURL.Path = "/api/systemusers/" + userId
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return user, err
	}

	// Parse response
	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &user)
	return user, err
}

package jcclient

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetUser is a struct that represents a Jumpcloud user group
func (c *Client) GetUser(userId string) (user SystemUser, err error) {
	c.HostURL.Path = "/api/systemusers/" + userId
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers
	response, _ := c.HTTPClient.Do(req)
	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &user)
	return user, err
}

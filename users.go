package jcclient

import (
	"encoding/json"
	"net/http"
)

// GetUser is a struct that represents a Jumpcloud user group
func (c *Client) GetUser(userId string) (user SystemUser, err error) {
	c.HostURL.Path = "/api/systemusers/" + userId
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	body, _ := c.doRequest(req)
	err = json.Unmarshal(body, &user)
	return user, err
}

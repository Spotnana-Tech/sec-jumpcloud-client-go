package jumpcloud

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// GetUser returns the details of a user by ID
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
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &user)
	return user, err
}

// GetUserIDFromEmail returns the details of a user
func (c *Client) GetUserIDFromEmail(userEmail string) (string, error) {
	// Prepare request
	var searchStruct struct {
		TotalCount int `json:"totalCount"`
		Results    []struct {
			ID string `json:"_id"`
		}
	}
	// Filter on email and return _id field only
	params := url.Values{
		"filter": {"email:$eq:" + userEmail},
		"fields": {"_id"},
	}
	c.HostURL.Path = "/api/systemusers"
	c.HostURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	// Parse response
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchStruct)

	// Returning the first result
	// If no results the return empty string, this means that the search failed
	if len(searchStruct.Results) == 0 {
		return "", err
	} else {
		// If we have results, return the first result
		return searchStruct.Results[0].ID, err
	}
}

// GetUserEmailFromID returns the details of a user
func (c *Client) GetUserEmailFromID(userID string) (string, error) {
	// Prepare request
	var searchStruct struct {
		TotalCount int `json:"totalCount"`
		Results    []struct {
			Email string `json:"email"`
		}
	}
	// Filter on userId and return email field only
	params := url.Values{
		"filter": {"_id:$eq:" + userID},
		"fields": {"email"},
	}
	c.HostURL.Path = "/api/systemusers"
	c.HostURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	// Parse response
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchStruct)

	// Returning the first result
	// If no results the return empty string, this means that the search failed
	if len(searchStruct.Results) == 0 {
		return "", err
	} else {
		// If we have results, return the first result
		return searchStruct.Results[0].Email, err
	}
}

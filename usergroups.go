package jcclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// GetAllUserGroups returns a list of all users in Jumpcloud
func (c *Client) GetAllUserGroups() (allUserGroups UserGroups, err error) {
	totalRecords := 0
	c.HostURL.Path = "/api/v2/usergroups"
	c.HostURL.RawQuery = "limit=100&skip=0"
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()
	// Set our totalRecords count and pull out data out
	totalRecords, err = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int
	body, err := io.ReadAll(response.Body)                                 // response body is []byte
	err = json.Unmarshal(body, &allUserGroups)                             // Unmarshal the JSON into our struct
	if err != nil {
		return nil, err
	}

	// While all groups is less than the total number of records...
	if len(allUserGroups) < totalRecords {
		var tempData UserGroups // Create a temp slice to hold the data
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		if err != nil {
			return nil, err
		}
		newStep := strconv.Itoa(currentSkip + stepValue)

		c.HostURL.RawQuery = "limit=100&skip=" + newStep
		req.URL = c.HostURL
		response, err := c.HTTPClient.Do(req)
		body, err := io.ReadAll(response.Body) // response body is []byte
		err = json.Unmarshal(body, &tempData)
		allUserGroups = append(allUserGroups, tempData...)
		if err != nil {
			return nil, err
		}
	}
	return allUserGroups, err
}

// CreateUserGroup creates a new user group
func (c *Client) CreateUserGroup(newUser UserGroup) (userGroup NewUserGroup, err error) {
	c.HostURL.Path = "/api/v2/usergroups"
	jsonBody, err := json.Marshal(newUser)
	request, err := http.NewRequest(
		http.MethodPost,
		c.HostURL.String(),
		bytes.NewReader(jsonBody),
	)
	request.Header = c.Headers
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

// CreateUserGroups creates multiple user groups
func (c *Client) CreateUserGroups(newUserGroups []UserGroup) (userGroups []NewUserGroup, err error) {
	for _, usergroup := range newUserGroups {
		new, err := c.CreateUserGroup(usergroup)
		if err != nil {
			return nil, err
		}
		userGroups = append(userGroups, new)
	}
	return userGroups, err
}

// GetUserGroup query for a specific user group by ID
func (c *Client) GetUserGroup(groupId string) (userGroup UserGroup, err error) {
	c.HostURL.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodGet,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

// DeleteUserGroup deletes a user group
func (c *Client) DeleteUserGroup(groupId string) (err error) {
	c.HostURL.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodDelete,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()
	return err
}

// GetGroupMembers returns a list of users in a group
func (c *Client) GetGroupMembers(groupId string) (groupMembers GroupMembership, err error) {
	totalRecords := 0
	c.HostURL.Path = "/api/v2/usergroups/" + groupId + "/members"
	c.HostURL.RawQuery = "limit=100&skip=0"
	request, err := http.NewRequest(
		http.MethodGet,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers
	response, _ := c.HTTPClient.Do(request)
	defer response.Body.Close()

	totalRecords, _ = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int
	// Set our totalRecords count and pull out data out
	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &groupMembers)
	if err != nil {
		return nil, err
	}

	for len(groupMembers) < totalRecords {
		var tempData GroupMembership // Create a temp slice to hold the data
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		newStep := strconv.Itoa(currentSkip + stepValue)
		c.HostURL.RawQuery = "limit=100&skip=" + newStep
		request.URL = c.HostURL
		response, err := c.HTTPClient.Do(request)
		body, err := io.ReadAll(response.Body) // response body is []byte
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &tempData)
		groupMembers = append(groupMembers, tempData...)
	}
	return groupMembers, err
}

package jumpcloud

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// GetAllUserGroups returns a list of all usergroups in Jumpcloud
func (c *Client) GetAllUserGroups() (allUserGroups UserGroups, err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups"
	c.HostURL.RawQuery = "limit=100&skip=0"
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()

	// Set our totalRecords count and pull out data out
	totalRecords, err := strconv.Atoi(response.Header.Get("x-total-count"))
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &allUserGroups)
	if err != nil {
		return nil, err
	}

	// While all groups is less than the total number of records...
	if len(allUserGroups) < totalRecords {
		var tempData UserGroups // Create a temp slice to hold the data

		// Calculate new step value and add it to params
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		if err != nil {
			return nil, err
		}
		newStep := strconv.Itoa(currentSkip + stepValue)

		// Update URL to include new params
		c.HostURL.RawQuery = "limit=100&skip=" + newStep
		req.URL = c.HostURL

		// Send Request
		response, err := c.HTTPClient.Do(req)
		body, err := io.ReadAll(response.Body)
		err = json.Unmarshal(body, &tempData)

		// Add tempData to allUserGroups
		allUserGroups = append(allUserGroups, tempData...)
		if err != nil {
			return nil, err
		}
	}
	return allUserGroups, err
}

// CreateUserGroup creates a new user group
func (c *Client) CreateUserGroup(newUser UserGroup) (userGroup UserGroup, err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups"
	jsonBody, err := json.Marshal(newUser)
	request, err := http.NewRequest(
		http.MethodPost,
		c.HostURL.String(),
		bytes.NewReader(jsonBody),
	)
	request.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()

	// Parse response
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

// CreateUserGroups creates multiple user groups
func (c *Client) CreateUserGroups(newUserGroups []UserGroup) (userGroups []UserGroup, err error) {
	// Iterate through the user groups and create them
	for _, usergroup := range newUserGroups {
		// Create the user group
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
	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodGet,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()

	// Parse response
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

// DeleteUserGroup deletes a user group
func (c *Client) DeleteUserGroup(groupId string) (err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodDelete,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers
	// Send request
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()

	// TODO -- Check for 204
	return err
}

// GetGroupMembers returns a list of users in a group
func (c *Client) GetGroupMembers(groupId string) (groupMembers GroupMembership, err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups/" + groupId + "/members"
	c.HostURL.RawQuery = "limit=100&skip=0"
	request, err := http.NewRequest(
		http.MethodGet,
		c.HostURL.String(),
		nil,
	)
	request.Header = c.Headers

	// Send request
	response, _ := c.HTTPClient.Do(request)
	defer response.Body.Close()

	// Set our totalRecords count and pull out data out
	totalRecords, _ := strconv.Atoi(response.Header.Get("x-total-count"))
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &groupMembers)
	if err != nil {
		return nil, err
	}

	// While all group members is less than the total number of records...
	for len(groupMembers) < totalRecords {
		var tempData GroupMembership // Create a temp slice to hold the data
		// Calculate new step value and add it to params
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		newStep := strconv.Itoa(currentSkip + stepValue)
		c.HostURL.RawQuery = "limit=100&skip=" + newStep

		// Update URL to include new params
		request.URL = c.HostURL
		// Send Request
		response, err := c.HTTPClient.Do(request)
		body, err := io.ReadAll(response.Body) // response body is []byte
		if err != nil {
			return nil, err
		}

		// Parse response, add results to tempData
		err = json.Unmarshal(body, &tempData)
		groupMembers = append(groupMembers, tempData...)
	}
	return groupMembers, err
}

// GetUserGroupByName returns a group by name
func (c *Client) GetUserGroupByName(groupName string) (userGroup UserGroup, err error) {
	// Prepare request
	var results UserGroups
	c.HostURL.Path = "/api/v2/usergroups"
	// Limiting results to 1, and filtering by name... this should be unique
	c.HostURL.RawQuery = "limit=1&filter=name:eq:" + groupName
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()

	// Parse response
	body, err := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &results)   // Unmarshal the JSON into our struct

	// If no results the return empty userGroup, this means that the search failed
	if len(results) == 0 {
		return userGroup, err
	} else {
		// If we have results, return the first result
		userGroup = results[0]
		return userGroup, err
	}
}

// UpdateUserGroup updates a user group
func (c *Client) UpdateUserGroup(groupId string, updatedUserGroup UserGroup) (userGroup UserGroup, err error) {
	// If update does not contain a description field, set it to the old description
	// We do this because the API will overwrite the description field with an empty string if not passed

	if updatedUserGroup.Description == "" {
		oldGroupData, _ := c.GetUserGroup(groupId)
		if oldGroupData.Description != "" {
			updatedUserGroup.Description = oldGroupData.Description
		}
	}

	// Prepare request
	c.HostURL.Path = "/api/v2/usergroups/" + groupId
	jsonBody, err := json.Marshal(updatedUserGroup)
	request, err := http.NewRequest(
		http.MethodPut,
		c.HostURL.String(),
		bytes.NewReader(jsonBody),
	)
	request.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(request)
	defer response.Body.Close()

	// Parse response
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

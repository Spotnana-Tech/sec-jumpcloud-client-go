package jcclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// GetAllApplications returns all applications
func (c *Client) GetAllApplications() (allApplications AllApps, err error) {
	// Prepare request
	totalApps := 0
	c.HostURL.Path = "/api/v2/applications/"
	c.HostURL.RawQuery = "limit=100&skip=0"
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, _ := c.HTTPClient.Do(req)
	defer response.Body.Close()

	// Parse response
	// totalApps is the total number of records
	totalApps, err = strconv.Atoi(response.Header.Get("x-total-count"))
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &allApplications)
	if err != nil {
		return nil, err
	}

	// While all groups is less than the total number of records
	if len(allApplications) < totalApps {
		var tempData AllApps // Create a temp slice to hold the data

		// Get step and skip values
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		if err != nil {
			return nil, err
		}

		// Calculate new step value and add it to params
		newStep := strconv.Itoa(currentSkip + stepValue)
		c.HostURL.RawQuery = "limit=100&skip=" + newStep
		req.URL = c.HostURL // Update URL to include new params

		// Send Request
		response, err := c.HTTPClient.Do(req)

		// Parse response, add results to tempData
		body, err := io.ReadAll(response.Body)
		err = json.Unmarshal(body, &tempData)

		// Add tempData to allApplications
		allApplications = append(allApplications, tempData...)
		if err != nil {
			return nil, err
		}
	}
	return allApplications, err
}

// GetApplication returns a single application
func (c *Client) GetApplication(appId string) (application App, err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/applications/" + appId
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, _ := c.HTTPClient.Do(req)

	// Parse response
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &application)
	return application, err
}

// GetAppAssociations returns a single application
// groupType can be either "user_group" or "user"
func (c *Client) GetAppAssociations(appId string, groupType string) (appAssociations AppAssociations, err error) {
	// Prepare request
	totalAssociations := 0
	c.HostURL.Path = "/api/v2/applications/" + appId + "/associations"
	c.HostURL.RawQuery = "targets=" + groupType + "&limit=100&skip=0"
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers

	// Send request
	response, _ := c.HTTPClient.Do(req)
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &appAssociations)

	totalAssociations, err = strconv.Atoi(response.Header.Get("x-total-count"))
	// While all groups is less than the total number of records...
	if len(appAssociations) < totalAssociations {
		var tempData AppAssociations // Create a temp slice to hold the data
		stepValue, err := strconv.Atoi(c.HostURL.Query().Get("limit"))
		currentSkip, err := strconv.Atoi(c.HostURL.Query().Get("skip"))
		if err != nil {
			return nil, err
		}
		newStep := strconv.Itoa(currentSkip + stepValue)

		c.HostURL.RawQuery = "targets=user_group&limit=100&skip=" + newStep
		req.URL = c.HostURL
		response, err := c.HTTPClient.Do(req)
		body, err := io.ReadAll(response.Body)
		err = json.Unmarshal(body, &tempData)
		appAssociations = append(appAssociations, tempData...)
		if err != nil {
			return nil, err
		}
	}
	return appAssociations, err
}

// AssociateGroupWithApp associates a group with an application
func (c *Client) AssociateGroupWithApp(appId string, groupId string) (err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/applications/" + appId + "/associations"
	j, _ := json.Marshal(AppAssociationModifier{
		ID:   groupId,
		OP:   "add",
		Type: "user_group",
	})
	bodyReader := bytes.NewReader(j)
	req, err := http.NewRequest(http.MethodPost, c.HostURL.String(), bodyReader)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	if response.StatusCode == 204 {
		// 204 = OK
		return nil
	}
	return err
}

// RemoveGroupFromApp removes a group from an application
func (c *Client) RemoveGroupFromApp(appId string, groupId string) (err error) {
	// Prepare request
	c.HostURL.Path = "/api/v2/applications/" + appId + "/associations"
	j, _ := json.Marshal(map[string]string{
		"id":   groupId,
		"op":   "remove",
		"type": "user_group",
	})
	bodyReader := bytes.NewReader(j)
	req, err := http.NewRequest(http.MethodPost, c.HostURL.String(), bodyReader)
	req.Header = c.Headers

	// Send request
	response, err := c.HTTPClient.Do(req)
	defer response.Body.Close()
	if response.StatusCode == 204 {
		// 204 -- OK
		return nil
	}
	return err
}

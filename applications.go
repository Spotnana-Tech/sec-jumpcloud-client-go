package jcclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// GetAllApplications is a struct that represents a Jumpcloud user group
func (c *Client) GetAllApplications() (allApplications AllApps, err error) {
	totalApps := 0
	c.HostURL.Path = "/api/v2/applications/"
	c.HostURL.RawQuery = "limit=100&skip=0"
	req, err := http.NewRequest(http.MethodGet, c.HostURL.String(), nil)
	req.Header = c.Headers
	response, _ := c.HTTPClient.Do(req)
	defer response.Body.Close()
	totalApps, err = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int

	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &allApplications)
	if err != nil {
		return nil, err
	}

	// While all groups is less than the total number of records...
	if len(allApplications) < totalApps {
		var tempData AllApps // Create a temp slice to hold the data
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
		allApplications = append(allApplications, tempData...)
		if err != nil {
			return nil, err
		}
	}
	return allApplications, err
}

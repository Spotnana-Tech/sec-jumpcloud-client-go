package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (jc *JC) GetGroups() (allGroups []UserGroup, err error) {
	var totalRecords int
	// Set API call details and make the request
	jc.Url.Path = "/api/v2/groups"
	request, err := http.NewRequest(
		http.MethodGet,
		jc.Url.String(),
		nil,
	)

	request.Header = jc.Headers
	response, _ := jc.Client.Do(request)

	// Set our totalRecords count and pull out data out
	totalRecords, _ = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int
	body, _ := io.ReadAll(response.Body)                                 // response body is []byte
	err = json.Unmarshal(body, &allGroups)                               // Unmarshal the JSON into our struct

	// While all groups is less than the total number of records...
	if len(allGroups) < totalRecords {

		var tempData []UserGroup // Create a temp slice to hold the data
		stepValue, _ := strconv.Atoi(jc.Url.Query().Get("limit"))
		currentSkip, _ := strconv.Atoi(jc.Url.Query().Get("skip"))
		newStep := strconv.Itoa(currentSkip + stepValue)

		jc.Url.RawQuery = "limit=100&skip=" + newStep
		request.URL = &jc.Url
		response, _ := jc.Client.Do(request)
		body, _ := io.ReadAll(response.Body) // response body is []byte
		_ = json.Unmarshal(body, &tempData)
		allGroups = append(allGroups, tempData...)

	}
	return allGroups, err
}

func (jc *JC) GetGroupMembers(groupId string) (groupMembers []map[string]string, err error) {
	var totalRecords int
	// Set API call details and make the request
	jc.Url.RawQuery = "limit=100&skip=0"
	jc.Url.Path = "/api/v2/usergroups/" + groupId + "/membership"
	request, err := http.NewRequest(
		http.MethodGet,
		jc.Url.String(),
		nil,
	)
	request.Header = jc.Headers
	response, _ := jc.Client.Do(request)
	defer response.Body.Close()

	totalRecords, _ = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int
	// Set our totalRecords count and pull out data out
	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &groupMembers)

	for len(groupMembers) < totalRecords {
		var tempData []map[string]string // Create a temp slice to hold the data
		stepValue, _ := strconv.Atoi(jc.Url.Query().Get("limit"))
		currentSkip, _ := strconv.Atoi(jc.Url.Query().Get("skip"))
		newStep := strconv.Itoa(currentSkip + stepValue)
		jc.Url.RawQuery = "limit=100&skip=" + newStep
		request.URL = &jc.Url
		response, _ := jc.Client.Do(request)
		body, _ := io.ReadAll(response.Body) // response body is []byte
		_ = json.Unmarshal(body, &tempData)
		groupMembers = append(groupMembers, tempData...)
	}
	return groupMembers, err
}

func (jc *JC) GetUser(userId string) (user User, err error) {
	// Set API call details and make the request
	jc.Url.RawQuery = "limit=100&skip=0"
	jc.Url.Path = "/api/systemusers/" + userId
	request, err := http.NewRequest(
		http.MethodGet,
		jc.Url.String(),
		nil,
	)
	request.Header = jc.Headers
	response, _ := jc.Client.Do(request)
	defer response.Body.Close()

	// Set our totalRecords count and pull out data out
	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &user)
	return user, err
}

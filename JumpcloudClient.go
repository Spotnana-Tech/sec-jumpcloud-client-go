package jcclient

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var _ = godotenv.Load()
var JCClient = JC{
	// A pre-configured client
	Url: url.URL{
		Scheme:   "https",
		Host:     "console.jumpcloud.com",
		RawQuery: "limit=100&skip=0",
	},
	Headers: http.Header{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
		"x-api-key":    {os.Getenv("JC_API_KEY")}, // JCClient API via env var, maybe pull from config file?
	},
	Client: http.Client{Timeout: 10 * time.Second},
}

func NewClient(ApiKey string) (jc JC) {
	jc.Url = url.URL{
		Scheme:   "https",
		Host:     "console.jumpcloud.com",
		RawQuery: "limit=100&skip=0",
	}
	jc.Headers = http.Header{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
		"x-api-key":    {ApiKey}, // JCClient API via user input
	}
	jc.Client = http.Client{Timeout: 10 * time.Second}
	return jc
}

func (jc *JC) GetAllUserGroups() (allUserGroups []UserGroup, err error) {
	var totalRecords int
	// Set API call details and make the request
	jc.Url.Path = "/api/v2/usergroups"
	request, err := http.NewRequest(
		http.MethodGet,
		jc.Url.String(),
		nil,
	)
	request.Header = jc.Headers
	response, err := jc.Client.Do(request)

	// Set our totalRecords count and pull out data out
	totalRecords, err = strconv.Atoi(response.Header.Get("x-total-count")) // Converting str to int
	body, err := io.ReadAll(response.Body)                                 // response body is []byte
	err = json.Unmarshal(body, &allUserGroups)                             // Unmarshal the JSON into our struct

	// While all groups is less than the total number of records...
	if len(allUserGroups) < totalRecords {

		var tempData []UserGroup // Create a temp slice to hold the data
		stepValue, _ := strconv.Atoi(jc.Url.Query().Get("limit"))
		currentSkip, _ := strconv.Atoi(jc.Url.Query().Get("skip"))
		newStep := strconv.Itoa(currentSkip + stepValue)

		jc.Url.RawQuery = "limit=100&skip=" + newStep
		request.URL = &jc.Url
		response, _ := jc.Client.Do(request)
		body, _ := io.ReadAll(response.Body) // response body is []byte
		_ = json.Unmarshal(body, &tempData)
		allUserGroups = append(allUserGroups, tempData...)

	}
	return allUserGroups, err
}

func (jc *JC) GetUserGroup(groupId string) (userGroup map[string]any, err error) {
	// Set API call details and make the request
	jc.Url.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodGet,
		jc.Url.String(),
		nil,
	)
	request.Header = jc.Headers
	response, _ := jc.Client.Do(request)
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

func (jc *JC) CreateUserGroup(groupMap map[string]string) (userGroup map[string]string, err error) {
	// Set API call details and make the request
	jc.Url.Path = "/api/v2/usergroups"
	jsonBody, err := json.Marshal(groupMap)
	request, err := http.NewRequest(
		http.MethodPost,
		jc.Url.String(),
		bytes.NewReader(jsonBody),
	)
	request.Header = jc.Headers
	response, err := jc.Client.Do(request)
	defer response.Body.Close()

	// Set our totalRecords count and pull out data out
	body, err := io.ReadAll(response.Body) // response body is []byte
	err = json.Unmarshal(body, &userGroup)
	return userGroup, err
}

func (jc *JC) DeleteUserGroup(groupId string) (err error) {
	// Set API call details and make the request
	jc.Url.Path = "/api/v2/usergroups/" + groupId
	request, err := http.NewRequest(
		http.MethodDelete,
		jc.Url.String(),
		nil,
	)
	request.Header = jc.Headers
	response, _ := jc.Client.Do(request)
	defer response.Body.Close()
	//body, _ := io.ReadAll(response.Body) // response body is []byte
	//err = json.Unmarshal(body, &deleted)
	return err
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

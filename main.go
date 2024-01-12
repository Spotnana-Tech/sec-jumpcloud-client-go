package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	// Load .env file for api key
	err := godotenv.Load()
	if err != nil {
		return
	}
	// Create our Jumpcloud API client
	jumpcloudApiClient := JC{
		Url: url.URL{
			Scheme:   "https",
			Host:     "console.jumpcloud.com",
			Path:     "/api",
			RawQuery: "limit=100&skip=0",
		},
		Headers: http.Header{
			"Accept":       {"application/json"},
			"Content-Type": {"application/json"},
			"x-api-key":    {os.Getenv("JC_API_KEY")}, // Jumpcloud API via env var, maybe pull from config file?
		},
		Client: http.Client{Timeout: 10 * time.Second},
	}

	// Get all userGroups
	groups, err := jumpcloudApiClient.GetGroups() // Get all groups
	for _, group := range groups {
		// Get all groupMembers
		members, _ := jumpcloudApiClient.GetGroupMembers(group.Id) // Get all members of group.Id
		fmt.Println(group.Name, "-", group.Id, "-", len(members), "members")

		// Get each groupMember's info
		for _, member := range members {
			user, _ := jumpcloudApiClient.GetUser(member["id"]) // Get user info for each member
			fmt.Println(user.Id, user.Displayname, user.Email, user.Department)
		}

		fmt.Println()
	}
}

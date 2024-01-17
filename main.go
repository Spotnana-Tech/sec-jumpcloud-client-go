package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"time"
)

var _ = godotenv.Load()
var JCClient = JC{
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

func main() {
	start := time.Now()

	groups, _ := JCClient.GetAllUserGroups() // Get all groups
	fmt.Println("Total Groups:", len(groups))
	elapsed := time.Since(start)
	fmt.Println("[!] Total runtime:", elapsed.Round(time.Millisecond))

	/* Example Workflow: Get all Groups, their members, and the member details

	allGroups, err := JCClient.GetAllUserGroups()

	for _, group := range allGroups {

		// Get all groupMembers
		members, _ := JCClient.GetGroupMembers(group.Id)
		fmt.Println(group.Name, "-", group.Id, "-", len(members), "members")
		fmt.Println("--------------------------------------------------")

		// Get group details
		groupDetails, _ := JCClient.GetUserGroup(group.Id)
		fmt.Println(groupDetails["id"], groupDetails["name"], groupDetails["description"])

			// Get each groupMember's info
			for _, member := range members {
				user, _ := JCClient.GetUser(member["id"])
				fmt.Println(user.Id, user.Displayname, user.Email, user.Department)
			}*/
}

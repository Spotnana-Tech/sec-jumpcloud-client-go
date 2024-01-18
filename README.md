# sec-jumpcloud-client-go
### A go client for the Jumpcloud API
[![Tests](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml/badge.svg)](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml)
# Getting Started
Clone the repo and run tests locally
```shell
git clone git@github.com:Spotnana-Tech/sec-jumpcloud-client-go.git
cd sec-jumpcloud-client-go
go test -v
```
## Example Usage
Below is a simple example of how to use the client to get all groups and their members.
```go
package main

import (
	"fmt"
	"github.com/Spotnana-Tech/sec-jumpcloud-client-go"
	"log"
	"os"
)

func main() {
	// Create a new Jumpcloud client, pulling the API key from the environment
	c, err := jcclient.NewClient(os.Getenv("JC_API_KEY"))
	if err != nil {
		log.Panic("Error creating client:", err)
	}
    
	groups, err := c.GetAllUserGroups()             // Get all groups
	for _, group := range groups {                  // Loop through each group
		members, _ := c.GetGroupMembers(group.ID)   // Get all groupMembers in each group
		fmt.Println(group.Name, group.ID, "-", len(members), "members")
		// example-group 1234567890 - 21 members
		
		for _, member := range members {        // Loop through each groupMember
			user, _ := c.GetUser(member.To.ID) // Lookup each member and get their details
			fmt.Println(user.ID, user.Displayname, user.Email, user.Department)
			// 0987654321 John Doe jdoe@spotnana.com Engineering
		}
	}
}
```

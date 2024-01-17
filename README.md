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

#### Test Output
```shell
=== RUN   TestCreateAndDeleteUserGroup
New Group ID 65a856555c27e500013b10da
Lookup Group ID 65a856555c27e500013b10da
Group ID 65a856555c27e500013b10da Deleted
--- PASS: TestCreateAndDeleteUserGroup (0.80s)
=== RUN   TestGetAllUserGroups
Total Groups Returned: 95
--- PASS: TestGetAllUserGroups (0.15s)
PASS

```

## Example Usage

```go
// Example Workflow: Get all Groups, their members, and the member details
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

	// Get all groups
	groups, err := c.GetAllUserGroups()
	for _, group := range groups {

		// Get all groupMembers in each group
		members, _ := c.GetGroupMembers(group.ID)
		fmt.Println(group.Name, group.ID, "-", len(members), "members")

		// Lookup each member and get their details
		for _, member := range members {
			user, _ := c.GetUser(member.To.ID) // Jumpcloud JSON response is ugly
			fmt.Println(user.ID, user.Displayname, user.Email, user.Department)
		}
	}
}
```

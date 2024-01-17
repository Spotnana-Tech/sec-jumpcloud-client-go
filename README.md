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
--- PASS: TestCreateAndDeleteUserGroup (1.08s)
=== RUN   TestGetAllUserGroups
--- PASS: TestGetAllUserGroups (0.14s)
PASS
ok      sec-jumpcloud-client-go 1.417s
```

## Example Usage
```go
// Example Workflow: Get all Groups, their members, and the member details
package main

import (
    "fmt"
    "github.com/Spotnana-Tech/sec-jumpcloud-client-go"
)

func main() {
	// Create a new Jumpcloud client
	Jumpcloud := jcclient.JCClient
	// Get all userGroups
	allGroups, err := Jumpcloud.GetAllUserGroups()

	for _, group := range allGroups {
		// Get all groupMembers
		members, _ := Jumpcloud.GetGroupMembers(group.Id)
		fmt.Println(group.Name, "-", group.Id, "-", len(members), "members")

		// Get group details
		groupDetails, _ := Jumpcloud.GetUserGroup(group.Id)
		fmt.Println(groupDetails["id"], groupDetails["name"], groupDetails["description"])

		// Get each groupMember's info
		for _, member := range members {
			user, _ := Jumpcloud.GetUser(member["id"])
			fmt.Println(user.Id, user.Displayname, user.Email, user.Department)
		}
	}
}
```

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
	if err != nil {log.Panic("Error creating client:", err)}
	
	// Get all usergroups and print them
	g, err := c.GetAllUserGroups() 
	if err != nil {log.Panic("Error getting groups:", err)}
	fmt.Println("Number of groups:" , len(g))
	fmt.Println("Groups:", g)
}
```

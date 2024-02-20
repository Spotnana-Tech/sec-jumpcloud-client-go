# sec-jumpcloud-client-go
### A go client for Jumpcloud
[![Tests](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml/badge.svg)](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml)

# Example Usage
```go
package main

import (
	"github.com/Spotnana-Tech/sec-jumpcloud-client-go"
)

func main() {
	// Create a new Jumpcloud client
	c, err := jcclient.NewClient("api_key")
	
	// Get all usergroups
	g, err := c.GetAllUserGroups() 
}
```

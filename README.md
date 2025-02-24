# sec-jumpcloud-client-go
### A Go client for Jumpcloud

This project is archived.

_"Jumpcloud" is a trademark of Jumpcloud, Inc.
"Go" is a trademark of Google LLC or its affiliate ("Google") for its programming language (see https://go.dev/brand). 
These marks are used nominatively to indicate the nature and function of Spotnana's 
source code, which is neither sponsored or endorsed by Jumpcloud, Inc. or Google._

[![Tests](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml/badge.svg)](https://github.com/Spotnana-Tech/sec-jumpcloud-client-go/actions/workflows/go.yml)

# Example Usage
```go
package main

import (
	"github.com/Spotnana-Tech/sec-jumpcloud-client-go"
)

func main() {
	// Create a new Jumpcloud client
	c, err := jumpcloud.NewClient("api_key")
	
	// Get all usergroups
	g, err := c.GetAllUserGroups() 
}
```

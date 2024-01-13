package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

var Jumpcloud = JC{
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

func TestMain(m *testing.M) {
	godotenv.Load()
	os.Exit(m.Run())
}

func TestCreateAndDeleteUserGroup(t *testing.T) {
	// Create userGroup
	newGroup, _ := Jumpcloud.CreateUserGroup(map[string]string{
		"name":        "SecJumpcloudTestGroup",
		"description": "Created via sec-jumpcloud-client-go unit test, if not deleted, please delete me",
	})
	// Get userGroup
	testNewGroup, _ := Jumpcloud.GetUserGroup(newGroup["id"])

	// Compare the two
	if testNewGroup["id"] != newGroup["id"] {
		t.Errorf("Created groupID %q does not match %q",
			newGroup["id"],
			testNewGroup["id"],
		)
	}

	// Delete userGroup
	_ = Jumpcloud.DeleteUserGroup(newGroup["id"])
	isGroupDeleted, _ := Jumpcloud.GetUserGroup(newGroup["id"])
	if isGroupDeleted["message"] != "Not Found" {
		t.Errorf("Unable to delete test-created groupID %q Error Message: %q",
			newGroup["id"],
			isGroupDeleted["message"],
		)
	}
}

func TestGetAllUserGroups(t *testing.T) {
	groups, _ := Jumpcloud.GetAllUserGroups()
	if len(groups) == 0 {
		t.Errorf("No groups returned")
	}
}

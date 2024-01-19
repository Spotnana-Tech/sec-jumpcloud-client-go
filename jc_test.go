package jcclient

import (
	"math/rand"
	"os"
	"testing"
)

func TestCreateAndDeleteUserGroup(t *testing.T) {
	// Create userGroup
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	newGroup, err := c.CreateUserGroup(map[string]string{
		"name":        "sec-jumpcloud-client-go-unit-test",
		"description": "Created via sec-jumpcloud-client-go unit test, please delete me!",
	})

	// Get userGroup
	testNewGroup, err := c.GetUserGroup(newGroup.ID)
	// Check for errors, check for identical groupIDs
	if testNewGroup.ID != newGroup.ID {
		t.Errorf("Unable to create group, or created group does not match ID lookup %v", err)
	}

	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
	isGroupDeleted, err := c.GetUserGroup(newGroup.ID)

	if isGroupDeleted.ID != "" {
		t.Errorf("Unable to delete test-created groupID %v", err)
	}
}

func TestGetAllUserGroups(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.GetAllUserGroups()

	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

func TestGetAllApps(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	apps, err := c.GetAllApplications()

	if len(apps) == 0 {
		t.Errorf("No apps returned")
		t.Errorf("Function Error: %q", err)
	}
}

func TestGetRandomUser(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groupId := "6479fcdf1be9850001728dec"
	users, err := c.GetGroupMembers(groupId)

	if len(users) == 0 {
		t.Errorf("No users returned")
		t.Errorf("Function Error: %q", err)
	}
	// Random index from users slice
	randomInt := rand.Int() % len(users)
	// Id of random user
	randomUserId := users[randomInt].To.ID
	// Lookup random user
	randomUser, err := c.GetUser(randomUserId)

	if randomUser.ID == "" || randomUser.Email == "" || randomUser.Username == "" {

		t.Errorf("No user returned")
		t.Errorf("Function Error: %q", err)
	}
}

func TestGetApp(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	app, err := c.GetApplication("64798af00ee9439afdfd9955")

	if app.ID != "64798af00ee9439afdfd9955" {
		t.Errorf("No app returned")
		t.Errorf("Function Error: %q", err)
	}

	associations, err := c.GetAppGroupAssociations("64798af00ee9439afdfd9955")
	if len(associations) == 0 {
		t.Errorf("No application group associations returned %v %v", app.ID, app.DisplayName)
	}
}

// TODO More testing!

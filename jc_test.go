package jcclient

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateAndDeleteUserGroup(t *testing.T) {
	// Create userGroup
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	newGroup, err := c.CreateUserGroup(map[string]string{
		"name":        "SecJumpcloudTestGroup",
		"description": "Created via sec-jumpcloud-client-go unit test, if not deleted, please delete me",
	})
	fmt.Println("New Group ID", newGroup.ID)
	// Get userGroup
	testNewGroup, err := c.GetUserGroup(newGroup.ID)
	fmt.Println("Lookup Group ID", testNewGroup.ID)
	// Check for errors, check for identical groupIDs
	if testNewGroup.ID != newGroup.ID {
		t.Errorf("Unable to create group, or created group does not match ID lookup %v", err)
	}

	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
	isGroupDeleted, err := c.GetUserGroup(newGroup.ID)
	// Check for errors, check for empty groupID
	if isGroupDeleted.ID == "" {
		fmt.Println("Group ID", newGroup.ID, "Deleted")
	}
	if isGroupDeleted.ID != "" {
		t.Errorf("Unable to delete test-created groupID %v", err)
	}
}

func TestGetAllUserGroups(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.GetAllUserGroups()
	if len(groups) > 0 {
		fmt.Println("Total Groups Returned:", len(groups))
	}
	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

// TODO More testing!

package main

import (
	"testing"
)

func TestCreateAndDeleteUserGroup(t *testing.T) {
	// Create userGroup
	newGroup, err := Jumpcloud.CreateUserGroup(map[string]string{
		"name":        "SecJumpcloudTestGroup",
		"description": "Created via sec-jumpcloud-client-go unit test, if not deleted, please delete me",
	})
	// Get userGroup
	testNewGroup, err := Jumpcloud.GetUserGroup(newGroup["id"])

	// Check for errors, check for identical groupIDs
	if testNewGroup["id"] == nil || testNewGroup["id"] != newGroup["id"] {
		t.Errorf("Unable to create group, or created group does not match ID lookup %v", err)
	}
	//Compare the two
	//if testNewGroup["id"] != newGroup["id"] {
	//	t.Logf("Created groupID %v does not match %v",
	//		newGroup["id"],
	//		testNewGroup["id"],
	//	)
	//	t.Errorf("Function Error: %v", err)
	//}

	// Delete userGroup
	err = Jumpcloud.DeleteUserGroup(newGroup["id"])
	isGroupDeleted, err := Jumpcloud.GetUserGroup(newGroup["id"])
	if isGroupDeleted["message"] != "Not Found" {
		t.Errorf("Unable to delete test-created groupID %v", err)

	}
}

func TestGetAllUserGroups(t *testing.T) {
	groups, err := Jumpcloud.GetAllUserGroups()
	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

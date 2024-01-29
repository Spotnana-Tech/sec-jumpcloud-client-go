package jcclient

import (
	"math/rand"
	"os"
	"testing"
)

func TestClient_UserGroups_CreateAndDeleteUserGroup(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))

	// Create userGroup
	testName := "sec-jumpcloud-client-go-unit-test"
	newGroupData := UserGroup{
		Name:        testName,
		Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
	}
	newGroup, err := c.CreateUserGroup(newGroupData)

	// Get userGroup, by ID and Name
	testNewGroup, err := c.GetUserGroup(newGroup.ID)
	testNewGroupByName, err := c.GetUserGroupByName(testName)

	// Check for errors, check for identical groupIDs
	if testNewGroup.ID != newGroup.ID {
		t.Errorf("Unable to create group, or created group does not match ID lookup %v", err)
	}
	if testNewGroupByName[0].Name != newGroupData.Name {
		t.Errorf("Unable to lookup group by name %v", err)
	}

	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
	isGroupDeleted, err := c.GetUserGroup(newGroup.ID)

	// Ensure group is deleted
	if isGroupDeleted.ID != "" {
		t.Errorf("Unable to delete test-created groupID %v", err)
	}
}

func TestClient_UserGroups_CreateMultipleUserGroups(t *testing.T) {
	// Create userGroup
	c, _ := NewClient(os.Getenv("JC_API_KEY"))
	newGroups := []UserGroup{
		UserGroup{
			Name:        "sec-jumpcloud-client-go-unit-test-01",
			Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
		},
		UserGroup{
			Name:        "sec-jumpcloud-client-go-unit-test-02",
			Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
		},
	}

	createdGroups, _ := c.CreateUserGroups(newGroups)
	for _, createdGroup := range createdGroups {
		// Get userGroup
		testNewGroup, err := c.GetUserGroup(createdGroup.ID)
		// Check for errors, check for identical groupIDs
		if testNewGroup.ID != createdGroup.ID {
			t.Errorf("Unable to create group, or created group does not match ID lookup %v", err)
		}
		// Delete userGroup
		err = c.DeleteUserGroup(createdGroup.ID)
		isGroupDeleted, err := c.GetUserGroup(createdGroup.ID)

		if isGroupDeleted.ID != "" {
			t.Errorf("Unable to delete test-created groupID %v", err)
		}
	}
}

func TestClient_UserGroups_GetAllUserGroups(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.GetAllUserGroups()

	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

func TestClient_Users_GetRandomUser(t *testing.T) {
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

func TestClient_Apps_GetAllApps(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	apps, err := c.GetAllApplications()

	if len(apps) == 0 {
		t.Errorf("No apps returned")
		t.Errorf("Function Error: %q", err)
	}
}

func TestClient_Apps_GetApp(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	app, err := c.GetApplication("64798af00ee9439afdfd9955")

	if app.ID != "64798af00ee9439afdfd9955" {
		t.Errorf("No app returned")
		t.Errorf("Function Error: %q", err)
	}

	associations, err := c.GetAppAssociations(
		"64798af00ee9439afdfd9955",
		"user_group",
	)
	if len(associations) == 0 {
		t.Errorf("No application group associations returned %v %v", app.ID, app.DisplayName)
	}
}

// TODO This test is broken, runs too slow, useless
//func TestClient_Apps_Client_GetAllAppAssociations(t *testing.T) {
//	c, err := NewClient(os.Getenv("JC_API_KEY"))
//	app, err := c.GetApplication("64798af00ee9439afdfd9955")
//
//	if app.ID != "64798af00ee9439afdfd9955" {
//		t.Errorf("No app returned")
//		t.Errorf("Function Error: %q", err)
//	}
//
//	// This runs for a really long time and I'm not sure why
//	associations, err := c.GetAllAppAssociations()
//
//	//fmt.Println(associations)
//	if len(associations) == 0 {
//		t.Errorf("No application group associations returned %v %v", app.ID, app.DisplayName)
//	}
//}

func TestClient_Apps_AssociateGroupWithApp(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	newGroup, err := c.CreateUserGroup(UserGroup{
		Name:        "sec-jumpcloud-client-go-unit-test-app-association",
		Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
	})
	awsSSOPOC, _ := c.GetApplication("632b3aae90fb7290ddb5667d") // AWS SSO POC App ID

	// Get userGroup
	newGroupId, err := c.GetUserGroup(newGroup.ID)

	// Associate group with app
	err = c.AssociateGroupWithApp(awsSSOPOC.ID, newGroupId.ID)
	if err != nil {
		t.Errorf("No application group associations returned %v", err)
	}

	// Remove group from app
	err = c.RemoveGroupFromApp(awsSSOPOC.ID, newGroupId.ID)
	if err != nil {
		t.Errorf("Unable to remove group from app %v", err)
	}

	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
	isGroupDeleted, err := c.GetUserGroup(newGroup.ID)

	if isGroupDeleted.ID != "" {
		t.Errorf("Unable to delete test-created groupID %v", err)
	}
}

// TODO More testing!

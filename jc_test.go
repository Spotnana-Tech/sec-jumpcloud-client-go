package jumpcloud

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

// Check if user is in group
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// TestClient_UserGroups_CreateAndDeleteUserGroup the creation and deletion of usergroups
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
	if testNewGroupByName.Name != newGroupData.Name {
		t.Errorf("Unable to lookup group by name %v", err)
	}

	// Update UserGroup name to include "-updated"
	update := UserGroup{
		Name: testName + "-updated",
	}
	updatedGroup, err := c.UpdateUserGroup(newGroup.ID, update)

	// Verify UserGroup name changed
	result, err := c.GetUserGroupByName(testName + "-updated")

	// Check result values
	if updatedGroup.Name != result.Name {
		t.Errorf("Unable to update group %v", err)
	}

	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
	isGroupDeleted, err := c.GetUserGroup(newGroup.ID)

	// Ensure group is deleted
	if isGroupDeleted.ID != "" {
		t.Errorf("Unable to delete test-created groupID %v", err)
	}
}

// TestClient_UserGroups_CreateMultipleUserGroups creating multiple groups via function call
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

// TestClient_UserGroups_GetAllUserGroups getting all groups via pagination
func TestClient_UserGroups_GetAllUserGroups(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.GetAllUserGroups()
	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

// TestClient_UserGroups_GetUserGroup getting a single group via ID
func TestClient_UserGroups_GetUserGroup(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.GetAllUserGroups()
	targetGroupId := groups[0].ID
	g, err := c.GetUserGroup(targetGroupId)
	if g.ID != targetGroupId {
		t.Errorf("No group returned")
		t.Errorf("Function Error: %q", err)
	}
}

// TestClient_Users_GetRandomUser getting a user
func TestClient_Users_GetRandomUser(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	var groupId string

	// Get all groups, find a group with members
	groups, err := c.GetAllUserGroups()
	for _, group := range groups {
		m, _ := c.GetGroupMembers(group.ID)
		if len(m) > 0 {
			groupId = group.ID
			break
		}
	}

	// Get group members
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

	randomUserEmailFromIdLookup, _ := c.GetUserEmailFromID(randomUserId)
	randomUserIDFromEmailLookup, err := c.GetUserIDFromEmail(randomUser.Email)

	if randomUserIDFromEmailLookup != randomUserId {
		t.Errorf("Unable to lookup user by email")
	}
	if randomUserEmailFromIdLookup != randomUser.Email {
		t.Errorf("Unable to lookup user by ID")
	}
}

// TestClient_Apps_GetAllApps getting all apps via pagination
func TestClient_Apps_GetAllApps(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	apps, err := c.GetAllApplications()

	if len(apps) == 0 {
		t.Errorf("No apps returned")
		t.Errorf("Function Error: %q", err)
	}
}

// TestClient_Apps_GetApp getting a single app via ID
func TestClient_Apps_GetApp(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	var targetAppId string
	var targetAppName string
	apps, err := c.GetAllApplications()
	for _, app := range apps {
		asoc, _ := c.GetAppAssociations(app.ID, "user_group")
		if len(asoc) > 0 {
			targetAppId = app.ID
			targetAppName = app.DisplayName
			break
		}
	}

	associations, err := c.GetAppAssociations(
		targetAppId,
		"user_group",
	)
	if len(associations) == 0 {
		t.Errorf("No application group associations returned %v %v", targetAppId, targetAppName)
	}
	if err != nil {
		t.Errorf("Function Error: %q", err)
	}
}

// TestClient_Apps_AssociateGroupWithApp associating and removing a group from an app
// This test looks for an application with `test` in the name and associates a group with it
// This test will fail if there are no applications with `test` in the name
func TestClient_Apps_AssociateGroupWithApp(t *testing.T) {
	var targetAppId string
	c, err := NewClient(os.Getenv("JC_API_KEY"))

	// Create group to associate with app
	newGroup, err := c.CreateUserGroup(UserGroup{
		Name:        "sec-jumpcloud-client-go-unit-test-app-association",
		Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
	})

	// Find a test app
	apps, err := c.GetAllApplications()
	for _, app := range apps {
		if strings.Contains(app.DisplayLabel, "test") {
			targetAppId = app.ID
		}
	}
	// Get userGroup
	newGroupId, err := c.GetUserGroup(newGroup.ID)

	// Associate group with app
	err = c.AssociateGroupWithApp(targetAppId, newGroupId.ID)
	if err != nil {
		t.Errorf("No application group associations returned %v", err)
	}

	// Remove group from app
	err = c.RemoveGroupFromApp(targetAppId, newGroupId.ID)
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

// TestClient_Groups_AddToGroup selects a random user, creates a group, and assigns the user to the group
func TestClient_Groups_AddToGroup(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	if err != nil {
		t.Errorf("Error creating client: %v", err)
	}
	// Get all groups, find a group with members
	var groupId string
	groups, err := c.GetAllUserGroups()
	for _, group := range groups {
		m, _ := c.GetGroupMembers(group.ID)
		if len(m) > 0 {
			groupId = group.ID
			break
		}
	}
	// Get group members
	users, err := c.GetGroupMembers(groupId)

	if len(users) == 0 {
		t.Errorf("No users returned")
		t.Errorf("Function Error: %q", err)
	}
	// Random index from users slice
	randomInt := rand.Int() % len(users)

	// Id of random user
	randomUserId := users[randomInt].To.ID

	// Create userGroup
	newGroup, err := c.CreateUserGroup(UserGroup{
		Name:        "sec-jumpcloud-client-go-unit-test-user-association",
		Description: "Created via sec-jumpcloud-client-go unit test, please delete me!",
	})

	ok, _ := c.AddUserToGroup(newGroup.ID, randomUserId)
	if !ok {
		t.Errorf("Unable to add user to group")
	}

	// Create slice of userIds
	var groupMembers []string
	members, _ := c.GetGroupMembers(newGroup.ID)
	for _, m := range members {
		groupMembers = append(groupMembers, m.To.ID)
	}
	// Check for user membership
	result := contains(groupMembers, randomUserId)
	if !result {
		t.Errorf("Unable to find user in group")
	}

	// Remove user from group
	ok2, _ := c.RemoveUserFromGroup(newGroup.ID, randomUserId)
	if !ok2 {
		t.Errorf("Unable to remove user from group")
	}

	// Create slice of userIds
	var groupMembers2 []string
	members, _ = c.GetGroupMembers(newGroup.ID)
	for _, m := range members {
		groupMembers2 = append(groupMembers2, m.To.ID)
	}

	// Check for user membership
	result2 := contains(groupMembers2, randomUserId)
	if result2 {
		t.Errorf("User still a group member after removal")
	}
	// Delete userGroup
	err = c.DeleteUserGroup(newGroup.ID)
}

// TestClient_Groups_SearchForGroup searches for a group by name
func TestClient_Groups_SearchForGroup(t *testing.T) {
	c, err := NewClient(os.Getenv("JC_API_KEY"))
	groups, err := c.SearchUserGroups("name", "a", 3)
	for _, g := range groups {
		fmt.Println(g.Name)
	}
	if len(groups) == 0 {
		t.Errorf("No groups returned")
		t.Errorf("Function Error: %q", err)
	}
}

package snjumpcloud

import (
	"net/http"
	"net/url"
	"time"
)

// JC is our Client API client
type JC struct {
	Url     url.URL
	Headers http.Header
	Client  http.Client
}

// UserGroup from https://docs.jumpcloud.com/api/2.0/index.html#tag/Groups/operation/groups_list
type UserGroup struct {
	Attributes  map[string]string
	Description string
	Email       string
	Id          string
	Name        string
	Type        string
}

// UserGroupDetails Something is broken about this... using map[string]string for now
type UserGroupDetails struct {
	Attributes struct {
		Sudo struct {
			Enabled         bool `json:"enabled"`
			WithoutPassword bool `json:"withoutPassword"`
		} `json:"sudo"`
		LdapGroups []struct {
			Name string `json:"name"`
		} `json:"ldapGroups"`
		PosixGroups []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"posixGroups"`
		Radius struct {
			Reply []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"reply"`
		} `json:"radius"`
		SambaEnabled bool `json:"sambaEnabled"`
	} `json:"attributes"`
	Description string `json:"description"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	MemberQuery struct {
		QueryType string `json:"queryType"`
		Filters   []struct {
			Field    string `json:"field"`
			Operator string `json:"operator"`
			Value    string `json:"value"`
		} `json:"filters"`
	} `json:"memberQuery"`
	MemberQueryExemptions []struct {
		Attributes struct {
		} `json:"attributes"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"memberQueryExemptions"`
	MemberSuggestionsNotify bool   `json:"memberSuggestionsNotify"`
	MembershipMethod        string `json:"membershipMethod"`
	Name                    string `json:"name"`
	SuggestionCounts        struct {
		Add    int `json:"add"`
		Remove int `json:"remove"`
		Total  int `json:"total"`
	} `json:"suggestionCounts"`
	Type string `json:"type"`
}

// User from https://docs.jumpcloud.com/api/1.0/index.html#tag/Systemusers/operation/systemusers_get
type User struct {
	Id                string `json:"_id"`
	AccountLocked     bool   `json:"account_locked"`
	AccountLockedDate string `json:"account_locked_date"`
	Activated         bool   `json:"activated"`
	Addresses         []struct {
		Country         string `json:"country"`
		ExtendedAddress string `json:"extendedAddress"`
		ID              string `json:"id"`
		Locality        string `json:"locality"`
		PoBox           string `json:"poBox"`
		PostalCode      string `json:"postalCode"`
		Region          string `json:"region"`
		StreetAddress   string `json:"streetAddress"`
		Type            string `json:"type"`
	} `json:"addresses"`
	AllowPublicKey bool   `json:"allow_public_key"`
	AlternateEmail string `json:"alternateEmail"`
	Attributes     []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attributes"`
	BadLoginAttempts               int    `json:"badLoginAttempts"`
	Company                        string `json:"company"`
	CostCenter                     string `json:"costCenter"`
	Created                        string `json:"created"`
	CreationSource                 string `json:"creationSource"`
	Department                     string `json:"department"`
	Description                    string `json:"description"`
	DisableDeviceMaxLoginAttempts  bool   `json:"disableDeviceMaxLoginAttempts"`
	Displayname                    string `json:"displayname"`
	Email                          string `json:"email"`
	EmployeeIdentifier             string `json:"employeeIdentifier"`
	EmployeeType                   string `json:"employeeType"`
	EnableManagedUID               bool   `json:"enable_managed_uid"`
	EnableUserPortalMultifactor    bool   `json:"enable_user_portal_multifactor"`
	ExternalDn                     string `json:"external_dn"`
	ExternalPasswordExpirationDate string `json:"external_password_expiration_date"`
	ExternalSourceType             string `json:"external_source_type"`
	ExternallyManaged              bool   `json:"externally_managed"`
	Firstname                      string `json:"firstname"`
	JobTitle                       string `json:"jobTitle"`
	Lastname                       string `json:"lastname"`
	LdapBindingUser                bool   `json:"ldap_binding_user"`
	Location                       string `json:"location"`
	ManagedAppleID                 string `json:"managedAppleId"`
	Manager                        string `json:"manager"`
	Mfa                            struct {
		Configured     bool      `json:"configured"`
		Exclusion      bool      `json:"exclusion"`
		ExclusionDays  int       `json:"exclusionDays"`
		ExclusionUntil time.Time `json:"exclusionUntil"`
	} `json:"mfa"`
	MfaEnrollment struct {
		OverallStatus  string `json:"overallStatus"`
		PushStatus     string `json:"pushStatus"`
		TotpStatus     string `json:"totpStatus"`
		WebAuthnStatus string `json:"webAuthnStatus"`
	} `json:"mfaEnrollment"`
	Middlename             string `json:"middlename"`
	Organization           string `json:"organization"`
	PasswordDate           string `json:"password_date"`
	PasswordExpirationDate string `json:"password_expiration_date"`
	PasswordExpired        bool   `json:"password_expired"`
	PasswordNeverExpires   bool   `json:"password_never_expires"`
	PasswordlessSudo       bool   `json:"passwordless_sudo"`
	PhoneNumbers           []struct {
		ID     string `json:"id"`
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"phoneNumbers"`
	PublicKey     string `json:"public_key"`
	RecoveryEmail struct {
		Address    string `json:"address"`
		Verified   bool   `json:"verified"`
		VerifiedAt string `json:"verifiedAt"`
	} `json:"recoveryEmail"`
	Relationships []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"relationships"`
	RestrictedFields []struct {
		Field string `json:"field"`
		ID    string `json:"id"`
		Type  string `json:"type"`
	} `json:"restrictedFields"`
	SambaServiceUser bool `json:"samba_service_user"`
	SSHKeys          []struct {
		ID         string `json:"_id"`
		CreateDate string `json:"create_date"`
		Name       string `json:"name"`
		PublicKey  string `json:"public_key"`
	} `json:"ssh_keys"`
	State       string   `json:"state"`
	Sudo        bool     `json:"sudo"`
	Suspended   bool     `json:"suspended"`
	Tags        []string `json:"tags"`
	TotpEnabled bool     `json:"totp_enabled"`
	UnixGUID    int      `json:"unix_guid"`
	UnixUID     int      `json:"unix_uid"`
	Username    string   `json:"username"`
}

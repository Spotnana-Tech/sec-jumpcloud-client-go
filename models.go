package jcclient

import (
	"time"
)

// SystemUser Users and their attributes
type SystemUser struct {
	ID                string `json:"_id,omitempty"`
	AccountLocked     bool   `json:"account_locked,omitempty"`
	AccountLockedDate string `json:"account_locked_date,omitempty"`
	Activated         bool   `json:"activated,omitempty"`
	Addresses         []struct {
		Country         string `json:"country,omitempty"`
		ExtendedAddress string `json:"extendedAddress,omitempty"`
		ID              string `json:"id,omitempty"`
		Locality        string `json:"locality,omitempty"`
		PoBox           string `json:"poBox,omitempty"`
		PostalCode      string `json:"postalCode,omitempty"`
		Region          string `json:"region,omitempty"`
		StreetAddress   string `json:"streetAddress,omitempty"`
		Type            string `json:"type,omitempty"`
	} `json:"addresses,omitempty"`
	AllowPublicKey bool   `json:"allow_public_key,omitempty"`
	AlternateEmail string `json:"alternateEmail,omitempty"`
	Attributes     []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"attributes,omitempty"`
	BadLoginAttempts               int    `json:"badLoginAttempts,omitempty"`
	Company                        string `json:"company,omitempty"`
	CostCenter                     string `json:"costCenter,omitempty"`
	Created                        string `json:"created,omitempty"`
	CreationSource                 string `json:"creationSource,omitempty"`
	Department                     string `json:"department,omitempty"`
	Description                    string `json:"description,omitempty"`
	DisableDeviceMaxLoginAttempts  bool   `json:"disableDeviceMaxLoginAttempts,omitempty"`
	Displayname                    string `json:"displayname,omitempty"`
	Email                          string `json:"email,omitempty"`
	EmployeeIdentifier             string `json:"employeeIdentifier,omitempty"`
	EmployeeType                   string `json:"employeeType,omitempty"`
	EnableManagedUID               bool   `json:"enable_managed_uid,omitempty"`
	EnableUserPortalMultifactor    bool   `json:"enable_user_portal_multifactor,omitempty"`
	ExternalDn                     string `json:"external_dn,omitempty"`
	ExternalPasswordExpirationDate string `json:"external_password_expiration_date,omitempty"`
	ExternalSourceType             string `json:"external_source_type,omitempty"`
	ExternallyManaged              bool   `json:"externally_managed,omitempty"`
	Firstname                      string `json:"firstname,omitempty"`
	JobTitle                       string `json:"jobTitle,omitempty"`
	Lastname                       string `json:"lastname,omitempty"`
	LdapBindingUser                bool   `json:"ldap_binding_user,omitempty"`
	Location                       string `json:"location,omitempty"`
	ManagedAppleID                 string `json:"managedAppleId,omitempty"`
	Manager                        string `json:"manager,omitempty"`
	Mfa                            struct {
		Configured     bool      `json:"configured,omitempty"`
		Exclusion      bool      `json:"exclusion,omitempty"`
		ExclusionDays  int       `json:"exclusionDays,omitempty"`
		ExclusionUntil time.Time `json:"exclusionUntil,omitempty"`
	} `json:"mfa,omitempty"`
	MfaEnrollment struct {
		OverallStatus  string `json:"overallStatus,omitempty"`
		PushStatus     string `json:"pushStatus,omitempty"`
		TotpStatus     string `json:"totpStatus,omitempty"`
		WebAuthnStatus string `json:"webAuthnStatus,omitempty"`
	} `json:"mfaEnrollment,omitempty"`
	Middlename             string `json:"middlename,omitempty"`
	Organization           string `json:"organization,omitempty"`
	PasswordDate           string `json:"password_date,omitempty"`
	PasswordExpirationDate string `json:"password_expiration_date,omitempty"`
	PasswordExpired        bool   `json:"password_expired,omitempty"`
	PasswordNeverExpires   bool   `json:"password_never_expires,omitempty"`
	PasswordlessSudo       bool   `json:"passwordless_sudo,omitempty"`
	PhoneNumbers           []struct {
		ID     string `json:"id,omitempty"`
		Number string `json:"number,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"phoneNumbers,omitempty"`
	PublicKey     string `json:"public_key,omitempty"`
	RecoveryEmail struct {
		Address    string `json:"address,omitempty"`
		Verified   bool   `json:"verified,omitempty"`
		VerifiedAt string `json:"verifiedAt,omitempty"`
	} `json:"recoveryEmail,omitempty"`
	Relationships []struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"relationships,omitempty"`
	RestrictedFields []struct {
		Field string `json:"field,omitempty"`
		ID    string `json:"id,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"restrictedFields,omitempty"`
	SambaServiceUser bool `json:"samba_service_user,omitempty"`
	SSHKeys          []struct {
		ID         string `json:"_id,omitempty"`
		CreateDate string `json:"create_date,omitempty"`
		Name       string `json:"name,omitempty"`
		PublicKey  string `json:"public_key,omitempty"`
	} `json:"ssh_keys,omitempty"`
	State       string   `json:"state,omitempty"`
	Sudo        bool     `json:"sudo,omitempty"`
	Suspended   bool     `json:"suspended,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TotpEnabled bool     `json:"totp_enabled,omitempty"`
	UnixGUID    int      `json:"unix_guid,omitempty"`
	UnixUID     int      `json:"unix_uid,omitempty"`
	Username    string   `json:"username,omitempty"`
}
type User struct {
	ID                string `json:"_id,omitempty"`
	AccountLocked     bool   `json:"account_locked,omitempty"`
	AccountLockedDate string `json:"account_locked_date,omitempty"`
	Activated         bool   `json:"activated,omitempty"`
	Addresses         []struct {
		Country         string `json:"country,omitempty"`
		ExtendedAddress string `json:"extendedAddress,omitempty"`
		ID              string `json:"id,omitempty"`
		Locality        string `json:"locality,omitempty"`
		PoBox           string `json:"poBox,omitempty"`
		PostalCode      string `json:"postalCode,omitempty"`
		Region          string `json:"region,omitempty"`
		StreetAddress   string `json:"streetAddress,omitempty"`
		Type            string `json:"type,omitempty"`
	} `json:"addresses,omitempty"`
	AllowPublicKey bool   `json:"allow_public_key,omitempty"`
	AlternateEmail string `json:"alternateEmail,omitempty"`
	Attributes     []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"attributes,omitempty"`
	BadLoginAttempts               int    `json:"badLoginAttempts,omitempty"`
	Company                        string `json:"company,omitempty"`
	CostCenter                     string `json:"costCenter,omitempty"`
	Created                        string `json:"created,omitempty"`
	CreationSource                 string `json:"creationSource,omitempty"`
	Department                     string `json:"department,omitempty"`
	Description                    string `json:"description,omitempty"`
	DisableDeviceMaxLoginAttempts  bool   `json:"disableDeviceMaxLoginAttempts,omitempty"`
	Displayname                    string `json:"displayname,omitempty"`
	Email                          string `json:"email,omitempty"`
	EmployeeIdentifier             string `json:"employeeIdentifier,omitempty"`
	EmployeeType                   string `json:"employeeType,omitempty"`
	EnableManagedUID               bool   `json:"enable_managed_uid,omitempty"`
	EnableUserPortalMultifactor    bool   `json:"enable_user_portal_multifactor,omitempty"`
	ExternalDn                     string `json:"external_dn,omitempty"`
	ExternalPasswordExpirationDate string `json:"external_password_expiration_date,omitempty"`
	ExternalSourceType             string `json:"external_source_type,omitempty"`
	ExternallyManaged              bool   `json:"externally_managed,omitempty"`
	Firstname                      string `json:"firstname,omitempty"`
	JobTitle                       string `json:"jobTitle,omitempty"`
	Lastname                       string `json:"lastname,omitempty"`
	LdapBindingUser                bool   `json:"ldap_binding_user,omitempty"`
	Location                       string `json:"location,omitempty"`
	ManagedAppleID                 string `json:"managedAppleId,omitempty"`
	Manager                        string `json:"manager,omitempty"`
	Mfa                            struct {
		Configured     bool      `json:"configured,omitempty"`
		Exclusion      bool      `json:"exclusion,omitempty"`
		ExclusionDays  int       `json:"exclusionDays,omitempty"`
		ExclusionUntil time.Time `json:"exclusionUntil,omitempty"`
	} `json:"mfa,omitempty"`
	MfaEnrollment struct {
		OverallStatus  string `json:"overallStatus,omitempty"`
		PushStatus     string `json:"pushStatus,omitempty"`
		TotpStatus     string `json:"totpStatus,omitempty"`
		WebAuthnStatus string `json:"webAuthnStatus,omitempty"`
	} `json:"mfaEnrollment,omitempty"`
	Middlename             string `json:"middlename,omitempty"`
	Organization           string `json:"organization,omitempty"`
	PasswordDate           string `json:"password_date,omitempty"`
	PasswordExpirationDate string `json:"password_expiration_date,omitempty"`
	PasswordExpired        bool   `json:"password_expired,omitempty"`
	PasswordNeverExpires   bool   `json:"password_never_expires,omitempty"`
	PasswordlessSudo       bool   `json:"passwordless_sudo,omitempty"`
	PhoneNumbers           []struct {
		ID     string `json:"id,omitempty"`
		Number string `json:"number,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"phoneNumbers,omitempty"`
	PublicKey     string `json:"public_key,omitempty"`
	RecoveryEmail struct {
		Address    string `json:"address,omitempty"`
		Verified   bool   `json:"verified,omitempty"`
		VerifiedAt string `json:"verifiedAt,omitempty"`
	} `json:"recoveryEmail,omitempty"`
	Relationships []struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"relationships,omitempty"`
	RestrictedFields []struct {
		Field string `json:"field,omitempty"`
		ID    string `json:"id,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"restrictedFields,omitempty"`
	SambaServiceUser bool `json:"samba_service_user,omitempty"`
	SSHKeys          []struct {
		ID         string `json:"_id,omitempty"`
		CreateDate string `json:"create_date,omitempty"`
		Name       string `json:"name,omitempty"`
		PublicKey  string `json:"public_key,omitempty"`
	} `json:"ssh_keys,omitempty"`
	State       string   `json:"state,omitempty"`
	Sudo        bool     `json:"sudo,omitempty"`
	Suspended   bool     `json:"suspended,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TotpEnabled bool     `json:"totp_enabled,omitempty"`
	UnixGUID    int      `json:"unix_guid,omitempty"`
	UnixUID     int      `json:"unix_uid,omitempty"`
	Username    string   `json:"username,omitempty"`
}

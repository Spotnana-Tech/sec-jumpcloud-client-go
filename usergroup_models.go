package jumpcloud

// UserGroup is sensitive due to the null values being empty structs when calling the API
type UserGroup struct {
	Attributes              *Attributes              `json:"attributes,omitempty"`
	Description             string                   `json:"description,omitempty"`
	Email                   string                   `json:"email,omitempty"`
	ID                      string                   `json:"id,omitempty"`
	MemberQuery             *MemberQuery             `json:"memberQuery,omitempty"`
	MemberQueryExemptions   *[]MemberQueryExemptions `json:"memberQueryExemptions,omitempty"`
	MemberSuggestionsNotify bool                     `json:"memberSuggestionsNotify,omitempty"`
	MembershipMethod        string                   `json:"membershipMethod,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	SuggestionCounts        *SuggestionCounts        `json:"suggestionCounts,omitempty"`
	Type                    string                   `json:"type,omitempty"`
}
type Sudo struct {
	Enabled         bool `json:"enabled,omitempty"`
	WithoutPassword bool `json:"withoutPassword,omitempty"`
}
type LdapGroups struct {
	Name string `json:"name,omitempty"`
}
type PosixGroups struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Reply struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type Radius struct {
	Reply []Reply `json:"reply,omitempty"`
}
type Attributes struct {
	Sudo         *Sudo          `json:"sudo,omitempty"`
	LdapGroups   *[]LdapGroups  `json:"ldapGroups,omitempty"`
	PosixGroups  *[]PosixGroups `json:"posixGroups,omitempty"`
	Radius       *Radius        `json:"radius,omitempty"`
	SambaEnabled bool           `json:"sambaEnabled,omitempty"`
}
type Filters struct {
	Field    string `json:"field,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}
type MemberQuery struct {
	QueryType string     `json:"queryType,omitempty"`
	Filters   *[]Filters `json:"filters,omitempty"`
}
type MemberQueryExemptions struct {
	Attributes *Attributes `json:"attributes,omitempty"`
	ID         string      `json:"id,omitempty"`
	Type       string      `json:"type,omitempty"`
}
type SuggestionCounts struct {
	Add    int `json:"add,omitempty"`
	Remove int `json:"remove,omitempty"`
	Total  int `json:"total,omitempty"`
}

// UserGroups is a slice of UserGroup
type UserGroups []UserGroup

// GroupMembership is a slice of struct containing members of a group
type GroupMembership []struct {
	Attributes struct {
	} `json:"attributes,omitempty"`
	From struct {
		Attributes struct {
		} `json:"attributes,omitempty"`
		ID   string `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"from,omitempty"`
	To struct {
		Attributes struct {
		} `json:"attributes,omitempty"`
		ID   string `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"to,omitempty"`
}

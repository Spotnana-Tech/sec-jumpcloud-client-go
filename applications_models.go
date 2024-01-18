package jcclient

type AllApps []App

type App struct {
	ID           string `json:"_id,omitempty"`
	Name         string `json:"name,omitempty"`
	CatalogItem  any    `json:"catalogItem,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	DisplayLabel string `json:"displayLabel,omitempty"`
	Description  string `json:"description,omitempty"`
	Color        any    `json:"color,omitempty"`
	Logo         struct {
		URL string `json:"url,omitempty"`
	} `json:"logo,omitempty"`
	Provision any `json:"provision,omitempty"`
	Sso       struct {
		Bookmark struct {
			URL string `json:"url,omitempty"`
		} `json:"bookmark,omitempty"`
		Type        string `json:"type,omitempty"`
		SpErrorFlow bool   `json:"spErrorFlow,omitempty"`
		Hidden      bool   `json:"hidden,omitempty"`
		Active      bool   `json:"active,omitempty"`
		Beta        bool   `json:"beta,omitempty"`
		URL         string `json:"url,omitempty"`
		Jit         struct {
			Supported bool `json:"supported,omitempty"`
			Enabled   bool `json:"enabled,omitempty"`
		} `json:"jit,omitempty"`
	} `json:"sso,omitempty"`
	Status       string `json:"status,omitempty"`
	Organization string `json:"organization,omitempty"`
}

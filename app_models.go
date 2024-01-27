package jcclient

import "time"

// AllApps Apps and their attributes
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
type AppDetail struct {
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
		Hidden                  bool      `json:"hidden,omitempty"`
		IdpCertificateUpdatedAt time.Time `json:"idpCertificateUpdatedAt,omitempty"`
		IdpPrivateKeyUpdatedAt  time.Time `json:"idpPrivateKeyUpdatedAt,omitempty"`
		IdpCertExpirationAt     time.Time `json:"idpCertExpirationAt,omitempty"`
		Beta                    bool      `json:"beta,omitempty"`
		Type                    string    `json:"type,omitempty"`
		URL                     string    `json:"url,omitempty"`
		Active                  bool      `json:"active,omitempty"`
		CertExpired             bool      `json:"certExpired,omitempty"`
		Jit                     struct {
			Supported bool `json:"supported,omitempty"`
			Enabled   bool `json:"enabled,omitempty"`
		} `json:"jit,omitempty"`
	} `json:"sso,omitempty"`
	Status       string `json:"status,omitempty"`
	Organization string `json:"organization,omitempty"`
}
type AppAssociations []struct {
	Attributes any `json:"attributes,omitempty"`
	To         struct {
		Attributes struct {
			LdapGroups []struct {
				Name string `json:"name,omitempty"`
			} `json:"ldapGroups,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"to,omitempty"`
}

// NewApp is the structure of a new SSO app request
type NewApp struct {
	ID     string `json:"_id,omitempty"`
	Active bool   `json:"active,omitempty"`
	Beta   bool   `json:"beta,omitempty"`
	Color  string `json:"color,omitempty"`
	Config struct {
		AcsURL struct {
			Label    string `json:"label,omitempty"`
			Options  string `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   string `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   string `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"acsUrl,omitempty"`
		ConstantAttributes struct {
			Label    string `json:"label,omitempty"`
			Mutable  bool   `json:"mutable,omitempty"`
			Options  any    `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   any    `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type  string `json:"type,omitempty"`
			Value []struct {
				Name     string `json:"name,omitempty"`
				ReadOnly bool   `json:"readOnly,omitempty"`
				Required bool   `json:"required,omitempty"`
				Value    string `json:"value,omitempty"`
				Visible  bool   `json:"visible,omitempty"`
			} `json:"value,omitempty"`
			Visible bool `json:"visible,omitempty"`
		} `json:"constantAttributes,omitempty"`
		DatabaseAttributes struct {
			Position int `json:"position,omitempty"`
		} `json:"databaseAttributes,omitempty"`
		IdpCertificate struct {
			Label    string `json:"label,omitempty"`
			Options  string `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   string `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   string `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"idpCertificate,omitempty"`
		IdpEntityID struct {
			Label    string `json:"label,omitempty"`
			Options  string `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   string `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   string `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"idpEntityId,omitempty"`
		IdpPrivateKey struct {
			Label    string `json:"label,omitempty"`
			Options  string `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   string `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   string `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"idpPrivateKey,omitempty"`
		SignAssertion struct {
			Label    string `json:"label,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   bool   `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"signAssertion,omitempty"`
		SignResponse struct {
			Label    string `json:"label,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   bool   `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"signResponse,omitempty"`
		SpEntityID struct {
			Label    string `json:"label,omitempty"`
			Options  string `json:"options,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Toggle   string `json:"toggle,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   string `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"spEntityId,omitempty"`
		SpErrorFlow struct {
			Label    string `json:"label,omitempty"`
			Position int    `json:"position,omitempty"`
			ReadOnly bool   `json:"readOnly,omitempty"`
			Required bool   `json:"required,omitempty"`
			Tooltip  struct {
				Template  string `json:"template,omitempty"`
				Variables struct {
					Icon    string `json:"icon,omitempty"`
					Message string `json:"message,omitempty"`
				} `json:"variables,omitempty"`
			} `json:"tooltip,omitempty"`
			Type    string `json:"type,omitempty"`
			Value   bool   `json:"value,omitempty"`
			Visible bool   `json:"visible,omitempty"`
		} `json:"spErrorFlow,omitempty"`
	} `json:"config,omitempty"`
	Created            string `json:"created,omitempty"`
	DatabaseAttributes []struct {
	} `json:"databaseAttributes,omitempty"`
	Description  string `json:"description,omitempty"`
	DisplayLabel string `json:"displayLabel,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	LearnMore    string `json:"learnMore,omitempty"`
	Logo         struct {
		Color string `json:"color,omitempty"`
		URL   string `json:"url,omitempty"`
	} `json:"logo,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Sso          struct {
		Beta                bool      `json:"beta,omitempty"`
		Hidden              bool      `json:"hidden,omitempty"`
		IdpCertExpirationAt time.Time `json:"idpCertExpirationAt,omitempty"`
		Jit                 bool      `json:"jit,omitempty"`
		Type                string    `json:"type,omitempty"`
	} `json:"sso,omitempty"`
	SsoURL string `json:"ssoUrl,omitempty"`
}

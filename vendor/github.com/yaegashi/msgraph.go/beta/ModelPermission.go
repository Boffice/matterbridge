// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Permission undocumented
type Permission struct {
	// Entity is the base model of Permission
	Entity
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// GrantedTo undocumented
	GrantedTo *IdentitySet `json:"grantedTo,omitempty"`
	// GrantedToIdentities undocumented
	GrantedToIdentities []IdentitySet `json:"grantedToIdentities,omitempty"`
	// HasPassword undocumented
	HasPassword *bool `json:"hasPassword,omitempty"`
	// InheritedFrom undocumented
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`
	// Invitation undocumented
	Invitation *SharingInvitation `json:"invitation,omitempty"`
	// Link undocumented
	Link *SharingLink `json:"link,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
}

// PermissionScope undocumented
type PermissionScope struct {
	// Object is the base model of PermissionScope
	Object
	// AdminConsentDescription undocumented
	AdminConsentDescription *string `json:"adminConsentDescription,omitempty"`
	// AdminConsentDisplayName undocumented
	AdminConsentDisplayName *string `json:"adminConsentDisplayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Origin undocumented
	Origin *string `json:"origin,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// UserConsentDescription undocumented
	UserConsentDescription *string `json:"userConsentDescription,omitempty"`
	// UserConsentDisplayName undocumented
	UserConsentDisplayName *string `json:"userConsentDisplayName,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

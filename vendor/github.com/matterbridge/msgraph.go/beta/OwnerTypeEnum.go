// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OwnerType undocumented
type OwnerType int

const (
	// OwnerTypeVUnknown undocumented
	OwnerTypeVUnknown OwnerType = 0
	// OwnerTypeVCompany undocumented
	OwnerTypeVCompany OwnerType = 1
	// OwnerTypeVPersonal undocumented
	OwnerTypeVPersonal OwnerType = 2
)

// OwnerTypePUnknown returns a pointer to OwnerTypeVUnknown
func OwnerTypePUnknown() *OwnerType {
	v := OwnerTypeVUnknown
	return &v
}

// OwnerTypePCompany returns a pointer to OwnerTypeVCompany
func OwnerTypePCompany() *OwnerType {
	v := OwnerTypeVCompany
	return &v
}

// OwnerTypePPersonal returns a pointer to OwnerTypeVPersonal
func OwnerTypePPersonal() *OwnerType {
	v := OwnerTypeVPersonal
	return &v
}
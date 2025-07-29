package utils

type Permission int

const (
	AdminOnly Permission = 1 << iota
	CreatePublic
	ReadPublic
	UpdatePublic
	DeletePublic
	CreateProtected
	ReadProtected
	UpdateProtected
	DeleteProtected
)

var permissionText = map[Permission]string{
	AdminOnly:       "Admin",
	CreatePublic:    "Create Public",
	ReadPublic:      "Read Public",
	UpdatePublic:    "Update Public",
	DeletePublic:    "Delete Public",
	CreateProtected: "Create Protected",
	ReadProtected:   "Read Protected",
	UpdateProtected: "Update Protected",
	DeleteProtected: "Delete Protected",
}

func (p Permission) String() string {
	return permissionText[p]
}

func fullPermission() Permission {
	var perm Permission = 0
	for k, _ := range permissionText {
		perm += k
	}

	return perm
}

func publicCRUD() Permission {
	return CreatePublic | ReadPublic | UpdatePublic | DeletePublic
}

func protectedCRUD() Permission {
	return CreateProtected | ReadProtected | UpdateProtected | DeleteProtected
}

var RolePermissions = map[string]Permission{
	"admin":  fullPermission(),
	"user":   publicCRUD() | protectedCRUD(),
	"viewer": ReadPublic | ReadProtected,
}

func HasPermission(role string, permission Permission) bool {
	perms, exists := RolePermissions[role]
	if !exists {
		return false
	}

	return perms&permission == permission
}

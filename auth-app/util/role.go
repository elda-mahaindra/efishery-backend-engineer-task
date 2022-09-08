package util

// Constants for all supported roles
const (
	SUPER = "super"
	ADMIN = "admin"
	USER  = "user"
)

var AvailableRoles = []string{"super", "admin", "user"}

// IsSupportedRole returns true if the role is supported
func IsSupportedRole(role string) bool {
	switch role {
	case SUPER, ADMIN, USER:
		return true
	}
	return false
}

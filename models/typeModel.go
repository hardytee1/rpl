package models

// Define a custom type for Role
type Role string

// Define the allowed roles as constants
const (
	RoleUser    Role = "USER"
	RoleTeacher Role = "TEACHER"
	RoleAdmin   Role = "ADMIN"
)

// ValidateRole checks if a given role is valid
func ValidateRole(role Role) bool {
	switch role {
	case RoleUser, RoleTeacher, RoleAdmin:
		return true
	default:
		return false
	}
}

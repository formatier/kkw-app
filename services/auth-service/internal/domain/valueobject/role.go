package valueobject

type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
)

func NewRole(val string) (Role, bool) {
	role := Role(val)
	switch role {
	case RoleStudent, RoleTeacher:
		return role, true
	default:
		return role, false
	}
}

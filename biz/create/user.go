package create

type User struct {
	Username string
	Password string
	Role     RoleType
}

type RoleType int

const (
	Customer RoleType = iota
	Collaborator
	Manager
	Administrator
)

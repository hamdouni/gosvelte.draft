package user

// User définie un utilisateur par son identifiant, son mot de passe et son rôle
type User struct {
	Username string
	Password string
	Role     Role
}

// Role définie le role que peut avoir un user
type Role int

const (
	Customer Role = iota
	Collaborator
	Manager
	Administrator
)

// New retourne un utilisateur validé
func New(un, pw string, ro Role) (*User, error) {
	if len(un) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(pw) < 4 {
		return nil, ErrPasswordTooShort
	}
	if ro < Customer || ro > Administrator {
		return nil, ErrUndefinedRole
	}

	u := &User{
		Username: un,
		Password: pw,
		Role:     ro,
	}
	return u, nil
}

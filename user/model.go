package user

// Credential définie un utilisateur par son identifiant, son mot de passe et son rôle dans un royaume (realm) donné
type Credential struct {
	Username string
	Password string
	Role     Role
	Realm    string
}

// Role définie le role que peut avoir un user
type Role int

const (
	Anomymous Role = iota
	Customer
	Collaborator
	Manager
	Administrator
)

// New retourne un utilisateur validé
func New(realm, username, password string, role Role) (*Credential, error) {
	if len(username) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(password) < 4 {
		return nil, ErrPasswordTooShort
	}
	if role < Customer || role > Administrator {
		return nil, ErrUndefinedRole
	}

	u := &Credential{
		Username: username,
		Password: password,
		Role:     role,
		Realm:    realm,
	}
	return u, nil
}

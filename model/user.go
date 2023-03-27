package model

import "admin/model/secure"

// User définie un utilisateur par son identifiant, son mot de passe et son rôle
type User struct {
	Username string
	Password string
	Role     Role
}

// UserStore est le système de stockage des utilisateurs
var UserStore UserStorage

// Contrat avec le service de stockage
type UserStorage interface {
	GetPasswordUser(username string) (encryptedPassword string)
	AddUser(user User) error
}

// NewUser retourne un utilisateur validé
func NewUser(un, pw string, ro Role) (*User, error) {
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

// AddUser enregistre un utilisateur dans le système de stockage
func AddUser(username, password string, role Role) error {
	hashed, err := secure.HashPassword(password)
	if err != nil {
		return err
	}
	u, err := NewUser(username, hashed, role)
	if err != nil {
		return err
	}
	UserStore.AddUser(*u)

	return nil
}

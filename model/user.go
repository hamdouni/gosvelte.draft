package model

import "admin/model/secure"

type User struct {
	Username string
	Password string
	Role     Role
}

var UserStore UserStorage

// Contrat avec le service de stockage
type UserStorage interface {
	GetPasswordUser(username string) (encryptedPassword string)
	AddUser(username, password string, role Role) error
}

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

func CheckPassword(username, password string) bool {
	hashed := UserStore.GetPasswordUser(username)

	return secure.CheckPassword(password, hashed)
}

func AddUser(username, password string, role Role) error {
	hashed, err := secure.HashPassword(password)
	if err != nil {
		return err
	}
	UserStore.AddUser(username, hashed, role)

	return nil
}

func Decrypt(message string) (string, error) {
	return secure.Decrypt(message)
}

func Encrypt(message string) (string, error) {
	return secure.Encrypt(message)
}

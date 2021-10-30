package usecase

import (
	"errors"
)

type User struct{ Username, Password string }

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
)

func NewUser(us, pw string) (*User, error) {
	if len(us) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(pw) < 8 {
		return nil, ErrPasswordTooShort
	}

	u := &User{
		Username: us,
		Password: pw,
	}
	return u, nil
}

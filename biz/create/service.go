package create

import (
	"app/biz/hash"
	"errors"
	"fmt"
)

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
)

func NewUser(us, pw string) (*User, error) {
	if len(us) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(pw) < 4 {
		return nil, ErrPasswordTooShort
	}

	hpw, err := hash.HashPassword(pw)
	if err != nil {
		return nil, fmt.Errorf("biz.create_user.HassPassword(%v): %v", pw, err)
	}

	u := &User{
		Username: us,
		Password: hpw,
	}
	return u, nil
}

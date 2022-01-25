package app

import (
	"app/hash"
	"errors"
	"fmt"
)

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

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
	ErrUndefinedRole    = errors.New("role undefined")
)

func NewUser(us, pw string, role RoleType) (*User, error) {
	if len(us) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(pw) < 4 {
		return nil, ErrPasswordTooShort
	}
	if role < Customer || role > Administrator {
		return nil, ErrUndefinedRole
	}

	hpw, err := hash.HashPassword(pw)
	if err != nil {
		return nil, fmt.Errorf("biz.create_user.HashPassword(%v): %v", pw, err)
	}

	u := &User{
		Username: us,
		Password: hpw,
		Role:     role,
	}
	return u, nil
}

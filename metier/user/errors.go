package user

import "errors"

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
	ErrUndefinedRole    = errors.New("role undefined")
)

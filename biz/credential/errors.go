package credential

import "errors"

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
	ErrUsernameUsed     = errors.New("username already used")
	ErrUndefinedRole    = errors.New("role undefined")
	ErrNotAuthorized    = errors.New("not authorized")
)

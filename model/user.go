package model

type User struct {
	Username string
	Password string
	Role     Role
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

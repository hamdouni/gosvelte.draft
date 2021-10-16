package biz

import "errors"

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrUsernameTooShort = errors.New("username too short")
)

func (b BIZ) CreateUser(us, pw string) error {
	if len(us) < 4 {
		return ErrUsernameTooShort
	}
	if len(pw) < 8 {
		return ErrPasswordTooShort
	}
	encryptedPassword, err := b.encryptPassword(pw)
	if err != nil {
		return err // Unexpected encrypt error
	}
	b.store.AddUser(us, encryptedPassword)
	return nil
}

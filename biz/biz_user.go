package biz

import "errors"

func (b BIZ) CreateUser(us, pw string) error {
	if len(us) < 8 || len(pw) < 8 {
		return errors.New("username or password too short")
	}
	b.store.AddUser(us, pw)
	return nil
}

func (b BIZ) Login(us, pw string) bool {
	if len(us) < 8 || len(pw) < 8 {
		return false
	}
	return pw == b.store.GetPasswordUser(us)
}

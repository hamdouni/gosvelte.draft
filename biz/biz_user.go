package biz

import "errors"

func (b BIZ) CreateUser(us, pw string) error {
	if len(us) < 8 || len(pw) < 8 {
		return errors.New("username or password too short")
	}
	encryptedPassword, err := b.encryptPassword(pw)
	if err != nil {
		return errors.New("error encrypting password in CreateUser " + err.Error())
	}
	b.store.AddUser(us, encryptedPassword)
	return nil
}

func (b BIZ) CheckPassword(us, pw string) bool {
	if len(us) < 8 || len(pw) < 8 {
		return false
	}
	return doPasswordMatch(b.store.GetPasswordUser(us), pw)
}

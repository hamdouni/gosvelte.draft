package web_test

import (
	"admin/app"
	"admin/model"
	"admin/store"
	"admin/web"
)

func init() {
	var fakeSec fakeSecurity
	fakeStore, _ := store.New()
	app.Config(&fakeStore, &fakeStore, fakeSec)
	model.AddUser("samething", "samething", 1)
	web.New(".")
}

type fakeSecurity struct{}

func (f fakeSecurity) Encrypt(s string) (string, error) {
	return s, nil
}
func (f fakeSecurity) Decrypt(s string) (string, error) {
	return s, nil
}

func (f fakeSecurity) HashPassword(pw string) (encryptedPassword string, err error) {
	return pw, nil
}
func (f fakeSecurity) CheckPassword(pw, hashed string) bool {
	return pw == "samething"
}

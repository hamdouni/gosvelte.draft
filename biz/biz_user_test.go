package biz_test

import "testing"

func TestCreateUserWrongLength(t *testing.T) {
	us := "username"
	pw := "xx"
	err := fakeBiz.CreateUser(us, pw)
	if err == nil {
		t.Fatalf("Waiting error but got nil")
	}
	msgerr := "username or password too short"
	if err.Error() != msgerr {
		t.Fatalf("Waiting error \"%v\" but got %v", msgerr, err.Error())
	}
}
func TestCreateUser(t *testing.T) {
	us := "username"
	pw := "password"
	err := fakeBiz.CreateUser(us, pw)
	if err != nil {
		t.Fatalf("Waiting no error but got %v", err)
	}
}

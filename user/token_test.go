package user_test

import (
	"testing"


	"wtk/user"
)

func TestToken(t *testing.T) {
	token, err := user.Token("fakeuser", "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	ok, err := user.CheckToken(token, "otherlocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if ok {
		t.Fatalf("Expecting check token to return false got %v", ok)
	}
	ok, err = user.CheckToken(token, "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if !ok {
		t.Fatalf("Expecting check token to return true got %v", ok)
	}
}

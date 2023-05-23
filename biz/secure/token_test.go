package secure_test

import (
	"testing"
	"wtk/biz"
	"wtk/biz/secure"
)

func TestToken(t *testing.T) {
	biz.Intialize(nil)
	token, err := secure.NewToken("fakeuser", "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	ok, err := secure.CheckToken(token, "otherlocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if ok {
		t.Fatalf("Expecting check token to return false got %v", ok)
	}
	ok, err = secure.CheckToken(token, "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if !ok {
		t.Fatalf("Expecting check token to return true got %v", ok)
	}
}

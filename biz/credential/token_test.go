package credential_test

import (
	"testing"
	"wtk/biz/credential"
)

func TestToken(t *testing.T) {
	token, err := credential.Token("fakeuser", "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	ok, err := credential.CheckToken(token, "otherlocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if ok {
		t.Fatalf("Expecting check token to return false got %v", ok)
	}
	ok, err = credential.CheckToken(token, "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if !ok {
		t.Fatalf("Expecting check token to return true got %v", ok)
	}
}

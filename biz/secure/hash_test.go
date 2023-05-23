package secure_test

import (
	"testing"
	"wtk/biz/secure"
)

func TestHash(t *testing.T) {
	err := secure.Init()
	if err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	testpass := "CKVght324!"

	hashedtext, err := secure.HashPassword(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !secure.CheckPassword(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

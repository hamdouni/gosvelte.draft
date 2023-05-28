package secure_test

import (
	"testing"
	"wtk/ext/secure"
)

func TestHash(t *testing.T) {
	sec, err := secure.New()
	if err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	testpass := "CKVght324!"

	hashedtext, err := sec.Hash(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !sec.CheckHash(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

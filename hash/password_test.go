package hash_test

import (
	"app/hash"
	"testing"
)

func TestHash(t *testing.T) {

	testpass := "CKVght324!"

	hashedtext, err := hash.HashPassword(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !hash.CheckPassword(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

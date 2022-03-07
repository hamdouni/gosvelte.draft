package secure_test

import (
	"admin/secure"
	"testing"
)

func TestHash(t *testing.T) {
	s, err := secure.New()
	if err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	testpass := "CKVght324!"

	hashedtext, err := s.HashPassword(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !s.CheckPassword(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

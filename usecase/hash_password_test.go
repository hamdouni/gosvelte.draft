package usecase_test

import (
	"app/usecase"
	"testing"
)

func TestHash(t *testing.T) {

	testpass := "CKVght324!"

	hashedtext, err := usecase.HashPassword(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !usecase.CheckPassword(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

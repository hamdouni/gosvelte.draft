package biz_test

import (
	"app/biz"
	"testing"
)

func TestHash(t *testing.T) {

	testpass := "CKVght324!"

	hashedtext, err := biz.HashPassword(testpass)
	if err != nil {
		t.Errorf("cannot hash: %s", err)
	}
	if !biz.CheckPassword(testpass, hashedtext) {
		t.Errorf("expected check password to be true")
	}
}

package secure_test

import (
	"testing"
)

// Robert C. Martin - 2019
var ciphertext = "Only test the parts of the application that you want to work"

func TestCrypt(t *testing.T) {
	s := &sec.Secure{}
	if err := s.Init(); err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	cryptedtext, err := s.Encrypt(ciphertext)
	if err != nil {
		t.Errorf("cannot crypt: %s", err)
	}
	decriptedtext, err := s.Decrypt(cryptedtext)
	if err != nil {
		t.Errorf("cannot decrypt: %s", err)
	}
	if ciphertext != decriptedtext {
		t.Errorf("expected %s got %s", ciphertext, decriptedtext)
	}
}

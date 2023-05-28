package secure

import (
	"testing"
)

// Robert C. Martin - 2019
var ciphertext = "Only test the parts of the application that you want to work"

func TestCrypt(t *testing.T) {
	sec, err := New()
	if err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	cryptedtext, err := sec.Encrypt(ciphertext)
	if err != nil {
		t.Errorf("cannot crypt: %s", err)
	}
	decriptedtext, err := sec.Decrypt(cryptedtext)
	if err != nil {
		t.Errorf("cannot decrypt: %s", err)
	}
	if ciphertext != decriptedtext {
		t.Errorf("expected %s got %s", ciphertext, decriptedtext)
	}
}

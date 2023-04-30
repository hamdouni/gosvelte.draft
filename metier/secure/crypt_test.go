package secure_test

import (
	"testing"
	"webtoolkit/metier/secure"
)

// Robert C. Martin - 2019
var ciphertext = "Only test the parts of the application that you want to work"

func TestCrypt(t *testing.T) {
	err := secure.Init()
	if err != nil {
		t.Errorf("cannot initialize security: %s", err)
	}

	cryptedtext, err := secure.Encrypt(ciphertext)
	if err != nil {
		t.Errorf("cannot crypt: %s", err)
	}
	decriptedtext, err := secure.Decrypt(cryptedtext)
	if err != nil {
		t.Errorf("cannot decrypt: %s", err)
	}
	if ciphertext != decriptedtext {
		t.Errorf("expected %s got %s", ciphertext, decriptedtext)
	}
}

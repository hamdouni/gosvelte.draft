package secure

import (
	"crypto/rand"
	"io"
)

// Secure store key
var Secure struct {
	key *[32]byte
}

// Init initialize secret key
func Init() error {
	key, err := newEncryptionKey()
	if err != nil {
		return err
	}
	Secure.key = key
	return nil
}

// newEncryptionKey generates a random 256-bit key
func newEncryptionKey() (*[32]byte, error) {
	key := [32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}
	return &key, nil
}

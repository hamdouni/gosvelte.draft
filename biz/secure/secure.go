package secure

import (
	"crypto/rand"
	"io"
)

// securestore store key
var securestore struct {
	key *[32]byte
}

// Init initialize secret key
func Init() error {
	key, err := newEncryptionKey()
	if err != nil {
		return err
	}
	securestore.key = key
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

package secure

import (
	"crypto/rand"
	"io"
)

// Secure store key
type Secure struct {
	key *[32]byte
}

// New initialize secret key
func New() (Secure, error) {
	key, err := newEncryptionKey()
	if err != nil {
		return Secure{}, err
	}
	return Secure{key: key}, nil
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

package secure

import (
	"crypto/rand"
	"io"
)

type keyType [32]byte

type SecureStore struct {
	key *keyType
}

func New() (sec SecureStore, err error) {
	key := keyType{}
	_, err = io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return sec, err
	}
	sec.key = &key
	return sec, nil
}

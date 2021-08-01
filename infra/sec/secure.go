// Package sec provides symmetric authenticated encryption using 256-bit AES-GCM with a random nonce.
// From cryptopasta - Written by George Tankersley <george.tankersley@gmail.com>
package sec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Secure store key
type Secure struct {
	key *[32]byte
}

// Init initialize secret key
func (s *Secure) Init() (err error) {
	s.key, err = newEncryptionKey()
	return err
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

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func (s Secure) Encrypt(plaintext string) (ciphertext string, err error) {
	block, err := aes.NewCipher(s.key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	sealed := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return string(sealed), nil
}

// Decrypt decrypts data using 256-bit AES-GCM. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func (s Secure) Decrypt(ciphertext string) (plaintext string, err error) {
	block, err := aes.NewCipher(s.key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}

	cipher := []byte(ciphertext)
	unsealed, err := gcm.Open(nil,
		cipher[:gcm.NonceSize()],
		cipher[gcm.NonceSize():],
		nil,
	)
	if err != nil {
		return "", err
	}
	return string(unsealed), nil
}

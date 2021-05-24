package biz

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// Encrypt password using bcrypt
func (b BIZ) encryptPassword(pw string) (ep string, err error) {
	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// Convert the hashed password to a base64 encoded string
	ep = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return ep, nil
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func doPasswordMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err != nil
}

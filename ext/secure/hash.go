package secure

import "golang.org/x/crypto/bcrypt"

// Hash using Bcrypt
func (sec SecureStore) Hash(pw string) (encryptedPassword string, err error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

// CheckHash return true if hash(pw)=hashed
func (sec SecureStore) CheckHash(pw, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
	return err == nil
}

package sec

import "golang.org/x/crypto/bcrypt"

// CheckPassword return true if hash(pw)=hashed
func (s *Secure) CheckPassword(pw, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
	return err == nil
}

// Hash password using Bcrypt
func (s *Secure) HashPassword(pw string) (encryptedPassword string, err error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

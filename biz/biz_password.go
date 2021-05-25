package biz

import "golang.org/x/crypto/bcrypt"

func (b BIZ) CheckPassword(us, pw string) bool {
	if len(us) < 8 || len(pw) < 8 {
		return false
	}
	return doPasswordMatch(b.store.GetPasswordUser(us), pw)
}

// Encrypt password using Bcrypt
func (b BIZ) encryptPassword(pw string) (encryptedPassword string, err error) {
	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
func doPasswordMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	// CompareHashAndPassword return nil on success and an error on failure.
	return err == nil
}

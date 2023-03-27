package model

import (
	"admin/model/secure"
	"fmt"
)

// Auth vérifie l'authentification de l'utilisateur et retourne un token signé.
func Auth(user, pass, address string) (string, error) {
	if !ValidatePassword(user, pass) {
		return "", fmt.Errorf("not authorized")
	}
	t, err := secure.NewToken(user, address)
	if err != nil {
		return "", err
	}
	return t, nil
}

// ValidatePassword vérifie la conformité de l'identifiant et mot de passe
func ValidatePassword(username, password string) bool {
	hashed := UserStore.GetPasswordUser(username)

	return secure.CheckPassword(password, hashed)
}

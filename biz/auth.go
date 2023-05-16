package biz

import (
	"fmt"
	"webtoolkit/biz/secure"
	"webtoolkit/biz/user"
)

// Auth vérifie l'authentification de l'utilisateur et retourne un token signé.
func Auth(realm, username, password, address string) (string, error) {
	if !user.CheckPassword(realm, username, password) {
		return "", fmt.Errorf("not authorized")
	}
	t, err := secure.NewToken(username, address)
	if err != nil {
		return "", err
	}
	return t, nil
}

package metier

import (
	"fmt"
	"webtoolkit/metier/secure"
	"webtoolkit/metier/user"
)

// Auth vérifie l'authentification de l'utilisateur et retourne un token signé.
func Auth(username, password, address string) (string, error) {
	if !user.CheckPassword(username, password) {
		return "", fmt.Errorf("not authorized")
	}
	t, err := secure.NewToken(username, address)
	if err != nil {
		return "", err
	}
	return t, nil
}

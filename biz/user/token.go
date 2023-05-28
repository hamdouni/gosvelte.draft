package user

import (
	"fmt"
	"strings"
	"time"
)

const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(24) * time.Hour

// Token génère un token d'authentification, numériquement signé.
// Il est constitué de 3 parties séparées par une barre verticale :
// - un élément identifiant l'utilisateur (son nom par exemple)
// - une adresse, une localisation d'où provient le token (ip par exemple)
// - la date de création du token
func Token(user, address string) (string, error) {
	timestamp := time.Now().Format(tokenTimeLayout)
	phrase := user + "|" + address + "|" + timestamp
	token, err := secure.Encrypt(phrase)
	if err != nil {
		return "", err
	}
	return token, nil
}

// CheckToken vérifie la validité du token selon 3 critères:
// D'abord que le token a bien été signé.
// Ensuite que sa durée de validité n'est pas expirée.
// Enfin que sa provenance est identique.
func CheckToken(token, address string) (bool, error) {
	code, err := secure.Decrypt(string(token))
	if err != nil {
		return false, err
	}
	parts := strings.Split(code, "|")

	curTime := time.Now()
	loginTime, err := time.Parse(tokenTimeLayout, parts[2])
	if err != nil {
		return false, fmt.Errorf("parsing token time %s: %s", parts[2], err)
	}
	dur := curTime.Sub(loginTime)
	if dur > tokenDuration {
		return false, nil
	}

	return (address == parts[1]), nil
}

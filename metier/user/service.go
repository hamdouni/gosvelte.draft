package user

import "webtoolkit/metier/secure"

// ValidatePassword vérifie la conformité de l'identifiant et mot de passe
func CheckPassword(username, password string) bool {
	hashed := store.GetPassword(username)
	return secure.CheckPassword(password, hashed)
}

// Add enregistre un utilisateur dans le système de stockage
func Add(username, password string, role Role) error {
	hashed, err := secure.HashPassword(password)
	if err != nil {
		return err
	}
	u, err := New(username, hashed, role)
	if err != nil {
		return err
	}
	store.Add(*u)

	return nil
}

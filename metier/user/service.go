package user

import "webtoolkit/metier/secure"

// ValidatePassword vérifie la conformité de l'identifiant et mot de passe
func CheckPassword(realm, username, password string) bool {
	hashed := store.GetPassword(realm, username)
	return secure.CheckPassword(password, hashed)
}

// Add enregistre un utilisateur dans le système de stockage
func Add(realm, username, password string, role Role) error {
	hashed, err := secure.HashPassword(password)
	if err != nil {
		return err
	}
	if store.ExistUsername(realm, username) {
		return ErrUsernameUsed
	}
	u, err := New(realm, username, hashed, role)
	if err != nil {
		return err
	}
	store.Add(*u)

	return nil
}

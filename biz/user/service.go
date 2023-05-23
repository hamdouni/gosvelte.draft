package user

import "webtoolkit/biz/secure"

// ValidatePassword vérifie la conformité de l'identifiant et mot de passe
func CheckPassword(realm, username, password string) bool {
	hashed := config.store.GetPassword(realm, username)
	return secure.CheckPassword(password, hashed)
}

// Add enregistre un utilisateur dans le système de stockage
func Add(realm, username, password string, role Role) error {
	hashed, err := secure.HashPassword(password)
	if err != nil {
		return err
	}
	if config.store.ExistUsername(realm, username) {
		return ErrUsernameUsed
	}
	u, err := New(realm, username, hashed, role)
	if err != nil {
		return err
	}
	return config.store.Add(*u)
}

// List retourne la liste des utilisateurs sans le password pour un realm donné
func List(realm string) (users []User, err error) {
	users, err = config.store.ListUsers(realm)
	if err != nil {
		return users, err
	}
	// purge password
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

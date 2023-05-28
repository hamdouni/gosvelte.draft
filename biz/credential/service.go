package credential

// ValidatePassword vérifie la conformité de l'identifiant et mot de passe
func CheckPassword(realm, username, password string) bool {
	hashed := store.GetPassword(realm, username)
	return secure.CheckHash(password, hashed)
}

// Auth vérifie l'authentification de l'utilisateur et retourne un token signé.
func Auth(realm, username, password, address string) (string, error) {
	if !CheckPassword(realm, username, password) {
		return "", ErrNotAuthorized
	}
	t, err := Token(username, address)
	if err != nil {
		return "", err
	}
	return t, nil
}

// Add enregistre un utilisateur dans le système de stockage
func Add(realm, username, password string, role Role) error {
	hashed, err := secure.Hash(password)
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
	return store.Add(*u)
}

// List retourne la liste des utilisateurs sans le password pour un realm donné
func List(realm string) (users []Credential, err error) {
	users, err = store.ListUsers(realm)
	if err != nil {
		return users, err
	}
	// purge password
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

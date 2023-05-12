package user

// store est le système de stockage des utilisateurs
var store Storage

// UseStore uses specified storage for user repository
func UseStore(s Storage) {
	store = s
}

// Contrat avec le service de stockage
type Storage interface {
	GetPassword(realm, username string) (encryptedPassword string)
	ExistUsername(realm, username string) bool
	Add(user User) error
}

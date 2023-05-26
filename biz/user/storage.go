package user

type Storage interface {
	GetPassword(realm, username string) (encryptedPassword string)
	ExistUsername(realm, username string) bool
	Add(user User) error
	ListUsers(realm string) ([]User, error)
}

var store Storage

// WithRepo permet de pluger le dépot de données
func WithRepo(s Storage) {
	store = s
}

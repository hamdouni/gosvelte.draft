package user

// Contrat avec le service de stockage
type Storage interface {
	GetPassword(realm, username string) (encryptedPassword string)
	ExistUsername(realm, username string) bool
	Add(user User) error
	ListUsers(realm string) ([]User, error)
	InitSchema() error
	Close() error
}

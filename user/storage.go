package user

type Storage interface {
	GetPassword(realm, username string) (encryptedPassword string)
	ExistUsername(realm, username string) bool
	Add(user Credential) error
	ListUsers(realm string) ([]Credential, error)
}

var store Storage

// WithRepo permet de pluger le dépot de données
func WithRepo(s Storage) {
	store = s
}

// Defensive coding
type defaultStore struct{}

func init() {
	WithRepo(defaultStore{})
}

func (ds defaultStore) GetPassword(realm, username string) string {
	panic("credential.GetPassword not implemented")
}

func (ds defaultStore) ExistUsername(realm, username string) bool {
	panic("credential.ExistUsername not implemented")
}

func (ds defaultStore) Add(user Credential) error {
	panic("credential.Add not implemented")
}

func (ds defaultStore) ListUsers(realm string) ([]Credential, error) {
	panic("credential.ListUsers not implemented")
}

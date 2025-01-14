package user

type Security interface {
	Hash(password string) (hased string, err error)
	CheckHash(password, hashed string) bool
	Encrypt(phrase string) (string, error)
	Decrypt(secret string) (string, error)
}

var secure Security

func WithSecurity(s Security) {
	secure = s
}

// Defensive coding

type defaultSecurity struct{}

func init() {
	WithSecurity(defaultSecurity{})
}

func (ds defaultSecurity) Hash(password string) (hased string, err error) {
	panic("credential.Hash not implemented")
}

func (ds defaultSecurity) CheckHash(password, hashed string) bool {
	panic("credential.CheckHash not implemented")
}

func (ds defaultSecurity) Encrypt(phrase string) (string, error) {
	panic("credential.Encrypt not implemented")
}

func (ds defaultSecurity) Decrypt(secret string) (string, error) {
	panic("credential.Decrypt not implemented")
}

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

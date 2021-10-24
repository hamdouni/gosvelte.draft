package web

// Contrat avec le service de sécurite
type secure interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	AuthUser(string, string) bool
}

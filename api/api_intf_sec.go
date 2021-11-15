package api

// Contrat avec le service de sécurite
type secure interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

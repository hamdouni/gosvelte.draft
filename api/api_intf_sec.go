package api

// Contrat avec le service de sécurite
type secure interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

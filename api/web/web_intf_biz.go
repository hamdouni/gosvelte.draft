package web

// Contrat avec le service business
type business interface {
	Bonjour(string) string
	Maj(string) string
	Min(string) string
	Historic() []string
	CheckPassword(string, string) bool
}

package api

type business interface {
	Bonjour(string) string
	Maj(string) string
	Min(string) string
	Historic() []string
	Login(string, string) bool
}

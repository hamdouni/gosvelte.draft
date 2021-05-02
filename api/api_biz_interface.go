package api

type business interface {
	Bonjour(string) string
	Maj(string) string
	Min(string) string
}

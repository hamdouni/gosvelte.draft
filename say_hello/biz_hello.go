package app

// Bonjour est une fonction business qui retourne bonjour et le nom passé en paramètre.
func Bonjour(s string) string {
	res := "Bonjour " + s + " depuis le business !"
	return res
}

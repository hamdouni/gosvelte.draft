package model

// Hello retourne bonjour et le nom passé en paramètre.
func Hello(s string) string {
	res := "Bonjour " + s + " depuis le business !"
	return res
}

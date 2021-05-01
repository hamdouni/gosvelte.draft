package biz

/*
	Bonjour est une fonction business qui retourne bonjour et le nom passé en paramètre.
*/
func (b BIZ) Bonjour(s string) string {
	return "Bonjour " + s + " depuis le business !"
}

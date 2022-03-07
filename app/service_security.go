package app

func Decrypt(message string) (string, error) {
	return app.secure.Decrypt(message)
}

func Encrypt(message string) (string, error) {
	return app.secure.Encrypt(message)
}

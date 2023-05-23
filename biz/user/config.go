package user

// config stock la configuration de l'application.
var config = struct {
	store Storage
}{}

// WithRepo permet de pluger le dépot de données
func WithRepo(s Storage) {
	config.store = s
}

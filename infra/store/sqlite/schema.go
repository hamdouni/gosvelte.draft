package sqlite

// InitSchema initialize le schéma
func (s Store) InitSchema() error {
	// schéma pour la table user
	q := `CREATE TABLE IF NOT EXISTS "user" (
		"id"		INTEGER PRIMARY KEY AUTOINCREMENT,
		"username"	VARCHAR(64),
		"password"	VARCHAR(64),
		"role"		VARCHAR(64),
		"realm"		VARCHAR(64)
	)`
	_, err := s.database.Exec(q)
	return err
}

package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Store is a SQLite storage.
type Store struct {
	dbpath   string
	pragma   string
	database *sql.DB
}

// New retourne un stockage sqlite
func New(path, params string) (sqlitedb Store, err error) {

	// check file exists
	_, err = os.Stat(path)
	if err != nil {
		return sqlitedb, err
	}

	dsn := fmt.Sprintf("%s?%s", path, params)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return sqlitedb, fmt.Errorf("error opening database %s got %w", path, err)
	}
	sqlitedb.database = db
	sqlitedb.dbpath = path
	sqlitedb.pragma = params
	return sqlitedb, err
}

func (s Store) Close() error {
	return s.database.Close()
}

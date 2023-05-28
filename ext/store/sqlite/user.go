package sqlite

import "wtk/biz/credential"

func (store Store) Add(u credential.Credential) error {
	q := `INSERT INTO user (username, password, role, realm) VALUES(?,?,?,?)`
	_, err := store.database.Exec(q, u.Username, u.Password, u.Role, u.Realm)
	return err
}

func (store Store) GetPassword(realm, username string) (encryptedPassword string) {
	q := `SELECT password FROM user WHERE username=? AND realm=?`
	row := store.database.QueryRow(q, username, realm)
	_ = row.Scan(&encryptedPassword)
	return encryptedPassword
}

func (store Store) ExistUsername(realm, username string) (exists bool) {
	q := `SELECT COUNT(*) FROM user WHERE username=? AND realm=?`
	row := store.database.QueryRow(q, username, realm)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	exists = count == 1
	return exists
}

func (store Store) ListUsers(realm string) (users []credential.Credential, err error) {
	q := `SELECT username, role FROM user WHERE realm=?`
	rows, err := store.database.Query(q, realm)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var u credential.Credential
		if err := rows.Scan(&u.Username, &u.Role); err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

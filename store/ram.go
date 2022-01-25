package store

import "app"

type Store struct {
	historic []string
	users    map[string]app.User
}

func (rs *Store) Init() {
	rs.users = make(map[string]app.User)
}

func (rs *Store) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs Store) ListeHistorique() []string {
	return rs.historic
}

func (rs *Store) AddUser(username, password string, role app.RoleType) error {
	rs.users[username] = app.User{
		Password: password,
		Role:     role,
	}
	return nil
}

func (rs *Store) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

package store

import (
	"admin/model"
)

type Store struct {
	historic []string
	users    map[string]model.User
}

func New() (Store, error) {
	var rs Store
	rs.users = make(map[string]model.User)
	return rs, nil
}

func (rs *Store) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs Store) ListeHistorique() []string {
	return rs.historic
}

func (rs *Store) AddUser(username, password string, role model.Role) error {
	rs.users[username] = model.User{
		Password: password,
		Role:     role,
	}
	return nil
}

func (rs *Store) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

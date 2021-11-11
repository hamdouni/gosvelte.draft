package store

import "app/biz"

type Store struct {
	historic []string
	users    map[string]string
}

func (rs *Store) Init() {
	rs.users = make(map[string]string)
}

func (rs *Store) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs Store) ListeHistorique() []string {
	return rs.historic
}

func (rs *Store) AddUser(user biz.User) error {
	rs.users[user.Username] = user.Password
	return nil
}

func (rs *Store) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username]
}

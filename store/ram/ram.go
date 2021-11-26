package ram

import "app/biz/create"

type Store struct {
	historic []string
	users    map[string]create.User
}

func (rs *Store) Init() {
	rs.users = make(map[string]create.User)
}

func (rs *Store) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs Store) ListeHistorique() []string {
	return rs.historic
}

func (rs *Store) AddUser(username, password string, role create.RoleType) error {
	rs.users[username] = create.User{
		Password: password,
		Role:     role,
	}
	return nil
}

func (rs *Store) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

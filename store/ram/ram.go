package ram

import (
	"admin/model"
)

type RAM struct {
	historic []string
	users    map[string]model.User
}

func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]model.User)
	return rs, nil
}

func (rs *RAM) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs RAM) ListeHistorique() []string {
	return rs.historic
}

func (rs *RAM) AddUser(username, password string, role model.Role) error {
	rs.users[username] = model.User{
		Password: password,
		Role:     role,
	}
	return nil
}

func (rs *RAM) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

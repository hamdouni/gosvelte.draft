package ram

import "webtoolkit/metier/user"

func (rs *RAM) Add(u user.User) error {
	rs.users[u.Username] = user.User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
	return nil
}

func (rs *RAM) GetPassword(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

func (rs *RAM) ExistUsername(username string) bool {
	_, exists := rs.users[username]
	return exists
}

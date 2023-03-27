package ram

import "admin/model"

func (rs *RAM) AddUser(user model.User) error {
	rs.users[user.Username] = model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	return nil
}

func (rs *RAM) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username].Password
}

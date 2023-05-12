package ram

import "webtoolkit/metier/user"

func key(realm, username string) string {
	return realm + ":" + username
}
func (rs *RAM) Add(u user.User) error {
	k := key(u.Realm, u.Username)
	rs.users[k] = user.User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
		Realm:    u.Realm,
	}
	return nil
}

func (rs *RAM) GetPassword(realm, username string) (encryptedPassword string) {
	k := key(realm, username)
	return rs.users[k].Password
}

func (rs *RAM) ExistUsername(realm, username string) bool {
	k := key(realm, username)
	_, exists := rs.users[k]
	return exists
}

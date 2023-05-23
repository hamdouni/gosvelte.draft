package ram

import "wtk/biz/user"

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

func (rs *RAM) ListUsers(realm string) (users []user.User, err error) {
	for _, u := range rs.users {
		if u.Realm == realm {
			users = append(users, u)
		}
	}
	return users, nil
}

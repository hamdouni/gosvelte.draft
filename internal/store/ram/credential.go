package ram

import "wtk/user"

func key(realm, username string) string {
	return realm + ":" + username
}

func (rs *RAM) Add(u user.Credential) error {
	k := key(u.Realm, u.Username)
	rs.users[k] = user.Credential{
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

func (rs *RAM) ListUsers(realm string) (users []user.Credential, err error) {
	for _, u := range rs.users {
		if u.Realm == realm {
			users = append(users, u)
		}
	}
	return users, nil
}

// ram does not use schema
func (rs *RAM) InitSchema() error {
	return nil
}

// ram does not close
func (rs *RAM) Close() error {
	return nil
}

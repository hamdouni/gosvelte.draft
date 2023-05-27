package ram

import "wtk/biz/user"

type RAM struct {
	users map[string]user.User
}

// New retourne une RAM pouvant stocker des données
func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]user.User)
	return rs, nil
}

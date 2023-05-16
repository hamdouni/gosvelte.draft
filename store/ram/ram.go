package ram

import "webtoolkit/biz/user"

type RAM struct {
	historic []string
	users    map[string]user.User
}

// New retourne une RAM pouvant stocker des données
func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]user.User)
	return rs, nil
}

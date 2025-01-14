package ram

import "wtk/user"

type RAM struct {
	users map[string]user.Credential
}

// New retourne une RAM pouvant stocker des données
func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]user.Credential)
	return rs, nil
}

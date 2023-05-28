package ram

import "wtk/biz/credential"

type RAM struct {
	users map[string]credential.Credential
}

// New retourne une RAM pouvant stocker des données
func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]credential.Credential)
	return rs, nil
}

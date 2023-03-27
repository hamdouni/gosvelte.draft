package ram

import (
	"admin/model"
)

type RAM struct {
	historic []string
	users    map[string]model.User
}

// New retourne une RAM pouvant stocker des données
func New() (RAM, error) {
	var rs RAM
	rs.users = make(map[string]model.User)
	return rs, nil
}

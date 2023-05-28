package biz

import (
	"wtk/biz/user"
)

func Initialize(sto Storage, sec Security) error {

	user.WithSecurity(sec)

	user.WithRepo(sto)

	return nil
}

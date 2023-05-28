package biz

import (
	"wtk/biz/credential"
)

func Initialize(sto Storage, sec Security) error {
	credential.WithSecurity(sec)
	credential.WithRepo(sto)
	return nil
}

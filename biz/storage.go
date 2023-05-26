package biz

import "wtk/biz/user"

type Storage interface {
	user.Storage
	InitSchema() error
	Close() error
}

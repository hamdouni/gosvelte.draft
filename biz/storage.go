package biz

import "wtk/user"

type Storage interface {
	user.Storage
	InitSchema() error
	Close() error
}

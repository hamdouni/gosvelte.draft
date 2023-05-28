package biz

import "wtk/biz/credential"

type Storage interface {
	credential.Storage
	InitSchema() error
	Close() error
}

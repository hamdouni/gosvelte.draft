package web_test

import (
	"app/pkg/ram"
	"app/pkg/web"
)

type fakeSecType struct{}

func (s *fakeSecType) AuthUser(u, p string) bool        { return u == p }
func (s *fakeSecType) Encrypt(p string) (string, error) { return p, nil }
func (s *fakeSecType) Decrypt(p string) (string, error) { return p, nil }

func init() {
	var fakeSec fakeSecType
	var fakeStore ram.RamStore
	fakeStore.Init()
	var fakeWeb web.WEB
	fakeWeb.Init(&fakeSec, &fakeStore, ".")
}

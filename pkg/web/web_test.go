package web_test

import (
	"app/pkg/ram"
	"app/pkg/web"
)

const attendu = "Fake Biz"

type fakeBizType struct{}

func (fakeBizType) Bonjour(string) string { return attendu }
func (fakeBizType) Maj(string) string     { return attendu }
func (fakeBizType) Min(string) string     { return attendu }
func (fakeBizType) Historic() []string    { return []string{attendu} }

type fakeSecType struct{}

func (s *fakeSecType) AuthUser(u, p string) bool        { return u == p }
func (s *fakeSecType) Encrypt(p string) (string, error) { return p, nil }
func (s *fakeSecType) Decrypt(p string) (string, error) { return p, nil }

func init() {
	var fakeBiz fakeBizType
	var fakeSec fakeSecType
	var fakeStore ram.RamStore
	fakeStore.Init()
	var fakeWeb web.WEB
	fakeWeb.Init(&fakeBiz, &fakeSec, &fakeStore, ".")
}

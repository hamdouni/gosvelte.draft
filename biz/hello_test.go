package biz_test

import (
	"testing"
	"wtk/biz"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le métier !"
	got := biz.Hello(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

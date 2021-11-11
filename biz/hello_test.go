package biz_test

import (
	"app/biz"
	"testing"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le business !"
	got := biz.Hello(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

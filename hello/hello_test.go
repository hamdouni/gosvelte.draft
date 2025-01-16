package hello_test

import (
	"testing"
	"wtk/hello"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le métier !"
	got := hello.Hello(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

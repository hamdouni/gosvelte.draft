package greet_test

import (
	"app/biz/greet"
	"testing"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le business !"
	got := greet.Hello(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

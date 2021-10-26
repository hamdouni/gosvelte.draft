package app_test

import (
	app "app/say_hello"
	"testing"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le business !"
	got := app.Bonjour(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

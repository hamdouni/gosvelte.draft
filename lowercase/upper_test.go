package app_test

import (
	app "app/uppercase"
	"testing"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := app.Maj(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

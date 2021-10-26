package app_test

import (
	app "app/lowercase"
	"testing"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := app.Min(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

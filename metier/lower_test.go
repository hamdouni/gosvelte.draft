package metier_test

import (
	"testing"
	"webtoolkit/metier"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := metier.Lower(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

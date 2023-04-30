package metier_test

import (
	"testing"
	"webtoolkit/metier"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := metier.Upper(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

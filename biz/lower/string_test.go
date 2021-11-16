package lower_test

import (
	"app/biz/lower"
	"testing"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := lower.String(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

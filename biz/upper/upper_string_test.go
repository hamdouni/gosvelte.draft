package upper_test

import (
	"app/biz/upper"
	"testing"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := upper.String(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

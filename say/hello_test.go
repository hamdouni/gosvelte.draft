package say_test

import (
	"app/say"
	"testing"
)

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le business !"
	got := say.Hello(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}

package strings_test

import (
	"testing"
)

func TestStrings(t *testing.T) {
	t.Parallel()
	s := "привет"
	l := len(s)

	t.Run("string is utf8", func(t *testing.T) {
		t.Parallel()
		got, want := s[0:3], "п\xd1" // slice over bytes, not runes
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	if l != 12 {
		t.Error("unxpected len")
	}
}

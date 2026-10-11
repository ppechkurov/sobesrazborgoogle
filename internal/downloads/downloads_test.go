package downloads_test

import (
	"testing"

	"github.com/ppechkurov/sobesrazborgoogle/internal/downloads"
)

func TestRun(t *testing.T) {
	err := downloads.Run(t.Context(), []string{
		"https://google.com",
		"https://github.com",
		"https://git.slonverse.xyz",
	})
	if err != nil {
		t.Fatal(err)
	}
}

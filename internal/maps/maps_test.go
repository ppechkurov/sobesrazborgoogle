package maps_test

import (
	"testing"

	"github.com/ppechkurov/sobesrazborgoogle/internal/maps"
)

func TestMaps(t *testing.T) {
	t.Parallel()
	a, b := maps.Run()

	if a != 20000 {
		t.Errorf("a: want %d, got %d", 20000, a)
	}

	if b != 10000 {
		t.Errorf("b: want %d, got %d", 10000, b)
	}
}

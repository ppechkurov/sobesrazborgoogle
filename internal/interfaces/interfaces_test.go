package interfaces_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/ppechkurov/sobesrazborgoogle/internal/interfaces"
)

func TestNilIfaceValueIsNotNil(t *testing.T) {
	isNil := interfaces.Run()
	if isNil {
		t.Errorf("unexpected nil interface")
	}
}

func TestNilIfaceAndNilValueIsNil(t *testing.T) {
	isNil := interfaces.RunNil()
	if !isNil {
		t.Errorf("unexpected NON nil interface")
	}

	var x any
	v, ok := x.(string) // safe
	if !ok {
		return
	}

	v = x.(string) // panic, if not a string
	switch v := x.(type) {
	case int, int64:
	case fmt.Stringer:
		_ = v
	}

	_ = v
}

func TestAlignment(t *testing.T) {
	t.Parallel()
	notAligned := unsafe.Sizeof(interfaces.Bad{})
	t.Logf("notAligned: %v\n", notAligned)

	aligned := unsafe.Sizeof(interfaces.Good{})
	t.Logf("aligned: %v\n", aligned)
}

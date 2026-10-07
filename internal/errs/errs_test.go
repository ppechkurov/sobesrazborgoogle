package errs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ppechkurov/sobesrazborgoogle/internal/errs"
)

func TestErrors(t *testing.T) {
	err := errs.MyError{Msg: "MyError msg"}
	t.Log(err)

	errWrapped := fmt.Errorf("wrapped: %w", err)
	t.Log(errWrapped)

	var asErr errs.MyError
	extracted := errors.As(errWrapped, &asErr)
	t.Logf("extracted: %v\n", extracted)

	if asType, ok := errors.AsType[errs.MyError](errWrapped); ok {
		t.Logf("as type: %v", asType)
	}

	t.Log(errors.Join(errors.New("joined error"), errWrapped, errWrapped))
	t.Log(fmt.Errorf("%w; %w", errWrapped, errWrapped))
}

func TestRecover(t *testing.T) {
	t.Log(errs.Recover())
}

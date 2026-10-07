package errs

import (
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func Recover() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v\n%s", r, debug.Stack())
		}
	}()

	if true {
		panic(nil)
	}

	return err
}

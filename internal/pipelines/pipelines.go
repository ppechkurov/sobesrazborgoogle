package pipelines

import (
	"context"
	"fmt"
)

func Pipe() {}

func Generate[T comparable](ctx context.Context, nums ...T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for _, i := range nums {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func Square[T int](ctx context.Context, items <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for i := range items {
			select {
			case out <- i * i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func Print[T comparable](ctx context.Context, in <-chan T) {
	for {
		select {
		case <-ctx.Done():
			return
		case n, more := <-in:
			if !more {
				return
			}
			fmt.Println(n)
		}
	}
}

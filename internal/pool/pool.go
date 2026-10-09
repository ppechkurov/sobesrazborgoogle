package pool

import (
	"context"
	"fmt"
	"sync"
)

type Task = func(ctx context.Context, val string) (string, error)

type Result struct {
	Val string
	Err error
}

func Run(ctx context.Context, workers uint, tasks <-chan Task) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	for i := range workers {
		wg.Go(func() {
			fmt.Println("worker", i, "started")
			defer fmt.Println("worker", i, "exited")

			for {
				select {
				case f, more := <-tasks:
					if !more {
						return
					}

					val, err := f(ctx, "task")

					select {
					case out <- Result{val, err}:
					case <-ctx.Done():
						return
					}

				case <-ctx.Done():
					return
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

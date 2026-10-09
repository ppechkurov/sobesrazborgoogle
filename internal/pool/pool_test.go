package pool_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ppechkurov/sobesrazborgoogle/internal/pool"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	tasks := make(chan pool.Task)

	for i := range 100 {
		wg.Go(func() {
			tasks <- func(ctx context.Context, val string) (string, error) {
				select {
				case <-time.After(time.Millisecond * 1000):
				case <-ctx.Done():
					return "", ctx.Err()
				}

				return fmt.Sprint("task ", i, " completed"), nil
			}
		})
	}
	go func() {
		wg.Wait()
		close(tasks)
	}()

	out := pool.Run(t.Context(), 4, tasks)

	for res := range out {
		if res.Err != nil {
			fmt.Println(res.Err)
		} else {
			fmt.Println(res.Val)
		}
	}
}

package channels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx context.Context) error {
	var wg sync.WaitGroup
	first := make(chan string, 1)
	second := make(chan string, 1)

	wg.Go(func() {
		defer close(first)

		for i := range 10 {
			select {
			case <-ctx.Done():
				fmt.Println("first cancelled...")
				return
			case first <- fmt.Sprintf("%d: first", i):
				time.Sleep(time.Second)
			}
		}
	})

	wg.Go(func() {
		defer close(second)

		for i := range 10 {
			select {
			case <-ctx.Done():
				fmt.Println("second cancelled...")
				return
			case second <- fmt.Sprintf("%d: second", i):
				time.Sleep(time.Second)
			}
		}
	})

	for _, ch := range []chan string{first, second} {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("ctx cancelled")
					return
				case msg, more := <-ch:
					if !more {
						return
					}
					fmt.Println(msg)
				}
			}
		})
	}

	wg.Wait()

	return ctx.Err()
}

func Merge[T comparable](ctx context.Context, chans ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	for _, ch := range chans {
		wg.Go(func() {
			for v := range ch {
				select {
				case out <- v:
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

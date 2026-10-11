package downloads

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Run(ctx context.Context, urls []string) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	sem := make(chan struct{}, 10)

	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error

lp:
	for i, url := range urls {
		select {
		case sem <- struct{}{}:
			wg.Go(func() {
				fmt.Println("started", i)
				defer fmt.Println("done", i)
				defer func() { <-sem }()

				err := download(ctx, url)
				if err != nil {
					once.Do(func() {
						firstErr = err
						cancel(err)
					})
				}
			})

		case <-ctx.Done():
			break lp
		}
	}

	wg.Wait()
	if firstErr != nil {
		return firstErr
	}

	return context.Cause(ctx)
}

func download(ctx context.Context, url string) error {
	client := http.Client{
		Transport: &http.Transport{},
		Timeout:   time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("doing request: %w", err)
	}
	defer res.Body.Close()

	fmt.Println(url, res.StatusCode)

	return nil
}

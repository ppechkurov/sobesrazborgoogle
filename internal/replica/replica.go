package replica

import (
	"context"
	"errors"
)

func Query(
	ctx context.Context,
	urls []string,
	call func(ctx context.Context, url, query string) (string, error),
) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	type res struct {
		msg string
		err error
	}

	ch := make(chan res, len(urls))
	sem := make(chan struct{}, 10)

	go func() {
		for _, url := range urls {
			select {
			case <-ctx.Done():
				return

			case sem <- struct{}{}:
				go func() {
					defer func() { <-sem }()
					msg, err := call(ctx, url, "some query?")
					ch <- res{msg: msg, err: err}
				}()
			}
		}
	}()

	errs := make([]error, 0, len(urls))
	for range len(urls) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()

		case r := <-ch:
			if r.err == nil {
				return r.msg, nil
			}
			errs = append(errs, r.err)
		}
	}

	return "", errors.Join(append([]error{errors.New("all replicas failed")}, errs...)...)
}

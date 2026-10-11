package pipelines

import (
	"context"
	"database/sql"
	"fmt"
)

func GetRows(ctx context.Context, db *sql.DB) (*sql.Rows, error) {
	return db.QueryContext(ctx,
		`
WITH RECURSIVE
  seq (n) AS (
    SELECT
      1
    UNION ALL
    SELECT
      n + 1
    FROM
      seq
    WHERE
      n < 1000
  )
SELECT
  n AS id,
  'user-' || n AS name,
  'https://example.com/users/' || n AS url
FROM
  seq
`,
	)
}

type Row struct {
	ID        int
	Name, url string
}

func Read(ctx context.Context, rows *sql.Rows) <-chan Row {
	out := make(chan Row)
	go func() {
		defer close(out)

		var res Row
		for rows.Next() {
			if err := rows.Scan(&res.ID, &res.Name, &res.url); err != nil {
				panic(fmt.Errorf("scanning row: %w", err))
			}

			select {
			case out <- res:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func Filter[T any](ctx context.Context, in <-chan T, fn func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)

		for v := range in {
			if !fn(v) {
				continue
			}

			select {
			case out <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

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

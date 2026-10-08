package replica_test

import (
	"context"
	"testing"
	"time"

	"github.com/ppechkurov/sobesrazborgoogle/internal/replica"
)

func TestQuery(t *testing.T) {
	call := func(ctx context.Context, url, query string) (string, error) {
		if url == "first" {
			time.Sleep(time.Millisecond * 1000)
			return "first", nil
		}
		return url, nil
	}

	for range 1000 {
		res, err := replica.Query(t.Context(), []string{"first", "second", "third"}, call)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(res)
		}
	}
}

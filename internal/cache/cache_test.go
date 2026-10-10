package cache_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ppechkurov/sobesrazborgoogle/internal/cache"
)

func TestGet(t *testing.T) {
	t.Parallel()
	c := cache.New[string, string]()

	t.Run("returns value by key", func(t *testing.T) {
		t.Parallel()
		c.Put("test_key", "test_val")

		got, ok := c.Get("test_key")
		if !ok {
			t.Fatal("value is expected")
		}

		want := "test_val"
		if want != got {
			t.Fatalf("want %v, got %v", want, got)
		}
	})

	t.Run("returns (zero,false) for non existing", func(t *testing.T) {
		t.Parallel()

		got, ok := c.Get("non existing key")
		if ok {
			t.Error("value should not exist")
		}
		if got != "" {
			t.Errorf("default value expected. got %v", got)
		}
	})

	t.Run("returns (zero,false) when key is expired", func(t *testing.T) {
		c = cache.New(
			cache.WithTTL[string, string](1),
		)
		t.Parallel()
		c.Put("test_key", "test_val")
		time.Sleep(time.Millisecond * 5)
		_, ok := c.Get("test_key")
		if ok {
			t.Fatal("expected expired cache entry")
		}
	})
}

func TestDelete(t *testing.T) {
	c := cache.New[string, string]()
	c.Put("test_key", "test_val")
	if _, ok := c.Get("test_key"); !ok {
		t.Fatal("value expected")
	}

	c.Delete("test_key")
	if val, ok := c.Get("test_key"); ok {
		t.Fatalf("value is on expected. got %v", val)
	}
}

func TestConcurrency(t *testing.T) {
	c := cache.New[string, int]()
	var wg sync.WaitGroup
	wg.Go(func() { c.Get("test_key") })
	wg.Go(func() { c.Put("test_key", 1) })
	wg.Go(func() { c.Delete("test_key") })
	wg.Wait()
}

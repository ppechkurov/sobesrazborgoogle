package maps

import (
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[key] += 1
}

func Run() (a, b int) {
	var wg sync.WaitGroup
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	wg.Go(func() {
		for range 10000 {
			c.Inc("a")
		}
	})

	wg.Go(func() {
		for range 10000 {
			c.Inc("a")
		}
	})

	wg.Go(func() {
		for range 10000 {
			c.Inc("b")
		}
	})

	wg.Wait()

	return c.counters["a"], c.counters["b"]
}

package cache

import (
	"sync"
	"time"
)

type option[K comparable, V any] func(*Cache[K, V])

type entry[V any] struct {
	val   V
	until int64
}

type Cache[K comparable, V any] struct {
	m   map[K]entry[V]
	mu  sync.RWMutex
	ttl int64
}

func New[K comparable, V any](opts ...option[K, V]) *Cache[K, V] {
	c := Cache[K, V]{
		m:   make(map[K]entry[V]),
		ttl: 1000,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

func WithTTL[K comparable, V any](ttl int64) option[K, V] {
	return func(c *Cache[K, V]) { c.ttl = ttl }
}

func (c *Cache[K, V]) Get(key K) (val V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ent, ok := c.m[key]
	now := time.Now().UnixMilli()
	until := ent.until
	if ok && (until >= now) {
		return ent.val, true
	}
	var zero V
	return zero, false
}

func (c *Cache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = entry[V]{value, time.Now().UnixMilli() + c.ttl}
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

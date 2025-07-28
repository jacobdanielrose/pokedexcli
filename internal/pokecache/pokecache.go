package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux   *sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	element, ok := c.cache[key]
	return element.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

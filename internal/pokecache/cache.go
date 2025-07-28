package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	v        map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	cache.interval = interval
	cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.v[key] = cacheEntry{
		time.Now(),
		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	element, ok := c.v[key]
	if ok {
		return element.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		for key, entry := range c.v {
			age := time.Since(entry.createdAt)
			if age > c.interval {
				delete(c.v, key)
			}
		}
	}
}

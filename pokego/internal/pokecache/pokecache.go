package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	if entry, ok := c.cache[key]; ok {
		val = entry.val
	}
	c.mu.Unlock()
	return val, ok
}

func (c *Cache) reapLoop() {
}

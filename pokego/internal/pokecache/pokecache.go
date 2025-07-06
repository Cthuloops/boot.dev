package pokecache

import (
	"log"
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

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{}
	newCache.cache = make(map[string]cacheEntry)
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	log.Printf("Add %s to cache", key)
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	log.Printf("Attempt to get cache entry: %s\n", key)
	c.mu.Lock()
	if entry, ok := c.cache[key]; ok {
		log.Printf("Cache entry found\n")
		val = entry.val
	}
	c.mu.Unlock()
	return val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		for key := range c.cache {
			timeDifference := time.Since(c.cache[key].createdAt)
			if timeDifference > interval {
				log.Printf("It's been %v since %s was created, attempting to delete", timeDifference, key)
				c.mu.Lock()
				delete(c.cache, key)
				c.mu.Unlock()
			}
		}
	}
}

func (c *Cache) PrintKeys() {
	c.mu.Lock()
	for key := range c.cache {
		log.Printf("Key: %s; Created at: %v.2", key, c.cache[key].createdAt)
	}
	c.mu.Unlock()
}

package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(URL string, val []byte) {
	c.cache[URL] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(URL string) ([]byte, bool) {
	result, ok := c.cache[URL]
	return result.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		dTime := time.Now().Add(-interval)
		for k, entry := range c.cache {
			if entry.createdAt.Before(dTime) {
				delete(c.cache, k)
			}
		}
	}
}

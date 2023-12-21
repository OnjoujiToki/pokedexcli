package pokecache

import "time"

// Path: internal/pokecache/pokecache.go
type Cache struct {
	cache map[string]cacheEntry
}
type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// this part will run every interval
		c.Reap(interval)
	}
}
func (c *Cache) Reap(interval time.Duration) {
	for k, v := range c.cache {
		if time.Since(v.createdAt) > interval {
			delete(c.cache, k)
		}
	}
}

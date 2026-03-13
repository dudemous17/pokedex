package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents the data stored in the cache
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache holds the map of cache entries and a mutex for concurrent access.
type Cache struct {
	mu    *sync.Mutex
	cache map[string]cacheEntry
}

// NewCache initializes a new Cache instance.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.cache[key]
	if !found {
		return nil, false
	}
	return entry.val, found
}

// reapLoop removes entries older than the interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for key, entry := range c.cache {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}

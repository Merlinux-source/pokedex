package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (e *cacheEntry) CreatedAt() time.Time { return e.createdAt }
func (e *cacheEntry) Value() []byte        { return e.value }

type Cache struct {
	mu       sync.RWMutex
	duration time.Duration
	entries  map[string]cacheEntry
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{time.Now(), value}
}

func (c *Cache) Get(key string) (value []byte, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.value, ok
}

func (c *Cache) readLoop(interval time.Duration) {
	// a simple cache invalidator.
	for ; ; time.Sleep(interval) {
		c.mu.RLock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(time.Now().Add(-c.duration)) {
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
		c.mu.RUnlock()
	}
}

func NewCache(duration time.Duration) (returnCache *Cache) {
	returnCache = &Cache{}
	returnCache.duration = duration
	returnCache.entries = make(map[string]cacheEntry)

	go returnCache.readLoop(duration) // initiate the cache invalidation in a new thread.

	return returnCache
}

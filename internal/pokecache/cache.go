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
	mu      sync.Mutex
	entries map[string]cacheEntry
}

var _cache *Cache

func NewCache(interval time.Duration) *Cache {
	if _cache == nil {
		_cache = &Cache{}
		_cache.reapLoop(interval)
	}
	return _cache
}

func (c *Cache) Add(key string, newVal []byte) {
	c.mu.Lock()
	if c.entries == nil {
		c.entries = make(map[string]cacheEntry)
	}
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: newVal}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cachedVlue, ok := c.entries[key]
	c.mu.Unlock()
	if ok {
		return cachedVlue.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for {
			// Block and wait for interval to pass
			<-ticker.C
			// Go over cache and delete old entries
			var keysToDelete []string
			timeLimit := time.Now()
			c.mu.Lock()
			for key, entry := range c.entries {
				// Check if entry is too old

				if entry.createdAt.Add(interval).Before(timeLimit) {
					keysToDelete = append(keysToDelete, key)
				}
			}

			//fmt.Println("Deleting cache...")
			for _, key := range keysToDelete {
				//fmt.Println(key)
				delete(c.entries, key)
			}
			c.mu.Unlock()
		}
	}()
}

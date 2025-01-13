package cache

import (
	"sync"
	"time"
)

type QueryCache struct {
	cache    map[string]*CacheEntry
	mutex    sync.RWMutex
	maxSize  int
	lifetime time.Duration
}

type CacheEntry struct {
	Value     interface{}
	CreatedAt time.Time
}

func NewQueryCache(maxSize int, lifetime time.Duration) *QueryCache {
	return &QueryCache{
		cache:    make(map[string]*CacheEntry),
		maxSize:  maxSize,
		lifetime: lifetime,
	}
}

func (c *QueryCache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if entry, exists := c.cache[key]; exists {
		if time.Since(entry.CreatedAt) < c.lifetime {
			return entry.Value
		}
		delete(c.cache, key)
	}
	return nil
}

func (c *QueryCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.cache) >= c.maxSize {
		c.evictOldest()
	}

	c.cache[key] = &CacheEntry{
		Value:     value,
		CreatedAt: time.Now(),
	}
}

func (c *QueryCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, entry := range c.cache {
		if oldestKey == "" || entry.CreatedAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.CreatedAt
		}
	}

	delete(c.cache, oldestKey)
}

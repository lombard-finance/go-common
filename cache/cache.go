package cache

import (
	"sync"
	"time"
)

const CACHE_TTL = time.Hour * 2

type Cache struct {
	ttl time.Duration
	// use sync.Map for concurrent access
	values sync.Map
}

type cacheEntry struct {
	value     any
	expiresAt time.Time
}

func NewCache(ttl *time.Duration) *Cache {
	defaultTTL := CACHE_TTL
	if ttl != nil {
		defaultTTL = *ttl
	}

	return &Cache{ttl: defaultTTL}
}

func (c *Cache) SetCache(key string, value any) {
	c.values.Store(key, cacheEntry{
		value:     value,
		expiresAt: time.Now().Add(c.ttl),
	})
}

func (c *Cache) GetCache(key string) any {
	entry, ok := c.values.Load(key)
	if !ok {
		return nil
	}

	// type assertion and ttl check
	ce := entry.(cacheEntry)
	if ce.expiresAt.Unix() < time.Now().Unix() {
		return nil
	}

	return ce.value
}

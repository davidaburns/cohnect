package cache

import (
	"sync"
	"time"
)

type cacheValue struct {
	value any
	expires *time.Time
}

type InMemoryCache struct {
	data map[string]cacheValue
	mtx sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]cacheValue),
	}
}

func (c *InMemoryCache) Set(key string, value any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.data[key] = cacheValue{
		value: value,
		expires: nil,
	}
}

func (c *InMemoryCache) SetTTL(key string, value any, ttl time.Duration) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	expires := time.Now().Add(ttl);
	c.data[key] = cacheValue{
		value: value,
		expires: &expires,
	}
}

func (c *InMemoryCache) Get(key string) (any, bool) {
	c.mtx.RLock()
	defer c.mtx.Unlock()

	value, ok := c.data[key];
	if !ok {
		return nil, false
	}

	if value.expires != nil && value.expires.Before(time.Now()) {
		delete(c.data, key)
		return nil, false
	}

	return value.value, true
}

func (c *InMemoryCache) Delete(key string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.data, key)
}

func (c* InMemoryCache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.data = make(map[string]cacheValue)
}
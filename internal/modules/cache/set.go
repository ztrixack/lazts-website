package cache

import "time"

func (c *module) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(c.ttl).UnixNano(),
	}
	c.mu.Unlock()
}

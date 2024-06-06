package cache

import (
	"sync"
	"time"
)

type Moduler interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type module struct {
	config *config
	data   map[string]CacheItem
	mu     sync.RWMutex
	ttl    time.Duration
}

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

var _ Moduler = (*module)(nil)

func New() *module {
	c := parseConfig()
	return &module{
		config: c,
		data:   make(map[string]CacheItem),
		ttl:    time.Duration(c.TTLSeconds) * time.Second,
	}
}

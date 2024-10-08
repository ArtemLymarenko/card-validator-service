package cache

import (
	lrucache "card-validator-service/pkg/lru_cache"
	"sync"
)

// Cache supports any implementation (LRU or LFU)
type Cache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, ok bool)
}

type CardIssuerCache struct {
	cache Cache[string, interface{}]
	mx    sync.RWMutex
}

func NewCardIssuerCache(cap int) *CardIssuerCache {
	return &CardIssuerCache{
		cache: lrucache.NewLRU[string, interface{}](cap),
		mx:    sync.RWMutex{},
	}
}

func (c *CardIssuerCache) Add(iin string, issuer interface{}) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cache.Put(iin, issuer)
}

func (c *CardIssuerCache) Get(iin string) interface{} {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.cache.Get(iin)
	if !ok {
		return nil
	}

	return value
}

package cache

import (
	"card-validator-service/internal/domain/model"
	lrucache "card-validator-service/pkg/lru_cache"
	"sync"
)

type Cache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, ok bool)
}

type CardIssuerCache struct {
	cache Cache[model.IIN, *model.CardIssuer]
	mx    sync.RWMutex
}

func NewCardIssuerCache(cap int) *CardIssuerCache {
	return &CardIssuerCache{
		cache: lrucache.NewLRU[model.IIN, *model.CardIssuer](cap),
		mx:    sync.RWMutex{},
	}
}

func (c *CardIssuerCache) Add(iin model.IIN, issuer *model.CardIssuer) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cache.Put(iin, issuer)
}

func (c *CardIssuerCache) Get(iin model.IIN) *model.CardIssuer {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.cache.Get(iin)
	if !ok {
		return nil
	}

	return value
}

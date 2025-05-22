package geecache

import (
	"bee-web/geecache/lru"
	"sync"
)

type cache struct {
	mu        sync.Mutex
	lru       *lru.Cache
	cacheByte int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheByte, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if ele, ok := c.lru.Get(key); ok {
		return ele.(ByteView), ok
	}

	return
}

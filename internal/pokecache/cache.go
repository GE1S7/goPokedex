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
	cacheMap map[string]cacheEntry
	sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	// create new cache
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)

	return &cache

}

func (c *Cache) Add(key string, val []byte) {
	// add new entry to the chache
	c.Lock()
	defer c.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// retrieve cache entry
	c.Lock()
	defer c.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return []byte{}, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	// delete cache entry if exists longer than the interval set
	for i := 0; ; i++ {
		time.Sleep(interval)
		c.Lock()
		for k, v := range c.cacheMap {
			t := time.Now()
			elapsed := t.Sub(v.createdAt)
			if elapsed > interval {
				delete(c.cacheMap, k)
			}
		}
		c.Unlock()
	}
}

package cache

import (
	"time"

	"github.com/FishGoddess/cachego"
)

type Cache struct {
	cache cachego.Cache
}

var cache *Cache

func init() {
	cache = NewCache()
}

func GetCache() *Cache {
	return cache
}

func NewCache() *Cache {
	cache := cachego.NewCache()
	return &Cache{
		cache: cache,
	}
}

func (c *Cache) Set(key string, value interface{}, ttl int64) {
	c.cache.Set(key, value, time.Duration(ttl*int64(time.Second)))
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Close() {
	c.cache.GC()
}

func (c *Cache) Del(key string) {
	c.cache.Remove(key)
}

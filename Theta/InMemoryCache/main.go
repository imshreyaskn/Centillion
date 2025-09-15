package main

import (
	"fmt"
	"time"
)

type CacheEntry struct {
	value any
	ttl   int64
}

type Cache struct {
	data map[string]CacheEntry
}

func newCache() *Cache {
	return &Cache{
		data: make(map[string]CacheEntry),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {

	expiration := time.Now().Add(ttl).UnixNano()

	c.data[key] = CacheEntry{
		value: value,
		ttl:   expiration,
	}

}

func (c *Cache) Get(key string) (any, bool) {
	CacheEntry, found := c.data[key]
	if !found {
		return nil, false
	}

	return CacheEntry.value, true
}

func main() {

	Cache := newCache()

	Cache.Set("User_1", "Alice", 5*time.Second)
	Cache.Set("api_key", "1ABC44", 5*time.Minute)

	val, ok := Cache.Get("api_key")

	if !ok {
		fmt.Println("Key Not Found")
	}

	fmt.Println(val)

}

package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[string]Item
	sync.RWMutex
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

func New() *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	expiration := time.Now().Add(ttl).UnixNano()

	c.items[key] = Item{
		Value:      value,
		Created:    time.Now(),
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, errors.New("cacheItem with such key doesn't exist")
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, errors.New("cacheItem is expired")
		}

	}

	return item.Value, nil
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

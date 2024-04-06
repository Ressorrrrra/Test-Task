package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items                 map[int]Item
	cleanUpInterval       time.Duration
	defaultExpirationTime time.Duration
	sync.RWMutex
}

type Item struct {
	Value          interface{}
	CreatedAt      time.Time
	ExpirationTime int64
}

func (c *Cache) Add(id int, item interface{}, expTime time.Duration) {
	if expTime == 0 {
		expTime = c.defaultExpirationTime
	}

	expiration := time.Now().Add(expTime).UnixNano()

	c.Lock()

	defer c.Unlock()

	c.items[id] = Item{
		Value:          item,
		ExpirationTime: expiration,
		CreatedAt:      time.Now(),
	}
}

func New(_defaultExpirationTime, _cleanUpInterval time.Duration) *Cache {

	_items := make(map[int]Item)

	cache := Cache{
		items:                 _items,
		defaultExpirationTime: _defaultExpirationTime,
		cleanUpInterval:       _cleanUpInterval,
	}

	if _cleanUpInterval > 0 {
		cache.StartGC()
	}

	return &cache
}

func (c *Cache) Get(key int) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	if item.ExpirationTime > 0 {
		if time.Now().UnixNano() > item.CreatedAt.UnixNano()+item.ExpirationTime {
			return nil, false
		}
	}

	return item.Value, true
}

func (c *Cache) Delete(key int) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {

	for {

		<-time.After(c.cleanUpInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

func (c *Cache) expiredKeys() (keys []int) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.ExpirationTime && i.ExpirationTime > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []int) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}

package cache

import (
	"image-previewer/pkg/cache/lru"
	"sync"
)

type lruCache struct {
	capacity int
	queue    lru.List
	items    map[string]*lru.Item
	mu       *sync.Mutex
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    lru.NewList(),
		items:    make(map[string]*lru.Item, capacity),
		mu:       new(sync.Mutex),
	}
}

func (c *lruCache) Set(key string, value interface{}) bool {
	c.mu.Lock()

	defer c.mu.Unlock()

	item := &cacheItem{key: key, value: value}
	if i, ok := c.items[key]; ok {
		i.Value = item
		c.queue.MoveToFront(i)

		return true
	}

	if c.queue.Len() == c.capacity {
		lastItem, ok := c.queue.Back().Value.(*cacheItem)
		if !ok {
			return false
		}

		c.queue.Remove(c.queue.Back())
		delete(c.items, lastItem.key)
	}

	pushed := c.queue.PushFront(item)
	c.items[key] = pushed

	return false
}

func (c *lruCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		cachedItem, ok := item.Value.(*cacheItem)
		if !ok {
			return nil, false
		}

		return cachedItem.value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = lru.NewList()
	c.items = make(map[string]*lru.Item, c.capacity)
}

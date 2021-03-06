package cache

import (
	"fmt"
	"sync"
)

type lruCache struct {
	capacity int
	queue    List
	items    map[string]*ListItem
	mu       *sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[string]*ListItem, capacity),
		mu:       new(sync.Mutex),
	}
}

func (c *lruCache) Set(i *Item) bool {
	c.mu.Lock()

	defer c.mu.Unlock()

	if lruItem, ok := c.items[i.Key]; ok {
		lruItem.Value = i
		c.queue.MoveToFront(lruItem)

		return true
	}

	if c.queue.Len() == c.capacity {
		lastItem, ok := c.queue.Back().Value.(*Item)
		if !ok {
			return false
		}

		c.queue.Remove(c.queue.Back())
		delete(c.items, lastItem.Key)
	}

	pushed := c.queue.PushFront(i)
	c.items[i.Key] = pushed

	return false
}

func (c *lruCache) Get(key string) (*Item, bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		cachedItem, ok := item.Value.(*Item)
		if !ok {
			return nil, false
		}

		return cachedItem, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[string]*ListItem, c.capacity)
}

func (c *lruCache) MakeCacheKeyResizes(width, height int, url string) string {
	return fmt.Sprintf("%d_%d_%s", width, height, url)
}

func (c *lruCache) MakeCacheKeyDownloaded(url string) string {
	return url
}

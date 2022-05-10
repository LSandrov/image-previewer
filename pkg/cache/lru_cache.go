package cache

import (
	"fmt"
	"image-previewer/pkg/cache/lru"
	"sync"
)

type lruCache struct {
	capacity int
	queue    lru.List
	items    map[string]*lru.Item
	mu       *sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    lru.NewList(),
		items:    make(map[string]*lru.Item, capacity),
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

		cachedItem, okCache := item.Value.(*Item)
		if okCache != true {
			return nil, false
		}

		return cachedItem, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = lru.NewList()
	c.items = make(map[string]*lru.Item, c.capacity)
}

func (c *lruCache) MakeCacheKeyResizes(width, height int, url string) string {
	return fmt.Sprintf("%d_%d_%s", width, height, url)
}

func (c *lruCache) MakeCacheKeyDownloaded(url string) string {
	return fmt.Sprintf("%s", url)
}

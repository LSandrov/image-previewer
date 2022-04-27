package cache

import (
	"image-previewer/pkg/cache/lru"
	"sync"
)

type LruCache struct {
	capacity int
	queue    *lru.List
	items    map[string]*lru.ListItem
	mu       *sync.Mutex
}

func NewLruCache(capacity int, queue *lru.List, items map[string]*lru.ListItem, mu *sync.Mutex) Cache {
	return &LruCache{capacity: capacity, queue: queue, items: items, mu: mu}
}

func (c *LruCache) Get(key string) (val []byte, err error) {
	//@TODO fixme
	return nil, nil
}
func (c *LruCache) Set(key string, val []byte) (err error) {
	//@TODO fixme
	return nil
}

package design

import (
	"container/list"
	"errors"
	"sync"
)

type Cache interface {
	Get(key any) (value any, ok bool)
	Put(key, value any) error
	Remove(key any) bool
	Clear()
	Len() int
	Cap() int
}

type entry struct {
	key   any
	value any
}

type LRUCache struct {
	mu      sync.RWMutex
	cap     int
	cache   map[any]*list.Element
	lruList *list.List
	onEvict func(key, value any)
}

type Option func(*LRUCache)

func New(capacity int, opts ...Option) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be positive")
	}

	c := &LRUCache{
		cap:     capacity,
		cache:   make(map[any]*list.Element, capacity),
		lruList: list.New(),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *LRUCache) Get(key any) (value any, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, exists := c.cache[key]; exists {
		c.lruList.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

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

func (c *LRUCache) Put(key, value any) error {
	if key == nil {
		return errors.New("key cannot be nil")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.cache[key]; exists {
		c.lruList.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return nil
	}

	ent := &entry{key: key, value: value}
	elem := c.lruList.PushFront(ent)
	c.cache[key] = elem

	if c.lruList.Len() > c.cap {
		c.evictOldest()
	}
	return nil
}

func (c *LRUCache) Remove(key any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.cache[key]; exists {
		c.removeElement(elem)
		return true
	}
	return false
}

func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.onEvict != nil {
		for key, elem := range c.cache {
			ent := elem.Value.(*entry)
			c.onEvict(key, ent.value)
		}
	}

	c.cache = make(map[any]*list.Element)
	c.lruList.Init()
}

func (c *LRUCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.cache)
}

func (c *LRUCache) Cap() int {
	return c.cap
}

func (c *LRUCache) Keys() []any {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]any, 0, len(c.cache))
	for elem := c.lruList.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(*entry).key)
	}
	return keys
}

func (c *LRUCache) evictOldest() {

}

package design

import "container/list"

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
	cap     int
	cache   map[any]*list.Element
	lruList *list.List
	onEvict func(key, value any)
}

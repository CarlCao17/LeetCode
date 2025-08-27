package design

import "errors"

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	key  K
	val  V
}

type list[K comparable, V any] struct {
	size int
	head *Node[K, V]
	tail *Node[K, V]
}

func NewList[K comparable, V any]() list[K, V] {
	return list[K, V]{}
}

type LRU[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	list[K, V]
	onEvict func(key K, val V)
}

type Option[K comparable, V any] func(*LRU[K, V])

func WithEvictCallback[K comparable, V any](callback func(key K, val V)) Option[K, V] {
	return func(l *LRU[K, V]) {
		l.onEvict = callback
	}
}

func New[K comparable, V any](capacity int, opts ...Option[K, V]) (*LRU[K, V], error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be positive")
	}

	lruList := NewList[K, V]()

	lru := &LRU[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
		lruList:  lruList,
	}

	for _, opt := range opts {
		opt(lru)
	}
	return lru, nil
}

func (lru *LRU[K, V]) Get(key K) (v V, exists bool) {
	node, exists := lru.cache[key]
	if !exists {
		return v, exists
	}
	lru.moveToHead(node)
	return node.val, true
}

func (lru *LRU[K, V]) Put(key K, val V) {
	if node, exists := lru.cache[key]; exists {
		node.val = val
		lru.moveToHead(node)
		return
	}

	newNode := &Node[K, V]{
		key: key,
		val: val,
	}

	lru.cache[key] = newNode
	lru.addToHead(newNode)

}

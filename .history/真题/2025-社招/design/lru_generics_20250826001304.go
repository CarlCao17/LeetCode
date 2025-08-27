package design

import "errors"

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

	lru := &LRU[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
		list:     NewList[K, V](),
	}

	for _, opt := range opts {
		opt(lru)
	}
	return lru, nil
}

func (lru *LRU[K, V]) Get(key K) (v V, exists bool) {
	node, exists := lru.cache[key]
	if !exists {
		return v, false
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
	lru.size++

	if lru.size > lru.capacity {
		lru.evictOldest()
	}
}

func (lru *LRU[K, V]) Remove(key K) bool {
	node, exists := lru.cache[key]
	if !exists {
		return false
	}

	lru.removeNode(node)
	lru.size--
	delete(lru.cache, key)

	if lru.onEvict != nil {
		lru.onEvict(key, node.val)
	}
	return true
}

func (lru *LRU[K, V]) Clear() {
	if lru.onEvict != nil {
		lru.ForEach[K, V](func(node *Node[K, V]) bool {
			lru.onEvict(node.key, node.val)
			return true
		})
	}

	lru.cache = make(map[K]*Node[K, V])
	lru.list = NewList[K, V]()
}

// Peek 查看值但不更新LRU顺序
func (lru *LRU[K, V]) Peek(key K) (v V, exists bool) {
	node, exists := lru.cache[key]
	if !exists {
		return v, false
	}

	return node.val, true
}

// Contains 检查key是否存在
func (lru *LRU[K, V]) Contains(key K) bool {
	_, exists := lru.cache[key]
	return exists
}

// Len 返回当前大小
func (lru *LRU[K, V]) Len() int {
	return lru.size
}

// Cap 返回容量
func (lru *LRU[K, V]) Cap() int {
	return lru.capacity
}

// Keys 返回所有key（从最新到最旧）
func (lru *LRU[K, V]) Keys() []K {
	keys := make([]K, 0, lru.size)
	lru.ForEach[K, V](func(node *Node[K, V]) bool {
		keys = append(keys, node.key)
		return true
	})
	return keys
}

// Values 返回所有value（从最新到最旧）
func (lru *LRU[K, V]) Values() []V {
	values := make([]V, 0, lru.size)
	lru.ForEach[K, V](func(node *Node[K, V]) bool {
		values = append(values, node.val)
		return true
	})
	return values
}

// Entries 返回所有键值对（从最新到最旧）
func (lru *LRU[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], 0, lru.size)
	lru.ForEach[K, V](func(node *Node[K, V]) bool {
		entries = append(entries, node.Entry)
		return true
	})
	return entries
}

type Entry[K comparable, V any] struct {
	Key K
	Val V
}

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	Entry[K, V]
}

type list[K comparable, V any] struct {
	size int
	head *Node[K, V]
	tail *Node[K, V]
}

func NewList[K comparable, V any]() list[K, V] {
	return list[K, V]{
		head: &Node[K, V]{},
		tail: &Node[K, V]{},
	}
}

func (l *list[K, V]) addToHead(node *Node[K, V]) {
	node.next = l.head.next
	l.head.next.prev = node

	l.head.next = node
	node.prev = l.head
}

func (l *list[K, V]) moveToHead(node *Node[K, V]) {

	l.addToHead(node)
}

// node 指针由使用方来保证消除
func (l *list[K, V]) removeNode(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// node.next = nil
	// node.prev = nil
}

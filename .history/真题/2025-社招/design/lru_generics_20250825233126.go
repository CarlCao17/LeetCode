package design

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

type LRU[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	lruList  list[K, V]
	onEvict  func(key K, val V)
}

type Option[K comparable, V any] func(*LRU[K, V])

func WithEvictCallback[K comparable, V any](callback func(key K, val V)) Option[K, V] {
	return func(l *LRU[K, V]) {
		l.onEvict = callback
	}
}

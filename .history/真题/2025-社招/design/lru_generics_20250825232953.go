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

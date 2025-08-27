package design

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	key  K
	val  V
}

package maps

import "sort"

func NewIntMap[V any](m map[int]V) *Map[int, V] {
	return NewMap[int, V](m, func(x, y *int) bool {
		return *x <= *y
	})
}

type Map[K comparable, V any] struct {
	m     map[K]V
	order *order[K]
}

type order[K comparable] struct {
	s  []K
	by By[K]
}

func (o *order[K]) Len() int {
	return len(o.s)
}

func (o *order[K]) Less(i, j int) bool {
	return o.by(&o.s[i], &o.s[j])
}

func (o *order[K]) Swap(i, j int) {
	o.s[i], o.s[j] = o.s[j], o.s[i]
}

type By[K comparable] func(x, y *K) bool

func NewMap[K comparable, V any](m map[K]V, less func(x, y *K) bool) *Map[K, V] {
	res := &Map[K, V]{
		m: m,
	}
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	o := &order[K]{s: keys, by: less}
	sort.Sort(o)
	res.order = o
	return res
}

func (m *Map[K, V]) Keys() []K {
	return m.order.s
}

func (m *Map[K, V]) Items() ([]K, []V) {
	v := make([]V, len(m.m))
	for i, k := range m.order.s {
		v[i] = m.m[k]
	}
	return m.order.s, v
}

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

func (m *Map[K, V]) Items2() <-chan Item[K, V] {
	c := make(chan Item[K, V])
	go func() {
		defer close(c)
		for _, k := range m.order.s {
			v, _ := m.m[k]
			c <- Item[K, V]{k, v}
		}
	}()
	return c
}

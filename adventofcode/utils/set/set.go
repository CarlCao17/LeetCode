package set

var defaultSize = 32

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](members ...T) *Set[T] {
	set := &Set[T]{}

	set.m = make(map[T]struct{}, If(len(members) == 0, defaultSize, len(members)))
	for _, mem := range members {
		set.m[mem] = struct{}{}
	}
	return set
}

func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	}
	return onFalse
}

func (s *Set[T]) Add(t T) (ok bool) {
	_, ok = s.m[t]
	s.m[t] = struct{}{}
	return ok
}

func (s *Set[T]) AddN(members ...T) {
	for _, m := range members {
		s.m[m] = struct{}{}
	}
}

func (s *Set[T]) AddAll(slices ...[]T) {
	for _, slice := range slices {
		for _, ss := range slice {
			s.m[ss] = struct{}{}
		}
	}
}

func (s *Set[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

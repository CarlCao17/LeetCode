package slices

import (
	"github.com/CarlCao17/LeetCode/adventofcode/utils/constraints"
)

func Sum[T constraints.RealNumber](s []T) T {
	return Reduce(s, RFSum[T])
}

func Max[T constraints.RealNumber](s []T) T {
	return Reduce(s, RFMax[T])
}

func Min[T constraints.RealNumber](s []T) T {
	if len(s) == 0 {
		panic("Min: should have length at least one")
	}
	if len(s) == 1 {
		return s[0]
	}
	return ReduceWithAcc(s, s[0], RFMin[T])
}

func Reduce[T any](s []T, f func(acc T, item T) T) T {
	var acc T
	return ReduceWithAcc(s, acc, f)
}

func ReduceWithAcc[T any](s []T, acc T, f func(acc T, item T) T) T {
	for _, item := range s {
		acc = f(acc, item)
	}
	return acc
}

func MapReduce[T any, V any](s []T, mapFunc func(T) V, initAcc V, reduceFunc func(acc V, item V) V) V {
	acc := initAcc
	for _, item := range s {
		acc = reduceFunc(acc, mapFunc(item))
	}
	return acc
}

type ReduceFunc[T any] func(acc T, item T) T

type MapFunc[T any, V any] func(item T) V

func RFSum[T constraints.RealNumber](acc T, item T) T {
	return acc + item
}

func RFMax[T constraints.RealNumber](acc T, item T) T {
	if item > acc {
		return item
	}
	return acc
}

func RFMin[T constraints.RealNumber](acc T, item T) T {
	if item < acc {
		return item
	}
	return acc
}

func Map[T any, V any](s []T, mapFunc func(T) V) []V {
	res := make([]V, len(s))
	for i, ss := range s {
		res[i] = mapFunc(ss)
	}
	return res
}

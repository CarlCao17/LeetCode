package array

import (
	"math/rand"
	"strconv"
	"strings"
)

const seed = 81192

func NewRandArray(n int, low, high int) []int {
	r := rand.New(rand.NewSource(seed))
	slice := make([]int, n)
	for i := range slice {
		slice[i] = int(r.Int63n(int64(high-low))) + low
	}
	return slice
}

func NewSortArrayAsString(n int) string {
	return PrintArray(NewSortArray(n))
}

func NewSortArray(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

func NewDupArray(n int, val int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = val
	}
	return slice
}

func NewRandArrayAsString(n int, low, high int) string {
	return PrintArray(NewRandArray(n, low, high))
}

func PrintArray(a []int) string {
	var b strings.Builder
	b.WriteRune('[')
	for i := 0; i < len(a); i++ {
		if i > 0 {
			b.WriteRune(',')
		}
		b.WriteString(strconv.Itoa(a[i]))
	}
	b.WriteRune(']')
	return b.String()
}

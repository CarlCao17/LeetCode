package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(rand.Intn(5))
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	n := len(nums) - 1
	l, r := 0, len(nums)-1
	for n >= 0 {
		lsq := nums[l] * nums[l]
		rsq := nums[r] * nums[r]
		if lsq >= rsq {
			res[n] = lsq
			l++
		} else {
			res[n] = rsq
			r--
		}
		n--
	}
	return res
}

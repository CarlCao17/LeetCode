package main

import (
	"fmt"
	"math"
	"os"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func main() {
	file, err := os.Open("./day_06/input.txt")
	if err != nil {
		panic(err)
	}
	records := GetInput(file)
	nums := getNumOfWays(records)
	fmt.Printf("num of ways = %v, multiply=%d\n", nums, slices.Multi(nums))
}

func getNumOfWays(records []Record) []int {
	res := make([]int, len(records))
	for i, record := range records {
		res[i] = getNumOfSolutions(record.Time, record.Distance)
	}
	return res
}

// 相当于求解
// 满足 x^2 -t*x + d <= 0 的正整数的个数，x 表示充电的时间
func getNumOfSolutions(t int, d int) int {
	delta2 := t*t - 4*d
	if delta2 < 0 {
		return 0
	}
	if delta2 == 0 {
		return 1
	}
	halfOfDelta := math.Sqrt(float64(delta2)) / 2
	x1, x2 := float64(t)/2-halfOfDelta, float64(t)/2+halfOfDelta
	n1, n2 := int(x1+1), int(math.Ceil(x2-1)) // n1 是 > x1 的最小整数，n2 是 < x2 的最小整数
	return n2 - n1 + 1
}

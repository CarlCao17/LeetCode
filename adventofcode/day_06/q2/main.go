package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func main() {
	file, err := os.Open("./day_06/input.txt")
	if err != nil {
		panic(err)
	}
	t, d := GetInput(file)
	fmt.Printf("time=%d, distance=%d, solution=%d\n", t, d, getNumOfSolutions(t, d))
}

// 相当于求解
// 满足 x^2 -t*x + d <= 0 的正整数的个数，x 表示充电的时间
func getNumOfSolutions(t int64, d int64) int64 {
	delta2 := t*t - 4*d
	if delta2 < 0 {
		return 0
	}
	if delta2 == 0 {
		return 1
	}
	halfOfDelta := math.Sqrt(float64(delta2)) / 2
	x1, x2 := float64(t)/2-halfOfDelta, float64(t)/2+halfOfDelta
	n1, n2 := int64(x1+1), int64(math.Ceil(x2-1)) // n1 是 > x1 的最小整数，n2 是 < x2 的最小整数
	return n2 - n1 + 1
}

func GetInput(r io.Reader) (t int64, d int64) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	timeLine := scanner.Text()
	t = slices.StrToInt64(strings.ReplaceAll(timeLine[5:], " ", ""))
	scanner.Scan()
	distanceLine := scanner.Text()
	d = slices.StrToInt64(strings.ReplaceAll(distanceLine[9:], " ", ""))
	return
}

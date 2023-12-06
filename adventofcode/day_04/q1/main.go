package main

import (
	"fmt"

	. "github.com/CarlCao17/LeetCode/adventofcode/day_04"
)

func main() {
	cards := GetInput()
	wp := getWinningPoints(cards)
	fmt.Printf("winning points = %d\n", wp)
}

func getWinningPoints(cards []Card) int64 {
	var points int64
	for _, card := range cards {
		p := int64(0)
		cnt := 0
		for _, n := range card.Numbers {
			if card.WiningNumbers.Contains(n) {
				if p == 0 {
					p = 1
					continue
				}
				p <<= 1
				cnt++
			}
		}
		assertNoOverflow(cnt)
		points += p
	}
	assertSumNoOverflow(points)
	return points
}

func assertNoOverflow(shift int) {
	if shift > 63 {
		panic(fmt.Sprintf("left shift should be not greater than 63, curr=%d", shift))
	}
}

func assertSumNoOverflow(n int64) {
	if n < 0 {
		panic(fmt.Sprintf("sum of winning points overflow, curr=%d", n))
	}
}

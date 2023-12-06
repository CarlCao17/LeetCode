package main

import (
	"fmt"

	. "github.com/CarlCao17/LeetCode/adventofcode/day_04"
	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func main() {
	cards := GetInput()

	scratchCardNums := getScratchCardNums(cards)
	fmt.Printf("total scratch card = %v, sum=%d\n", scratchCardNums, slices.Sum(scratchCardNums))
}

func getScratchCardNums(cards []Card) []int {
	scratchCardNums := slices.NewSliceWith(1, len(cards))
	scratchCardNums[0] = 1
	for idx, num := range scratchCardNums {
		card := &cards[idx]
		matchingNum := GetMatchingNumber(card)
		for next := 1; next < len(cards) && next <= matchingNum; next++ {
			scratchCardNums[idx+next] += num
		}
	}
	return scratchCardNums
}

func GetMatchingNumber(c *Card) int {
	return slices.MapReduce(c.Numbers, func(num int) int {
		if c.WiningNumbers.Contains(num) {
			return 1
		}
		return 0
	}, 0, slices.RFSum[int])
}

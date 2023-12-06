package day_04

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/set"
)

func GetInput() []Card {
	file, err := os.Open("./day_04/input.txt")
	if err != nil {
		panic(err)
	}
	cards := parseInput(file)
	return cards
}

type Card struct {
	CardNum       int
	WiningNumbers *set.Set[int]
	Numbers       []int
}

func parseInput(r io.Reader) []Card {
	res := make([]Card, 0, 200) // from the puzzle input
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		card := parseLine(line)
		res = append(res, card)
	}
	return res
}

func parseLine(s string) Card {
	card := Card{}
	idx := strings.Index(s, ":")
	fmt.Sscanf(s[:idx], "Card %d", &card.CardNum)
	s = s[idx+1:]
	parts := strings.Split(s, "|")
	assertHaveLenTwo(parts)
	win, nums := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	card.WiningNumbers = set.New(slices.Map(strings.Fields(win), slices.StrToInt)...)
	card.Numbers = slices.Map(strings.Fields(nums), slices.StrToInt)
	return card
}

func assertHaveLenTwo(s []string) {
	if len(s) != 2 {
		panic(fmt.Sprintf("the string slice should have len 2, got=%d", len(s)))
	}
}

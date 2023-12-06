package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils"
	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

type Record struct {
	Time     int
	Distance int
}

func GetInput(r io.Reader) []Record {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	timeLine := scanner.Text()
	times := slices.Map(strings.Fields(strings.TrimSpace(timeLine[5:])), slices.StrToInt)
	scanner.Scan()
	distanceLine := scanner.Text()
	distances := slices.Map(strings.Fields(strings.TrimSpace(distanceLine[9:])), slices.StrToInt)
	utils.AssertF(len(times) == len(distances), "time line length should equal to distance line, time=%d, distance=%d", len(times), len(distances))

	res := make([]Record, len(times))
	for i := 0; i < len(times); i++ {
		res[i] = Record{
			Time:     times[i],
			Distance: distances[i],
		}
	}
	return res
}

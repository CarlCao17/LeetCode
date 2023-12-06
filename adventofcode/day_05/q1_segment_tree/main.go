package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func main() {
	var r io.Reader
	file, err := os.Open("./day_05/input.txt")
	if err != nil {
		r = os.Stdin // for debug
	} else {
		r = file
	}
	solution := parseInput(r)
	locations := findLocations(solution)
	result := slices.Min(locations)
	fmt.Printf("locations=%v, result=%d\n", locations, result)
}

func findLocations(s Solution) []int64 {
	res := make([]int64, len(s.Seeds))
	for i, seed := range s.Seeds {
		res[i] = s.SeedToLocation(seed)
	}
	return res
}

type Solution struct {
	Seeds                 []int64
	seedToSoil            *TreeNode
	soilToFertilizer      *TreeNode
	fertilizerToWater     *TreeNode
	waterToLight          *TreeNode
	lightToTemperature    *TreeNode
	temperatureToHumidity *TreeNode
	humidityToLocation    *TreeNode
}

func (s *Solution) SeedToLocation(seed int64) int64 {
	soil := s.seedToSoil.Search(seed)
	fertilizer := s.soilToFertilizer.Search(soil)
	water := s.fertilizerToWater.Search(fertilizer)
	light := s.waterToLight.Search(water)
	temperature := s.lightToTemperature.Search(light)
	humidity := s.temperatureToHumidity.Search(temperature)
	location := s.humidityToLocation.Search(humidity)
	return location
}

// TreeNode 假设线段树区间没有交集
type TreeNode struct {
	lChild *TreeNode
	rChild *TreeNode
	srcToDst
}

func (n *TreeNode) String() string {
	if n == nil {
		return "<not exist node>"
	}
	var s strings.Builder
	s.WriteString("[")
	n.string(&s)
	s.WriteString("]")
	return s.String()
}

func (n *TreeNode) string(s *strings.Builder) {
	if n == nil {
		return
	}
	n.lChild.string(s)
	s.WriteString(fmt.Sprintf("{%d %d %d}", n.dstStart, n.srcStart, n.len))
	n.rChild.string(s)
}

func (n *TreeNode) Insert(o *TreeNode) *TreeNode {
	if n == nil {
		return o
	}
	if n.srcStart >= o.srcStart+o.len {
		n.lChild = n.lChild.Insert(o)
	} else if o.srcStart >= n.srcStart+n.len {
		n.rChild = n.rChild.Insert(o)
	} else {
		panic(fmt.Sprintf("should not have interleave between %s and %s", n, o))
	}
	return n
}

func (n *TreeNode) Search(num int64) int64 {
	o := n.search(num)
	if o == nil {
		return num
	}
	return o.dstStart + num - o.srcStart
}

func (n *TreeNode) search(num int64) *TreeNode {
	if n == nil || num >= n.srcStart && num < n.srcStart+n.len {
		return n
	}
	if num < n.srcStart {
		return n.lChild.search(num)
	}
	return n.rChild.search(num)
}

type srcToDst struct {
	dstStart int64
	srcStart int64
	len      int64
}

func parseLine(s string) (res srcToDst) {
	s = strings.TrimSpace(s)
	fmt.Sscanf(s, "%d %d %d", &res.dstStart, &res.srcStart, &res.len) //nolint
	return res
}

func parseInput(r io.Reader) Solution {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanMoreNewlines)

	solution := Solution{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line[:5] == "seeds" {
			seeds := slices.Map(strings.Split(strings.TrimSpace(line[6:]), " "), func(v string) int64 {
				parsed, _ := strconv.ParseInt(v, 10, 64)
				return parsed
			})
			solution.Seeds = seeds
			continue
		}

		lines := strings.Split(line, "\n")
		if len(lines) > 0 {
			lines = lines[1:] // delete the cmd line
		}
		tree := parseToSegmentTree(lines)
		switch line[:5] {
		case "seed-":
			solution.seedToSoil = tree
		case "soil-":
			solution.soilToFertilizer = tree
		case "ferti":
			solution.fertilizerToWater = tree
		case "water":
			solution.waterToLight = tree
		case "light":
			solution.lightToTemperature = tree
		case "tempe":
			solution.temperatureToHumidity = tree
		case "humid":
			solution.humidityToLocation = tree
		}
	}
	return solution
}

func parseToSegmentTree(lines []string) *TreeNode {
	var tree *TreeNode
	for _, line := range lines {
		segmentVal := parseLine(line)
		tree = tree.Insert(&TreeNode{srcToDst: segmentVal})
	}
	return tree
}

// ScanMoreNewlines 以多个换行符为划分
func ScanMoreNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 1, dropCRs(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCRs(data), nil
	}
	return 0, nil, nil
}

func dropCRs(data []byte) []byte {
	last := len(data) - 1
	for last >= 0 && data[last] == 'r' {
		last--
	}
	return data[0 : last+1]
}

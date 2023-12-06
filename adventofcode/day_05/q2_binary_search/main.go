package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

// q2 use binary search because the q1_segment_tree might not be balanced
func main() {
	var r io.Reader
	file, err := os.Open("./day_05/input.txt")
	if err != nil {
		r = os.Stdin // for debug
	} else {
		r = file
	}
	solution := parseInput(r)
	result := findMinLocations(solution)
	fmt.Printf("result=%d\n", result)
}

func findMinLocations(s Solution) int64 {
	res := int64(math.MaxInt64)
	for _, iv := range s.Seeds {
		for i := iv.start; i < iv.start+iv.len; i++ {
			res = min(res, s.SeedToLocation(i))
		}
	}
	return res
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

type Solution struct {
	//Seeds                 []int64  for q1
	Seeds                 []interval // for q2
	seedToSoil            *OrderIntervals
	soilToFertilizer      *OrderIntervals
	fertilizerToWater     *OrderIntervals
	waterToLight          *OrderIntervals
	lightToTemperature    *OrderIntervals
	temperatureToHumidity *OrderIntervals
	humidityToLocation    *OrderIntervals
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

type OrderIntervals struct {
	ordered   bool
	intervals []*srcToDst // Supposed that the interval have no interleave
}

func (oi *OrderIntervals) Sort() {
	sort.Slice(oi.intervals, func(i, j int) bool {
		l, r := oi.intervals[i], oi.intervals[j]
		return l.src.start+l.src.len <= r.src.start
	})
	oi.ordered = true
}

func (oi *OrderIntervals) Search(num int64) int64 {
	index := oi.BinarySearch(num)
	if index < 0 {
		return num
	}
	iv := oi.intervals[index]
	return num + (iv.dstStart - iv.src.start)
}

// BinarySearch will return the index if num is in the intervals,
// otherwise return -l-1, l is the interval index which it should be in
func (oi *OrderIntervals) BinarySearch(num int64) int {
	l, r := 0, len(oi.intervals)
	for l < r {
		m := l + (r-l)/2
		mv := oi.intervals[m].src
		if num < mv.start {
			r = m
		} else if num >= mv.start+mv.len {
			l = m + 1
		} else {
			return m
		}
	}
	return -l - 1
}

func (oi *OrderIntervals) String() string {
	if !oi.ordered {
		oi.Sort()
	}
	var s strings.Builder
	s.WriteString("{")
	for i, iv := range oi.intervals {
		if i > 0 {
			s.WriteByte(' ')
		}
		s.WriteString(fmt.Sprintf("{%d %d %d}", iv.dstStart, iv.src.start, iv.src.len))
	}
	s.WriteString("}")
	return s.String()
}

type srcToDst struct {
	dstStart int64
	src      interval
}

type interval struct {
	start int64
	len   int64
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
			assertEven(len(seeds))
			for i := 0; i < len(seeds); i += 2 {
				solution.Seeds = append(solution.Seeds, interval{start: seeds[i], len: seeds[i+1]})
			}
			continue
		}

		lines := strings.Split(line, "\n")
		if len(lines) > 0 {
			lines = lines[1:] // delete the cmd line
		}
		orderIv := parseLines(lines)
		switch line[:5] {
		case "seed-":
			solution.seedToSoil = orderIv
		case "soil-":
			solution.soilToFertilizer = orderIv
		case "ferti":
			solution.fertilizerToWater = orderIv
		case "water":
			solution.waterToLight = orderIv
		case "light":
			solution.lightToTemperature = orderIv
		case "tempe":
			solution.temperatureToHumidity = orderIv
		case "humid":
			solution.humidityToLocation = orderIv
		}
	}

	return solution
}

func parseLines(lines []string) *OrderIntervals {
	res := &OrderIntervals{
		intervals: make([]*srcToDst, 0, 16),
	}
	for _, line := range lines {
		iv := parseLine(line)
		res.intervals = append(res.intervals, iv)
	}
	res.Sort()
	return res
}

func parseLine(s string) *srcToDst {
	res := &srcToDst{}
	s = strings.TrimSpace(s)
	fmt.Sscanf(s, "%d %d %d", &res.dstStart, &res.src.start, &res.src.len) //nolint
	return res
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

func assertEven(n int) {
	if n&1 == 1 {
		panic(fmt.Sprintf("the number should be even, n=%d", n))
	}
}

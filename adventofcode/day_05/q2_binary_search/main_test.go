package main

import (
	"sort"
	"testing"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func TestOrderIntervals_BinarySearch(t *testing.T) {
	lines := []string{
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
	}
	// [0,7), [7,11), [11, 52), [53, 61)
	oi := parseLines(lines)
	if got := oi.BinarySearch(0); got != 0 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 0, got, 0)
	}
	if got := oi.BinarySearch(3); got != 0 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 3, got, 0)
	}
	if got := oi.BinarySearch(62); got != -5 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 62, got, -5)
	}
	if got := oi.BinarySearch(61); got != -5 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 61, got, -5)
	}
	if got := oi.BinarySearch(45); got != 2 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 45, got, 2)
	}
	if got := oi.BinarySearch(11); got != 2 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 11, got, 2)
	}
	if got := oi.BinarySearch(45); got != 2 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 45, got, 2)
	}
	if got := oi.BinarySearch(52); got != 2 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 52, got, 2)
	}
	if got := oi.BinarySearch(53); got != 3 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 53, got, 3)
	}
	if got := oi.BinarySearch(7); got != 1 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 7, got, 1)
	}
	if got := oi.BinarySearch(11); got != 2 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 11, got, 2)
	}

	lines = []string{
		"49 53 8",
		"0 11 30",
		"42 1 6",
		"57 7 4",
		"50 50 2",
	}
	// [1,7), [7,11), [11, 41), [50, 52), [53, 61)
	oi = parseLines(lines)
	if got := oi.BinarySearch(0); got != -1 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 0, got, -1)
	}
	if got := oi.BinarySearch(61); got != -6 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 61, got, -6)
	}
	if got := oi.BinarySearch(48); got != -4 {
		t.Errorf("BinarySearch %d: got=%d, expect=%d", 61, got, -4)
	}
}

func TestOrderIntervals_Sort(t *testing.T) {
	lines := []string{
		"3547471595 1239929038 174680800",
		"3052451552 758183681 481745357",
		"0 1427884524 1775655006",
		"2844087171 549819300 208364381",
		"3767989253 4004864866 5194940",
		"3534196909 1414609838 13274686",
		"1775655006 114264781 435554519",
		"4148908402 4010059806 146058894",
		"2729822390 0 114264781",
		"3773184193 4156118700 138848596",
		"2211209525 3203539530 518612865",
		"3912032789 3767989253 236875613",
	}
	oi := parseLines(lines)
	got := slices.Map(oi.intervals, func(t *srcToDst) int64 {
		return t.src.start
	})
	if !sort.SliceIsSorted(got, func(i, j int) bool {
		return got[i] < got[j]
	}) {
		t.Errorf("OrderIntervals sort failed: %v", got)
	}
}

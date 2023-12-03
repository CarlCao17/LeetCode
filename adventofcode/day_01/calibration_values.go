package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readInput(r io.Reader) (res []string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		if t == "" {
			continue
		}
		res = append(res, t)
	}
	return res
}

func main() {
	file, err := os.Open("/Users/caozhengcheng/go/src/leetcode/adventofcode/day_01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	input := readInput(file)

	numbers := make([]int, len(input))
	for _, s := range input {
		f, l := findDigits(s)
		number := f*10 + l
		numbers = append(numbers, number)
	}
	fmt.Printf("numbers=%v, sum=%d", numbers, sum(numbers))
}

func findDigits(s string) (f int, l int) {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			f = int(s[i] - '0')
			break
		}
		if i <= len(s)-5 {
			t := ContainLetterDigits(s[i:i+5], 5)
			if t != -1 {
				f = t
				break
			}
		}
		if i <= len(s)-4 {
			t := ContainLetterDigits(s[i:i+4], 4)
			if t != -1 {
				f = t
				break
			}
		}
		if i <= len(s)-3 {
			t := ContainLetterDigits(s[i:i+3], 3)
			if t != -1 {
				f = t
				break
			}
		}
	}
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] >= '0' && s[j] <= '9' {
			l = int(s[j] - '0')
			break
		}
		if j <= len(s)-3 {
			t := ContainLetterDigits(s[j:j+3], 3)
			if t != -1 {
				l = t
				break
			}
		}
		if j <= len(s)-4 {
			t := ContainLetterDigits(s[j:j+4], 4)
			if t != -1 {
				l = t
				break
			}
		}
		if j <= len(s)-5 {
			t := ContainLetterDigits(s[j:j+5], 5)
			if t != -1 {
				l = t
				break
			}
		}
	}
	return f, l
}

var (
	digitStr3 = []string{
		"one", "two", "six",
	}
	m3 = []int{
		1, 2, 6,
	}
	digitStr4 = []string{
		"zero", "four", "five", "nine",
	}
	m4 = []int{
		0, 4, 5, 9,
	}
	digitStr5 = []string{
		"three", "seven", "eight",
	}
	m5 = []int{
		3, 7, 8,
	}
)

func ContainLetterDigits(s string, n int) int {
	digitStrs := [][]string{
		3: digitStr3, 4: digitStr4, 5: digitStr5,
	}
	ms := [][]int{
		3: m3, 4: m4, 5: m5,
	}
	digitStr := digitStrs[n]
	m := ms[n]
	idx := ContainString(digitStr, s)
	if idx == -1 {
		return -1
	}
	return m[idx]
}

func ContainString(s []string, t string) int {
	for i, ss := range s {
		if ss == t {
			return i
		}
	}
	return -1
}

func sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

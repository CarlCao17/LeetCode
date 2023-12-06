package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils"
)

type Cubes struct {
	r int
	g int
	b int
}

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	var r io.Reader
	file, err := os.Open("./day_02/input.txt")
	if err != nil {
		r = os.Stdin
	} else {
		r = file
	}

	games := parseInput(r)
	res := make([]int, 0, len(games))
	for i, set := range games {
		possible := true
		for _, cube := range set {
			if cube.r > maxRed || cube.g > maxGreen || cube.b > maxBlue {
				possible = false
				break
			}
		}
		if possible {
			res = append(res, i)
		}
	}
	fmt.Printf("res = %v, sum=%d\n", res, utils.Sum(res))
}

func parseInput(r io.Reader) map[int][]Cubes {
	res := make(map[int][]Cubes)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		s := strings.SplitN(text, ":", 2)
		var idx int
		fmt.Sscanf(s[0], "Game %d", &idx)
		text = s[1]
		cubes := parseLine(text)
		res[idx] = cubes
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return res
}

func parseLine(line string) []Cubes {
	res := make([]Cubes, 0, 8)
	splits := strings.Split(line, ";")
	for _, s := range splits {
		ss := strings.Split(s, ",")
		var cube Cubes
		for _, sss := range ss {
			sss = strings.TrimSpace(sss)
			var num int
			var color string
			fmt.Sscanf(sss, "%d %s", &num, &color)
			switch color {
			case "blue":
				cube.b = num
			case "red":
				cube.r = num
			case "green":
				cube.g = num
			default:
				panic(fmt.Sprintf("color: %s, num: %d", color, num))
			}
		}
		res = append(res, cube)
	}
	return res
}

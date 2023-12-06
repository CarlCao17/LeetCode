package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils"
)

type Cube struct {
	r int
	g int
	b int
}

func main() {
	var r io.Reader
	file, err := os.Open("./day_02/input.txt")
	if err != nil {
		r = os.Stdin
	} else {
		r = file
	}

	games := parseInput(r)
	orderM := utils.NewIntMap(games)
	res := int64(0)
	for item := range orderM.Items2() {
		no, cubes := item.Key, item.Value
		cube := maxCube(cubes)
		fmt.Printf("Game %d red: %d, green: %d, blue: %d\n", no, cube.r, cube.g, cube.b)
		res += int64(cube.r * cube.g * cube.b)
	}
	fmt.Println(res)
}

func maxCube(cubes []Cube) Cube {
	res := Cube{}
	for _, cube := range cubes {
		if cube.r > res.r {
			res.r = cube.r
		}
		if cube.g > res.g {
			res.g = cube.g
		}
		if cube.b > res.b {
			res.b = cube.b
		}
	}
	return res
}

func parseInput(r io.Reader) map[int][]Cube {
	res := make(map[int][]Cube)
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

func parseLine(line string) []Cube {
	res := make([]Cube, 0, 8)
	splits := strings.Split(line, ";")
	for _, s := range splits {
		ss := strings.Split(s, ",")
		var cube Cube
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

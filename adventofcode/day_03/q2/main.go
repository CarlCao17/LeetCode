package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/CarlCao17/LeetCode/adventofcode/utils/slices"
)

func main() {
	file, err := os.Open("./day_03/input.txt")
	if err != nil {
		panic(err)
	}
	nodes, n := parseInput(file)
	PrintMatrix(nodes)
	partNumbers := getPartNumbers(nodes, n)
	gears := getGears(nodes, partNumbers)
	sum := slices.Sum(gears)
	fmt.Printf("Sum of gear ratios %v = %d\n", gears, sum)
}

func PrintMatrix(nodes [][]Node) {
	var sb strings.Builder
	for i := 1; i < len(nodes)-1; i++ {
		row := nodes[i]
		for j := 1; j < len(row)-1; j++ {
			sb.WriteString(row[j].String())
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func getPartNumbers(matrix [][]Node, n int) []*Node {
	res := make([]*Node, 0, n)
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i][j].kind == KindNumber && isAdjacentToSymbol(matrix, i, j) {
				res = append(res, &matrix[i][j])
			}
		}
	}
	return res
}

func getGears(matrix [][]Node, partNumbers []*Node) []int {
	partNumberSet := make(map[*Node]struct{}, len(partNumbers))
	for _, num := range partNumbers {
		partNumberSet[num] = struct{}{}
	}

	res := make([]int, 0, 4)
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i][j].kind == KindSymbol && matrix[i][j].symbol == "*" {
				adjPartNums := getAdjacentPartNumbers(matrix, i, j, partNumberSet)
				if len(adjPartNums) == 2 {
					res = append(res, adjPartNums[0].number*adjPartNums[1].number)
				}
			}
		}
	}
	return res
}

func getAdjacentPartNumbers(matrix [][]Node, i, j int, partSet map[*Node]struct{}) []*Node {
	res := make([]*Node, 0, 4)
	row := matrix[i]
	curr := row[j]
	left, right := &row[j-1], &row[j+1]

	if _, ok := partSet[left]; ok && left.colR == curr.colL {
		res = append(res, left)
	}
	if _, ok := partSet[right]; ok && right.colL == curr.colR {
		res = append(res, right)
	}
	lastRow := matrix[i-1]
	nextRow := matrix[i+1]
	for k := 0; k < len(lastRow); k++ {
		up := &lastRow[k]
		if _, ok := partSet[up]; ok && (up.colR >= curr.colL && up.colL <= curr.colR) {
			res = append(res, up)
		}
	}
	for k := 0; k < len(nextRow); k++ {
		down := &nextRow[k]
		if _, ok := partSet[down]; ok && (down.colR >= curr.colL && down.colL <= curr.colR) {
			res = append(res, down)
		}
	}

	return res
}

func isAdjacentToSymbol(matrix [][]Node, i, j int) bool {
	row := matrix[i]
	curr := row[j]
	left, right := row[j-1], row[j+1]
	if left.kind == KindSymbol && left.colR == curr.colL {
		return true
	}
	if right.kind == KindSymbol && right.colL == curr.colR {
		return true
	}
	lastRow := matrix[i-1]
	nextRow := matrix[i+1]
	for k := 0; k < len(lastRow); k++ {
		if up := lastRow[k]; up.kind == KindSymbol && (up.colR >= curr.colL && up.colL <= curr.colR) {
			return true
		}
	}
	for k := 0; k < len(nextRow); k++ {
		if down := nextRow[k]; down.kind == KindSymbol && (down.colR >= curr.colL && down.colL <= curr.colR) {
			return true
		}
	}
	return false
}

type Node struct {
	symbol string
	number int
	kind   int
	row    int // node locate in Row row, [colL, colR) character(bytes)
	colL   int
	colR   int
}

func (n *Node) String() string {
	switch n.kind {
	case KindPeriod:
		return "."
	case KindNumber:
		return strconv.FormatInt(int64(n.number), 10)
	case KindSymbol:
		return n.symbol
	default:
		return fmt.Sprintf("%v", *n)
	}
}

const (
	KindPeriod int = iota
	KindSymbol
	KindNumber
)

// expand the matrix boarder with sentinel
func parseInput(r io.Reader) ([][]Node, int) {
	scanner := bufio.NewScanner(r)

	res := make([][]Node, 0, 2)
	n := 0

	res = append(res, []Node{{kind: KindPeriod, row: -1}})
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		nodes := parseLine(line, row)
		res = append(res, nodes)
		n += len(nodes) - 2 // delete the sentinel
	}
	res = append(res, []Node{{kind: KindPeriod, row: row + 1}})
	return res, n
}

// Add sentinel into the start and end of the line
func parseLine(s string, row int) []Node {
	res := make([]Node, 0, 2)
	res = append(res, Node{kind: KindPeriod, row: row, colL: -1, colR: -1})
	i := 0
	for {
		for ; i < len(s) && s[i] == '.'; i++ {
		}
		if i >= len(s) {
			break
		}
		j := i
		if isDigit(s[i]) {
			for ; j < len(s) && isDigit(s[j]); j++ {
			}
			if j > i {
				num, _ := strconv.ParseInt(s[i:j], 10, 64)
				n := Node{number: int(num), kind: KindNumber, row: row, colL: i, colR: j}
				res = append(res, n)
			}
		} else {
			for ; j < len(s) && isSymbol(s[j]); j++ {
			}
			if j > i {
				n := Node{symbol: s[i:j], kind: KindSymbol, row: row, colL: i, colR: j}
				res = append(res, n)
			}
		}
		i = j
	}
	res = append(res, Node{kind: KindPeriod, row: row, colL: len(s), colR: len(s)})
	return res
}

func isDigit(d uint8) bool {
	return d >= '0' && d <= '9'
}

func isSymbol(d uint8) bool {
	return d != '.' && !isDigit(d)
}

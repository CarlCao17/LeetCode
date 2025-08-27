package main

func findInSortMatrix(ma [][]int, num int) bool {
	m, n := len(ma), len(ma[0])
	if m <= 0 || n <= 0 {
		return false
	}
	for i, j := 0, n-1; i >= 0 && j >= 0; {
		if ma[i][j] == num {
			return true
		}
		if ma[i][j] > num {
			i++
		} else {
			j--
		}
	}
	return false
}

package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for i >= 0 && j >= 0 {
		i = backspaceDelWith(s, i)
		j = backspaceDelWith(t, j)
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 {
			return false
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	if i > 0 {
		i = backspaceDelWith(s, i)
	}
	if j > 0 {
		j = backspaceDelWith(t, j)
	}
	if i < 0 && j < 0 {
		return true
	}
	return false
}

func backspaceDel(s string) int {
	return backspaceDelWith(s, len(s)-1)
}

func backspaceDelWith(s string, i int) int {
	for { // 连续消去足够多的 '#' 直到不能再消除为止
		if i < 0 || s[i] != '#' {
			break
		}
		num := 0
		for i >= 0 && s[i] == '#' {
			i--
			num++
		}
		for i >= 0 && num > 0 {
			if s[i] == '#' {
				num++
				i--
			} else {
				i--
				num--
			}
		}
	}
	return i
}

//
//func main() {
//	backspaceCompare("nzp#o#g", "b#nzp#o#g")
//}

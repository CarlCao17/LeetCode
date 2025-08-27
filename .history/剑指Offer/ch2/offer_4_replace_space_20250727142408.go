package ch2

import "strings"

func ReplaceSpace(s string) string {
	var sb strings.Builder
	nSpace := 0
	for _, c := range s {
		if c == ' ' {
			nSpace++
		}
	}
	sb.Grow(len(s) + nSpace*2)
	for _, c := range s {
		if c == ' ' {
			sb.WriteString("%20")
			continue
		}
		sb.WriteRune(c)
	}
	return sb.String()
}

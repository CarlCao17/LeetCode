package utils

import "fmt"

func Assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func AssertF(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Sprintf(format, args...))
	}
}

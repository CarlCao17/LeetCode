package utils

import "encoding/json"

func ToTwoDimensionSlices[T any](s string) [][]T {
	var t [][]T
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	return t
}

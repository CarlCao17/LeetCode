package dp

import "slices"

func maxEnvelopes(envelopes [][]int) int {
	slices.SortFunc(envelopes, func(a []int, b []int) int {
		if a[0] == b[0] {
			return b[1] - a[1]
		}
		return a[0] - b[0]
	})
	height := make([]int, 0, len(envelopes))
	height = append(height, envelopes[0][1])
	for i := 1; i < len(envelopes); i++ {
		if envelopes[i][1] > height[len(height)-1] {
			height = append(height, envelopes[i][1])
		} else {
			idx, _ := slices.BinarySearch(height, envelopes[i][1])
			height[idx] = envelopes[i][1]
		}
	}
	return len(height)
}

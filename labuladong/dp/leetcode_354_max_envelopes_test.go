package dp

import (
	"leetcode/utils"
	"testing"
)

func Test_maxEnvelopes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		envelopes [][]int
		want      int
	}{
		{
			envelopes: utils.ToTwoDimensionSlices[int](`[[5,4],[6,4],[6,7],[2,3]]`),
			want:      3,
		},
		{
			envelopes: utils.ToTwoDimensionSlices[int](`[[1,1],[1,1],[1,1]]`),
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxEnvelopes(tt.envelopes)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

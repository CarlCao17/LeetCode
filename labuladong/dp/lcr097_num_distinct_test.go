package dp

import "testing"

func Test_numDistinct(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want int
	}{
		{
			s: "babgbag", t: "bag",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDistinct(tt.s, tt.t)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

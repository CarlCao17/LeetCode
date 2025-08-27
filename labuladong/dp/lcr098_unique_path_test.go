package dp

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		m    int
		n    int
		want int
	}{
		{
			name: "100x100",
			m:    100, n: 100,
			want: 4631081169483718960,
		},
		{
			name: "100x1",
			m:    100, n: 1,
			want: 1,
		},
		{
			name: "1x100",
			m:    1, n: 100,
			want: 1,
		},
		{
			name: "3x7",
			m:    3, n: 7,
			want: 28,
		},
		{
			name: "3x2",
			m:    3, n: 2,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

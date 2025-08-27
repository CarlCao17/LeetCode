package dp

import "testing"

func Test_maxValueOf01Package(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		wt   []int
		vt   []int
		W    int
		want int
	}{
		{
			name: "case: descend",
			wt:   []int{1, 2, 3, 4, 5, 6, 7},
			vt:   []int{7, 6, 5, 4, 3, 2, 1},
			W:    7,
			want: 18,
		},
		{
			name: "case: ascend",
			wt:   []int{1, 2, 3, 4, 5, 6, 7},
			vt:   []int{1, 2, 3, 4, 5, 6, 7},
			W:    7,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxValueOf01Package(tt.wt, tt.vt, tt.W)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("maxValueOf01Package() = %v, want %v", got, tt.want)
			}
		})
	}
}

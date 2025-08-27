package ch2

import "testing"

func TestFindInSortMatrix(t *testing.T) {
	type args struct {
		ma  [][]int
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInSortMatrix(tt.args.ma, tt.args.num); got != tt.want {
				t.Errorf("FindInSortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

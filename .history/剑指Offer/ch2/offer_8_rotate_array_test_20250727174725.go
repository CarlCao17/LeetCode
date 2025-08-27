package ch2

import "testing"

func TestFindSmallestNumInRotateArray(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSmallestNumInRotateArray(tt.args.a); got != tt.want {
				t.Errorf("FindSmallestNumInRotateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

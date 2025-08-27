package ch2

import "testing"

func Test_binarySearchL(t *testing.T) {
	type args struct {
		a  []int
		lo int
		hi int
		n  int
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
			if got := binarySearchL(tt.args.a, tt.args.lo, tt.args.hi, tt.args.n); got != tt.want {
				t.Errorf("binarySearchL() = %v, want %v", got, tt.want)
			}
		})
	}
}

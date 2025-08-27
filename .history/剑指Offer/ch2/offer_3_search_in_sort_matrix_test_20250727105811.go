package ch2

import "testing"

func TestFindInSortMatrix(t *testing.T) {
	ma := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	type args struct {
		ma  [][]int
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				ma: ma,
				num: 1,
			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInSortMatrix(tt.args.ma, tt.args.num); got != tt.want {
				t.Errorf("FindInSortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

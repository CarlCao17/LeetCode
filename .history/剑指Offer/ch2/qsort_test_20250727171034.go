package ch2

import "testing"

func Test_partition(t *testing.T) {
	type args struct {
		arr []int
		lo  int
		hi  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "increasing seq",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				lo:  0,
				hi:  4,
			},
			want:  4,
			want1: []int{1, 2, 3, 4, 5},
		},
		{
			name: "decreasing seq",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
				lo:  0,
				hi:  4,
			},
			want:  0,
			want1: []int{1, 4, 3, 2, 5},
		},
		{
			name: "len=2",
			args: args{
				arr: []int{1, 3},
				lo:  0,
				hi:  1,
			},
			want:  1,
			want1: []int{1, 3},
		},
		{
			name: "len=3",
			args: args{
				arr: []int{1, 3, 2},
				lo:  0,
				hi:  2,
			},
			want:  1,
			want1: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.lo, tt.args.hi); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

package ch2

import "testing"

func Test_partition(t *testing.T) {
	type args struct {
		arr []int
		lo  int
		hi  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "increasing seq",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				lo:  0,
				hi:  4,
			},
			want: 4,
		},
		{
			name: "decreasing seq",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
				lo:  0,
				hi:  4,
			},
			want: 0,
		},
		{
			name: "len=1",
			args: args{
				arr: []int{1},
				lo:  0,
				hi:  0,
			},
			want: 0,
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

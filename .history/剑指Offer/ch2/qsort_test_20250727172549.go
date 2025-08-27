package ch2

import (
	"reflect"
	"testing"
)

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
			if !reflect.DeepEqual(tt.args.arr, tt.want1) {
				t.Errorf("after partition: arr should be %v, got=%v", tt.want1, tt.args.arr)
			}
		})
	}
}

func TestFindKLargest(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "corner case: empty array",
			args: args{
				arr: []int{},
				k:   0,
			},
			want: -1,
		},
		{
			name: "invalid param: k=0",
			args: args{
				arr: []int{1, 2, 3},
				k:   0,
			},
			want: -1,
		},
		{
			name: "invalid param: k",
			args: args{
				arr: []int{1, 2, 3},
				k:   4,
			},
			want: -1,
		},
		{
			name: "corner case: len(arr)=1",
			args: args{
				arr: []int{1},
				k:   1,
			},
			want: 1,
		},
		{
			name: "find k=5, len(arr)=5: increasing",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   5,
			},
			want: 5,
		},
		{
			name: "find k=5, len(arr)=5: decreasing",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
				k:   5,
			},
			want: 5,
		},
		{
			name: "find k=5, len(arr)=5: some order",
			args: args{
				arr: []int{5, 1, 2, 3, 4},
				k:   5,
			},
			want: 5,
		},
		{
			name: "find k=5, len(arr)=5: random",
			args: args{
				arr: []int{4, 3, 5, 1, 2},
				k:   5,
			},
			want: 5,
		},
		{
			name: "find k=1, len(arr)=5: random",
			args: args{
				arr: []int{5, 1, 2, 3, 4},
				k:   1,
			},
			want: 1,
		},
		{
			name: "find k=3, len(arr)=5: random",
			args: args{
				arr: []int{5, 1, 2, 3, 4},
				k:   3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKLargest(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("FindKLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

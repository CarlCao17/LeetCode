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
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  2,
			},
			want: 1,
		},
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  0,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  1,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  0,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  2,
			},
			want: 1,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  3,
			},
			want: 4,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  4,
			},
			want: 6,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  5,
			},
			want: 7,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  6,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchL(tt.args.a, tt.args.lo, tt.args.hi, tt.args.n); got != tt.want {
				t.Errorf("binarySearchL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchR(t *testing.T) {
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
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  2,
			},
			want: 1,
		},
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  0,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1},
				lo: 0,
				hi: 0,
				n:  1,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  0,
			},
			want: 0,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  2,
			},
			want: 3,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  3,
			},
			want: 5,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  4,
			},
			want: 6,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  5,
			},
			want: 7,
		},
		{
			args: args{
				a:  []int{1, 2, 2, 2, 3, 3, 4, 5},
				lo: 0,
				hi: 7,
				n:  6,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchL(tt.args.a, tt.args.lo, tt.args.hi, tt.args.n); got != tt.want {
				t.Errorf("binarySearchL() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				ma:  ma,
				num: 1,
			},
			want: true,
		},
		{
			args: args{
				ma:  ma,
				num: -1,
			},
			want: false,
		},
		{
			args: args{
				ma:  ma,
				num: 9,
			},
			want: true,
		},
		{
			args: args{
				ma:  ma,
				num: 8,
			},
			want: true,
		},
		{
			args: args{
				ma:  ma,
				num: 10,
			},
			want: true,
		},
		{
			args: args{
				ma:  ma,
				num: 5,
			},
			want: false,
		},
		{
			args: args{
				ma:  ma,
				num: 16,
			},
			want: false,
		},
		{
			args: args{
				ma:  ma,
				num: 12,
			},
			want: true,
		},
		{
			args: args{
				ma:  ma,
				num: 14,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInSortMatrix(tt.args.ma, tt.args.num); got != tt.want {
				t.Errorf("FindInSortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
